package search

import (
	"errors"
	"fmt"
	"reflect"
)

type Service interface {
}

type service struct {
	Tickets []*Ticket
	//other two

}

type Field struct {
	Type     string
	ValueMap map[string][]interface{}
}

func NewService() Service {
	return &service{}
}

func PrepareData(tickets []*Ticket, users []*User, organizations []*Organization) map[string]map[string]Field {
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
		"1": ProcessFieldList(tList),
		"2": ProcessFieldList(uList),
		"3": ProcessFieldList(oList),
	}
}

//ProcessFieldList func is to convert object list such as tickets to a map structure; map key is the struct field name (ex. Name). The map value is the Field struct;
//Field struct contains the 1. map key's type (ex. string, int) and 2. a nested value map;
//The nested value map has the struct field value in string format as the key (ex. "A Drama in Gabon", or "true"); the map value are a list of pointers to the structs which contains the map key;
//When a field contains an array, treat each string in the array as a separate value map key;
func ProcessFieldList(structList []interface{}) map[string]Field {
	fmt.Println(len(structList))

	fieldList := initFieldList(structList[0])
	for _, ticket := range structList {
		v := reflect.ValueOf(ticket).Elem()
		for k := range fieldList {
			updatedTicketList := []interface{}{}
			switch fieldList[k].Type {
			case "[]string":
				list, ok := v.FieldByName(k).Interface().([]string)
				if ok {
					for _, element := range list {
						matchedTicketList, _ := fieldList[k].ValueMap[element]
						updatedTicketList = append(matchedTicketList, ticket)
						fieldList[k].ValueMap[element] = updatedTicketList
					}
				}
				break
			default:
				fieldValue := fmt.Sprintf("%v", v.FieldByName(k))
				matchedTicketList, _ := fieldList[k].ValueMap[fieldValue]
				updatedTicketList = append(matchedTicketList, ticket)
				fieldList[k].ValueMap[fieldValue] = updatedTicketList
			}
		}
	}
	return fieldList
}

func initFieldList(instance interface{}) map[string]Field {
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

func SearchTicket(structKey, fieldKey, searchKey string, fieldListMap map[string]map[string]Field) (resultList []interface{}, err error) {
	fieldList, ok := fieldListMap[structKey]
	if !ok {
		err = errors.New("structKey not found")
		return
	}
	field, ok := fieldList[fieldKey]
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
