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

	fmt.Println("Please enter search value. The search value type is:", typeName)
	if typeName == "[]string" {
		fmt.Println("You just need to type in a string and any slices contain your search value is treated as matched slices")
	}
	isQuit, searchValueParam := s.InteractionService.GetUserInput()
	if isQuit {
		return
	}
	start := time.Now()
	results, err = s.retrieveResults(searchValueParam)
	colapsed := time.Now().Sub(start)
	results += fmt.Sprintf("%v", colapsed)
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
	//return type as well for notice
	return field.Type, nil
}

func (s *service) retrieveResults(param string) (results string, err error) {
	paramLowerCase := strings.ToLower(param)
	fieldMap, _ := s.StructMap[s.SelectedStructKey]
	Field, _ := fieldMap[s.SelectedFieldKey]
	resultsList, ok := Field.ValueMap[paramLowerCase]
	if !ok {
		err = errors.New("No results found")
		return
	}
	return s.processResults(resultsList)
}

func (s *service) processResults(resultsList []interface{}) (processedResults string, err error) {

	switch resultsList[0].(type) {
	case *data.Ticket:
		return processTicketResults(resultsList, s.StructMap)
	case *data.User:
		return processUserResults(resultsList, s.StructMap)
	case *data.Organization:
		return processOrganizationResults(resultsList, s.StructMap)
	}
	err = errors.New("No matched type for process")
	return
}

func processTicketResults(resultsList []interface{}, structMap map[string]map[string]data.Field) (processedResults string, err error) {
	//No need to check map contains the key here as if the struct map is not complete, the processData step should have already reported errors
	userMap, _ := structMap["2"]
	organizationMap, _ := structMap["3"]
	ticketsForDisplay := []data.TicketForDisplay{}

	for _, result := range resultsList {
		ticket := result.(*data.Ticket)
		//get assignee name
		ulist := getLinkedStructs(strconv.Itoa(ticket.AssigneeID), userMap["id"])
		//relationship between ticket and assignee is 1:1; thus take the first user pointer
		//add if length = 0 check???
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

		ticketForDisplay := data.TicketForDisplay{Ticket: *ticket, AssigneeName: assignee.Name, SubmitterName: submitter.Name, OrganizationName: org.Name}
		ticketsForDisplay = append(ticketsForDisplay, ticketForDisplay)
	}
	if len(ticketsForDisplay) == 0 {
		err = errors.New("No tickets are available in the search")
		return
	}
	b, err := json.Marshal(ticketsForDisplay)
	return string(b), err
}

func processUserResults(resultsList []interface{}, structMap map[string]map[string]data.Field) (processedResults string, err error) {
	ticketMap, _ := structMap["1"]
	organizationMap, _ := structMap["3"]
	usersForDisplay := []data.UserForDisplay{}
	for _, result := range resultsList {
		assignedTicketIDs := []string{}
		submittedTicketIDs := []string{}
		user := result.(*data.User)
		//get organization name that user belongs to
		orgList := getLinkedStructs(strconv.Itoa(user.OrganizationID), organizationMap["id"])
		org := &data.Organization{}
		if len(orgList) > 0 {
			org = orgList[0].(*data.Organization)
		}
		//get tickets that the user id is shown in the assignee field
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
	b, err := json.Marshal(usersForDisplay)
	return string(b), err
}

func processOrganizationResults(resultsList []interface{}, structMap map[string]map[string]data.Field) (processedResults string, err error) {
	ticketMap, _ := structMap["1"]
	userMap, _ := structMap["2"]
	orgsForDisplay := []data.OrganizationForDisplay{}
	for _, result := range resultsList {
		userNames := []string{}
		ticketIDs := []string{}
		organization := result.(*data.Organization)
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
	b, err := json.Marshal(orgsForDisplay)
	return string(b), err
}

func getLinkedStructs(value string, linkedField data.Field) (linkedStructs []interface{}) {
	//if there is no available key in the value map, return an empty linkedStructs back; this means the searched struct has no linked structs on the requested field
	linkedStructs, _ = linkedField.ValueMap[value]
	return
}
