package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"searchDemo/src/data"
	"strconv"
)

type Service interface {
	SetStructMap() (err error)
	SetSearchStruct(param string) (fieldMap map[string]data.Field, err error)
	SetSearchFieldValue(param string) error
	Search(param string) (results string, err error)
}

type service struct {
	DataService       data.Service
	StructMap         map[string]map[string]data.Field
	SelectedStructKey string
	SelectedFieldKey  string
}

func NewService(dataService data.Service) Service {
	return &service{DataService: dataService}
}

func (s *service) SetStructMap() (err error) {
	tickets, users, organizations, err := s.DataService.LoadFile()
	if err != nil {
		return
	}
	structMap, err := s.DataService.PrepareStructMap(tickets, users, organizations)
	if err == nil {
		s.StructMap = structMap
	}
	return
}

func (s *service) SetSearchStruct(param string) (fieldMap map[string]data.Field, err error) {
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

func (s *service) Search(param string) (results string, err error) {
	fieldMap, _ := s.StructMap[s.SelectedStructKey]
	Field, _ := fieldMap[s.SelectedFieldKey]
	resultsList, ok := Field.ValueMap[param]
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
		ulist, e := getLinkedStructs(strconv.Itoa(ticket.AssigneeID), userMap["ID"])
		if e != nil {
			// err = e
			// return
			continue
		}
		//relationship between ticket and assignee is 1:1; thus take the first user pointer
		//add if length = 0 check???
		assignee := ulist[0].(*data.User)

		ulist, e = getLinkedStructs(strconv.Itoa(ticket.SubmitterID), userMap["ID"])
		if e != nil {
			// err = e
			// return
			continue
		}
		submitter := ulist[0].(*data.User)

		orgList, e := getLinkedStructs(strconv.Itoa(ticket.OrganizationID), organizationMap["ID"])
		if e != nil {
			// err = e
			// return
			continue
		}
		org := orgList[0].(*data.Organization)

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
	for _, result := range resultsList {
		user := result.(*data.User)
		fmt.Printf("%+v\n", user)
	}
	//to do
	return "test", nil
}

func processOrganizationResults(resultsList []interface{}, structMap map[string]map[string]data.Field) (processedResults string, err error) {
	for _, result := range resultsList {
		organization := result.(*data.Organization)
		fmt.Printf("%+v\n", organization)
	}
	return "test", nil
}

func getLinkedStructs(value string, linkedField data.Field) (linkedStructs []interface{}, err error) {
	linkedStructs, ok := linkedField.ValueMap[value]
	if !ok || !(len(linkedStructs) > 0) {
		err = errors.New("No linked struct was found")
	}
	return
}
