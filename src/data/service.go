package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type Service interface {
	PrepareStructMap(tickets []*Ticket, users []*User, organizations []*Organization) (map[string]map[string]Field, error)
	LoadFile() (tickets []*Ticket, users []*User, organizations []*Organization, err error)
}

type service struct {
}

type Field struct {
	Type         string
	NameWithCase string
	ValueMap     map[string][]interface{}
}

func NewService() Service {
	return &service{}
}

func (s *service) PrepareStructMap(tickets []*Ticket, users []*User, organizations []*Organization) (structMap map[string]map[string]Field, err error) {
	//Convert tickets, users and organizations slice to []interface{} so they can share the same ProcessFieldList func which takes []interface{} as param
	if len(tickets) == 0 {
		err = errors.New("The given tickets data is empty")
		return
	}
	tList := make([]interface{}, len(tickets))
	for i, v := range tickets {
		tList[i] = v
	}

	if len(users) == 0 {
		err = errors.New("The given users data is empty")
		return
	}
	uList := make([]interface{}, len(users))
	for i, v := range users {
		uList[i] = v
	}

	if len(organizations) == 0 {
		err = errors.New("The given organizations data is empty")
		return
	}
	oList := make([]interface{}, len(organizations))
	for i, v := range organizations {
		oList[i] = v
	}
	structMap = map[string]map[string]Field{
		"1": ProcessFieldMap(tList),
		"2": ProcessFieldMap(uList),
		"3": ProcessFieldMap(oList),
	}
	return
}

//ProcessFieldMap func is to convert a struct list such as tickets to a map structure;
//map key is the struct field name with lower case (ex. name). The map value is a Field struct;
//Field struct contains the 1. map key's type (ex. string, int), 2. the key's original case value, (ex. Name) and 3. a nested value map
//The nested value map has the struct field value in string format, with lower case as the key (ex. "a drama in gabon", or "true");
//the map value is a list of pointers to the structs which contains the map key;
//When a field contains an string list, such as 'Tags', treat each string in the list as a separate value map key;
func ProcessFieldMap(structList []interface{}) map[string]Field {
	fieldMap := initFieldMap(structList[0])
	for _, s := range structList {
		v := reflect.ValueOf(s).Elem()
		for k := range fieldMap {
			//Since k is processed as lower case, we need to use fieldNameWithCase for reflect.value.FieldByName func to get the field value
			fieldNameWithCase := fieldMap[k].NameWithCase
			updatedPtrList := []interface{}{}
			switch fieldMap[k].Type {
			case "[]string":
				list, ok := v.FieldByName(fieldNameWithCase).Interface().([]string)
				if ok {
					for _, element := range list {
						//get the point list from ValueMap by the given key.
						//If the key does not exist, it returns an empty slice; so here we don't need to have extra checks to see whether reading key is OK
						matchedPtrList, _ := fieldMap[k].ValueMap[strings.ToLower(element)]
						updatedPtrList = append(matchedPtrList, s)
						fieldMap[k].ValueMap[element] = updatedPtrList
					}
				}
				break
			default:
				fieldValue := strings.ToLower(fmt.Sprintf("%v", v.FieldByName(fieldNameWithCase)))
				matchedPtrList, _ := fieldMap[k].ValueMap[fieldValue]
				updatedPtrList = append(matchedPtrList, s)
				fieldMap[k].ValueMap[fieldValue] = updatedPtrList
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
		//To support case insensitive search, we use field name in lower case as the key of fieldMap.
		//We also save the original case of the field name into the Field struct, this helps when we fill in the values to this initiated FieldMap later;
		fieldNameWithCase := v.Type().Field(i).Name
		n := strings.ToLower(fieldNameWithCase)
		t := f.Type().String()
		fieldMap[n] = Field{Type: t, ValueMap: map[string][]interface{}{}, NameWithCase: fieldNameWithCase}
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
