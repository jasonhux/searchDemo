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

func (s *service) PrepareData() {

}

func PrepareTickets(tickets []*Ticket) map[string]Field {
	//timenow
	fmt.Println(len(tickets))
	//start := time.Now()
	fieldList := SetupStructure(&Ticket{})
	for _, ticket := range tickets {
		v := reflect.ValueOf(ticket).Elem()
		for k := range fieldList {
			updatedTicketList := []interface{}{}
			//make sure array change and bool type change is required
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
	// colapse := time.Now().Sub(start)
	// fmt.Println(colapse)
	return fieldList
}

func SetupStructure(instance interface{}) map[string]Field {
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

func SearchTicket(key, searchKey string, fieldList map[string]Field) (resultList []interface{}, err error) {
	field, ok := fieldList[key]
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
