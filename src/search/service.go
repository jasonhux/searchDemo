package search

import (
	"fmt"
	"reflect"
	"time"
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

func PrepareTickets(tickets []*Ticket) {
	//timenow
	start := time.Now()
	fieldList := SetupStructure(&Ticket{})
	for _, ticket := range tickets {
		v := reflect.ValueOf(ticket).Elem()
		for k := range fieldList {
			fieldValue := fmt.Sprintf("%v", v.FieldByName(k))
			matchedTicketList, _ := fieldList[k].ValueMap[fieldValue]
			updatedTicketList := append(matchedTicketList, ticket)
			fieldList[k].ValueMap[fieldValue] = updatedTicketList
		}
	}
	colapse := time.Now().Sub(start)
	fmt.Println(colapse)
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
