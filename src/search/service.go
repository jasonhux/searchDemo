package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"searchDemo/src/data"
	"searchDemo/src/interaction"
	"strconv"
	"strings"
	"time"
)

type Service interface {
	SetStructMap() (err error)
	Search() (results string, isQuit bool, err error)
	DirectSearchWithValue() (results string, isQuit bool, err error)
	RequestNewSearch() bool
}

type service struct {
	DataService        data.Service
	InteractionService interaction.Service
	StructMap          map[string]map[string]data.Field
	SelectedStructKey  string
	SelectedFieldKey   string
}

func NewService(dataService data.Service, interactionService interaction.Service) Service {
	return &service{DataService: dataService, InteractionService: interactionService}
}

//Search func retrieves the user input and process the required search on the keywords given;
//It returns results in string format if the search is successful; isQuit as true if user type 'quit' during the interaction; and error message if any error happens
func (s *service) Search() (results string, isQuit bool, err error) {
	fmt.Println("Welcome to Zendesk search. The search param is case insensitive. You can type 'quit' to leave the application")
	fmt.Println("Select 1) Tickets or 2) Users or 3) Organizations")
	isQuit, searchStructParam := s.InteractionService.GetUserInput()
	if isQuit {
		return
	}
	fieldMap, err := s.setSearchStruct(searchStructParam)
	if err != nil {
		return
	}

	fmt.Println("Available search field")
	fmt.Println("======================")
	for k := range fieldMap {
		fmt.Println(k)
	}
	fmt.Println("======================")
	fmt.Println("Please enter a search field from the above list")
	isQuit, searchFieldParam := s.InteractionService.GetUserInput()
	if isQuit {
		return
	}
	typeName, err := s.setSearchFieldValue(searchFieldParam)
	if err != nil {
		return
	}

	fmt.Println("Please enter the search value. The search value type is:", typeName)
	if typeName == "[]string" {
		fmt.Println("You just need to type in a string and any slices contain your search value is treated as matched slices")
	}
	isQuit, searchValueParam := s.InteractionService.GetUserInput()
	if isQuit {
		return
	}
	start := time.Now()
	resultList, err := retrieveResults(s.SelectedStructKey, s.SelectedFieldKey, searchValueParam, s.StructMap)
	if err != nil {
		return
	}

	//Convert to json string for output
	resultsBytes, err := json.Marshal(resultList)
	if err != nil {
		return
	}
	results = string(resultsBytes)
	//remove later
	colapsed := time.Now().Sub(start)
	results += fmt.Sprintf("%v", colapsed)
	return
}

func (s *service) DirectSearchWithValue() (results string, isQuit bool, err error) {
	// fmt.Println("Please enter the search value.")
	// isQuit, value := s.InteractionService.GetUserInput()
	// if isQuit {
	// 	return
	// }
	value := "false"
	keyMap := map[string]string{
		"1": "tickets",
		"2": "users",
		"3": "organizations",
	}

	start := time.Now()
	combinedResultsMap := map[string][]interface{}{}
	for structKey := range s.StructMap {
		resultMapKey := keyMap[structKey]
		resultMap := map[string]bool{}

		for fieldKey := range s.StructMap[structKey] {
			resultList, err := retrieveResults(structKey, fieldKey, value, s.StructMap)
			if err != nil {
				continue
			}
			existingResultList, _ := combinedResultsMap[resultMapKey]
			updatedResultList := existingResultList
			for _, result := range resultList {
				pointerKey := fmt.Sprintf("%v", &result)
				isExist, _ := resultMap[pointerKey]
				if !isExist {
					resultMap[pointerKey] = true
					updatedResultList = append(updatedResultList, result)

				}

			}
			combinedResultsMap[resultMapKey] = updatedResultList
		}
	}
	resultsMapBytes, err := json.Marshal(combinedResultsMap)
	if err != nil {
		return
	}
	results = string(resultsMapBytes)
	colapsed := time.Now().Sub(start)
	fmt.Println(colapsed)
	return
}

func (s *service) RequestNewSearch() bool {
	fmt.Println("Type 'n' or 'quit' to quit or any other key to start a new search")
	isQuit, input := s.InteractionService.GetUserInput()
	if isQuit || input == "n" {
		return false
	}
	return true
}

func (s *service) SetStructMap() (err error) {
	start := time.Now()
	tickets, users, organizations, err := s.DataService.LoadFile()
	if err != nil {
		return
	}
	structMap, err := s.DataService.PrepareStructMap(tickets, users, organizations)
	if err == nil {
		s.StructMap = structMap
	}
	colapsed := time.Now().Sub(start)
	fmt.Println(colapsed)
	return
}

func (s *service) setSearchStruct(param string) (fieldMap map[string]data.Field, err error) {
	fieldMap, ok := s.StructMap[param]
	if !ok {
		err = errors.New("No struct found")
		return
	}
	s.SelectedStructKey = param
	return
}

func (s *service) setSearchFieldValue(param string) (fieldType string, err error) {
	paramLowerCase := strings.ToLower(param)
	fieldMap, _ := s.StructMap[s.SelectedStructKey]
	field, ok := fieldMap[paramLowerCase]
	if !ok {
		err = errors.New("No field found")
		return
	}
	s.SelectedFieldKey = paramLowerCase
	return field.Type, nil
}

func retrieveResults(structKey, FieldKey, param string, structMap map[string]map[string]data.Field) (results []interface{}, err error) {
	paramLowerCase := strings.ToLower(param)
	fieldMap, _ := structMap[structKey]
	Field, _ := fieldMap[FieldKey]
	resultsList, ok := Field.ValueMap[paramLowerCase]
	if !ok {
		err = errors.New("No results found")
		return
	}
	return processResults(resultsList, structMap)
}

func processResults(resultsList []interface{}, structMap map[string]map[string]data.Field) (processedResults []interface{}, err error) {

	switch resultsList[0].(type) {
	case *data.Ticket:
		return processTicketResults(resultsList, structMap)
	case *data.User:
		return processUserResults(resultsList, structMap)
	case *data.Organization:
		return processOrganizationResults(resultsList, structMap)
	}
	err = errors.New("No matched type for process")
	return
}

func processTicketResults(resultsList []interface{}, structMap map[string]map[string]data.Field) (processedResults []interface{}, err error) {
	//Skip to check map contains the key here as if the struct map is not complete, the processData step should have already reported errors.
	userMap, _ := structMap["2"]
	organizationMap, _ := structMap["3"]
	ticketsForDisplay := []data.TicketForDisplay{}

	for _, result := range resultsList {
		ticket := result.(*data.Ticket)

		//Find linked struct to the ticket, such as assignee, submitter and organization.

		ulist := getLinkedStructs(strconv.Itoa(ticket.AssigneeID), userMap["id"])
		//Relationship between ticket and assignee is 1:1; thus take the first user pointer
		//If ulist is empty, assignee is a empty struct
		assignee := &data.User{}
		if len(ulist) > 0 {
			assignee = ulist[0].(*data.User)
		}

		ulist = getLinkedStructs(strconv.Itoa(ticket.SubmitterID), userMap["id"])
		submitter := &data.User{}
		if len(ulist) > 0 {
			submitter = ulist[0].(*data.User)
		}

		orgList := getLinkedStructs(strconv.Itoa(ticket.OrganizationID), organizationMap["id"])
		org := &data.Organization{}
		if len(orgList) > 0 {
			org = orgList[0].(*data.Organization)
		}

		//Create a ticketForDisplay struct which contains the ticket information, and the linked struct values
		ticketForDisplay := data.TicketForDisplay{Ticket: *ticket, AssigneeName: assignee.Name, SubmitterName: submitter.Name, OrganizationName: org.Name}
		ticketsForDisplay = append(ticketsForDisplay, ticketForDisplay)
	}
	if len(ticketsForDisplay) == 0 {
		err = errors.New("No tickets are available in the search")
		return
	}
	for _, ticket := range ticketsForDisplay {
		processedResults = append(processedResults, ticket)
	}
	return
}

func processUserResults(resultsList []interface{}, structMap map[string]map[string]data.Field) (processedResults []interface{}, err error) {
	ticketMap, _ := structMap["1"]
	organizationMap, _ := structMap["3"]
	usersForDisplay := []data.UserForDisplay{}
	for _, result := range resultsList {

		assignedTicketIDs := []string{}
		submittedTicketIDs := []string{}
		user := result.(*data.User)

		//Find linked struct to the user, such as assigned tickets, submitted tickets and organization.
		orgList := getLinkedStructs(strconv.Itoa(user.OrganizationID), organizationMap["id"])
		org := &data.Organization{}
		if len(orgList) > 0 {
			org = orgList[0].(*data.Organization)
		}
		ticketList := getLinkedStructs(strconv.Itoa(user.ID), ticketMap["assigneeid"])

		for _, t := range ticketList {
			ticket := t.(*data.Ticket)
			assignedTicketIDs = append(assignedTicketIDs, ticket.ID)
		}

		ticketList = getLinkedStructs(strconv.Itoa(user.ID), ticketMap["submitterid"])

		for _, t := range ticketList {
			ticket := t.(*data.Ticket)
			submittedTicketIDs = append(submittedTicketIDs, ticket.ID)
		}

		userForDisplay := data.UserForDisplay{User: *user, OrganizationName: org.Name, SubmittedTicketIDs: submittedTicketIDs, AssignedTicketsIDs: assignedTicketIDs}
		usersForDisplay = append(usersForDisplay, userForDisplay)
	}
	if len(usersForDisplay) == 0 {
		err = errors.New("No users are available in the search")
		return
	}
	for _, user := range usersForDisplay {
		processedResults = append(processedResults, user)
	}
	return
}

func processOrganizationResults(resultsList []interface{}, structMap map[string]map[string]data.Field) (processedResults []interface{}, err error) {
	ticketMap, _ := structMap["1"]
	userMap, _ := structMap["2"]
	orgsForDisplay := []data.OrganizationForDisplay{}
	for _, result := range resultsList {
		userNames := []string{}
		ticketIDs := []string{}
		organization := result.(*data.Organization)

		//Find linked struct to the organization, such as all users and tickets with matched organization id.
		ticketList := getLinkedStructs(strconv.Itoa(organization.ID), ticketMap["organizationid"])
		for _, t := range ticketList {
			ticket := t.(*data.Ticket)
			ticketIDs = append(ticketIDs, ticket.ID)
		}

		userList := getLinkedStructs(strconv.Itoa(organization.ID), userMap["organizationid"])
		for _, t := range userList {
			user := t.(*data.User)
			userNames = append(userNames, user.Name)
		}
		orgForDisplay := data.OrganizationForDisplay{Organization: *organization, TicketIDs: ticketIDs, UserNames: userNames}
		orgsForDisplay = append(orgsForDisplay, orgForDisplay)
	}
	if len(orgsForDisplay) == 0 {
		err = errors.New("No organizations are available in the search")
		return
	}
	for _, org := range orgsForDisplay {
		processedResults = append(processedResults, org)
	}
	return
}

func getLinkedStructs(value string, linkedField data.Field) (linkedStructs []interface{}) {
	//If there is no available key in the value map, return an empty linkedStructs back; this means the searched struct has no linked structs on the requested field
	linkedStructs, _ = linkedField.ValueMap[value]
	return
}
