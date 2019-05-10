package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Service interface {
	SetSearchStruct(param string) (fieldMap map[string]Field, err error)
	SetSearchFieldValue(param string) error
	Search(param string) (resultsList []interface{}, err error)
}

type service struct {
	StructMap         map[string]map[string]Field
	SelectedStructKey string
	SelectedFieldKey  string
}

type Field struct {
	Type     string
	ValueMap map[string][]interface{}
}

func NewService(structMap map[string]map[string]Field) Service {
	return &service{StructMap: structMap}
}

func PrepareStructMap(tickets []*Ticket, users []*User, organizations []*Organization) map[string]map[string]Field {
	//Convert tickets, users and organizations slice to []interface{} so they can share the same ProcessFieldList func which takes []interface{} as param
	tList := make([]interface{}, len(tickets))
	for i, v := range tickets {
		tList[i] = v
	}
	uList := make([]interface{}, len(users))
	for i, v := range users {
		uList[i] = v
	}
	oList := make([]interface{}, len(organizations))
	for i, v := range organizations {
		oList[i] = v
	}
	return map[string]map[string]Field{
		"1": ProcessFieldMap(tList),
		"2": ProcessFieldMap(uList),
		"3": ProcessFieldMap(oList),
	}
}

//ProcessFieldMap func is to convert object list such as tickets to a map structure; map key is the struct field name (ex. Name). The map value is the Field struct;
//Field struct contains the 1. map key's type (ex. string, int) and 2. a nested value map;
//The nested value map has the struct field value in string format as the key (ex. "A Drama in Gabon", or "true"); the map value are a list of pointers to the structs which contains the map key;
//When a field contains an array, treat each string in the array as a separate value map key;
func ProcessFieldMap(structList []interface{}) map[string]Field {
	fmt.Println(len(structList))

	fieldMap := initFieldMap(structList[0])
	for _, ticket := range structList {
		v := reflect.ValueOf(ticket).Elem()
		for k := range fieldMap {
			updatedTicketList := []interface{}{}
			switch fieldMap[k].Type {
			case "[]string":
				list, ok := v.FieldByName(k).Interface().([]string)
				if ok {
					for _, element := range list {
						matchedTicketList, _ := fieldMap[k].ValueMap[element]
						updatedTicketList = append(matchedTicketList, ticket)
						fieldMap[k].ValueMap[element] = updatedTicketList
					}
				}
				break
			default:
				fieldValue := fmt.Sprintf("%v", v.FieldByName(k))
				matchedTicketList, _ := fieldMap[k].ValueMap[fieldValue]
				updatedTicketList = append(matchedTicketList, ticket)
				fieldMap[k].ValueMap[fieldValue] = updatedTicketList
			}
		}
	}
	return fieldMap
}

func initFieldMap(instance interface{}) map[string]Field {
	v := reflect.ValueOf(instance).Elem()
	fieldMap := map[string]Field{}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		n := v.Type().Field(i).Name
		t := f.Type().String()
		fieldMap[n] = Field{Type: t, ValueMap: map[string][]interface{}{}}
	}
	return fieldMap
}

func (s *service) SetSearchStruct(param string) (fieldMap map[string]Field, err error) {
	fieldMap, ok := s.StructMap[param]
	if !ok {
		err = errors.New("No struct found")
		return
	}
	s.SelectedStructKey = param
	return
}

func (s *service) SetSearchFieldValue(param string) error {
	fieldMap, _ := s.StructMap[s.SelectedStructKey]
	_, ok := fieldMap[param]
	if !ok {
		return errors.New("No field found")
	}
	s.SelectedFieldKey = param
	//return type as well for notice
	return nil
}

func (s *service) Search(param string) (resultsList []interface{}, err error) {
	fieldMap, _ := s.StructMap[s.SelectedStructKey]
	Field, _ := fieldMap[s.SelectedFieldKey]
	resultsList, ok := Field.ValueMap[param]
	if !ok {
		err = errors.New("No results found")
	}

	s.processResults(resultsList)

	return
}

func (s *service) processResults(resultsList []interface{}) (processedResultsList []interface{}, err error) {

	switch resultsList[0].(type) {
	case *Ticket:
		return processTicketResults(resultsList, s.StructMap)
	case *User:
		return processUserResults(resultsList, s.StructMap)
	case *Organization:
		return processOrganizationResults(resultsList, s.StructMap)
	}
	err = errors.New("No matched type for process")
	return
}

func processTicketResults(resultsList []interface{}, structMap map[string]map[string]Field) (processedResultsList []interface{}, err error) {
	//No need to check map contains the key here as if the struct map is not complete, the processData step should have already reported errors
	userMap, _ := structMap["2"]
	organizationMap, _ := structMap["3"]
	ticketsForDisplay := []TicketForDisplay{}

	for _, result := range resultsList {
		ticket := result.(*Ticket)
		//get assignee name
		ulist, e := getLinkedUsers(strconv.Itoa(ticket.AssigneeID), userMap["ID"])
		if e != nil {
			// err = e
			// return
			continue
		}
		//relationship between ticket and assignee is 1:1; thus take the first user pointer
		//add if length = 0 check???
		assignee := ulist[0]

		ulist, e = getLinkedUsers(strconv.Itoa(ticket.SubmitterID), userMap["ID"])
		if e != nil {
			// err = e
			// return
			continue
		}
		submitter := ulist[0]

		orgList, e := getLinkedOrganizations(strconv.Itoa(ticket.OrganizationID), organizationMap["ID"])
		if e != nil {
			// err = e
			// return
			continue
		}
		org := orgList[0]

		ticketForDisplay := TicketForDisplay{Ticket: *ticket, AssigneeName: assignee.Name, SubmitterName: submitter.Name, OrganizationName: org.Name}
		//fmt.Printf("%+v\n", TicketForDisplay)
		ticketsForDisplay = append(ticketsForDisplay, ticketForDisplay)
	}
	b, _ := json.Marshal(ticketsForDisplay)
	println(string(b))
	return resultsList, nil
}

func processUserResults(resultsList []interface{}, structMap map[string]map[string]Field) (processedResultsList []interface{}, err error) {
	for _, result := range resultsList {
		user := result.(*User)
		fmt.Printf("%+v\n", user)
	}
	return resultsList, nil
}

func processOrganizationResults(resultsList []interface{}, structMap map[string]map[string]Field) (processedResultsList []interface{}, err error) {
	for _, result := range resultsList {
		organization := result.(*Organization)
		fmt.Printf("%+v\n", organization)
	}
	return resultsList, nil
}

func getLinkedUsers(value string, userField Field) (users []*User, err error) {
	matchedUserPtrs, ok := userField.ValueMap[value]
	if !ok || !(len(matchedUserPtrs) > 0) {
		err = errors.New("No linked user was found")
		return

	}
	for _, userPtr := range matchedUserPtrs {
		users = append(users, userPtr.(*User))
	}
	return
}

func getLinkedOrganizations(value string, organizationField Field) (organizations []*Organization, err error) {
	matchedOrgPtrs, ok := organizationField.ValueMap[value]
	if !ok || !(len(matchedOrgPtrs) > 0) {
		err = errors.New("No linked organization was found")
		return

	}
	for _, orgPtr := range matchedOrgPtrs {
		organizations = append(organizations, orgPtr.(*Organization))
	}
	return
}
