package search

import (
	"errors"
	"fmt"
	"reflect"
)

type Service interface {
	SetSearchStruct(param string) error
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

func Search(structKey, fieldKey, searchKey string, structMap map[string]map[string]Field) (resultList []interface{}, err error) {
	fieldMap, ok := structMap[structKey]
	if !ok {
		err = errors.New("structKey not found")
		return
	}
	field, ok := fieldMap[fieldKey]
	if !ok {
		err = errors.New("field not found")
		return
	}
	//consider to do a type check for searchKey
	resultList, ok = field.ValueMap[searchKey]
	if !ok {
		err = errors.New("no results found")
	}
	return
}

func (s *service) SetSearchStruct(param string) error {
	_, ok := s.StructMap[param]
	if !ok {
		return errors.New("No struct found")
	}
	s.SelectedStructKey = param
	return nil
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
	return
}
