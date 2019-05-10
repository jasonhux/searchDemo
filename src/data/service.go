package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
)

type Service interface {
	PrepareStructMap(tickets []*Ticket, users []*User, organizations []*Organization) (map[string]map[string]Field, error)
	LoadFile() (tickets []*Ticket, users []*User, organizations []*Organization, err error)
}

type service struct {
}

type Field struct {
	Type     string
	ValueMap map[string][]interface{}
}

func NewService() Service {
	return &service{}
}

func (s *service) PrepareStructMap(tickets []*Ticket, users []*User, organizations []*Organization) (map[string]map[string]Field, error) {
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
	}, nil
	//dont forget to add error handling
}

//ProcessFieldMap func is to convert object list such as tickets to a map structure; map key is the struct field name (ex. Name). The map value is the Field struct;
//Field struct contains the 1. map key's type (ex. string, int) and 2. a nested value map;
//The nested value map has the struct field value in string format as the key (ex. "A Drama in Gabon", or "true"); the map value are a list of pointers to the structs which contains the map key;
//When a field contains an array, treat each string in the array as a separate value map key;
func ProcessFieldMap(structList []interface{}) map[string]Field {
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

func (s *service) LoadFile() (tickets []*Ticket, users []*User, organizations []*Organization, err error) {
	data, e := ioutil.ReadFile("./data/tickets.json")
	if e != nil {
		err = errors.New("read tickets.json file failed")
		return
	}
	e = json.Unmarshal(data, &tickets)
	if e != nil {
		err = errors.New("unmarshal tickets failed")
		return
	}
	data, e = ioutil.ReadFile("./data/users.json")
	if e != nil {
		err = errors.New("read users.json file failed")
		return
	}
	e = json.Unmarshal(data, &users)
	if e != nil {
		err = errors.New("unmarshal users failed")
		return
	}
	data, e = ioutil.ReadFile("./data/organizations.json")
	if e != nil {
		err = errors.New("read organizations.json file failed")
		return
	}
	e = json.Unmarshal(data, &organizations)
	if e != nil {
		err = errors.New("unmarshal organizations failed")
	}
	return
}
