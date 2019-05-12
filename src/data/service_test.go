package data_test

import (
	"errors"
	"searchDemo/src/data"
	"strings"
	"testing"
)

func TestFailedLoadFileError(t *testing.T) {
	testCases := map[string]struct {
		failLoadedFileName        string
		failUnMarshaledStructName string
		expectedErrorMessage      string
	}{
		"tickets: read file failed": {
			failLoadedFileName:   "tickets",
			expectedErrorMessage: "read tickets.json file failed",
		},
		"users: read file failed": {
			failLoadedFileName:   "users",
			expectedErrorMessage: "read users.json file failed",
		},
		"organizations: read file failed": {
			failLoadedFileName:   "organizations",
			expectedErrorMessage: "read organizations.json file failed",
		},
		"tickets: unmarshal failed": {
			failUnMarshaledStructName: "tickets",
			expectedErrorMessage:      "unmarshal tickets failed",
		},
		"users: unmarshal failed": {
			failUnMarshaledStructName: "users",
			expectedErrorMessage:      "unmarshal users failed",
		},
		"organizations: unmarshal failed": {
			failUnMarshaledStructName: "organizations",
			expectedErrorMessage:      "unmarshal organizations failed",
		},
	}
	for tc, tp := range testCases {
		mockSerializer := &mockSerializer{failLoadedFileName: tp.failLoadedFileName, failUnMarshaledStructName: tp.failUnMarshaledStructName}
		dataService := data.NewService(mockSerializer)
		_, _, _, err := dataService.LoadFile()
		if err == nil {
			t.Errorf("For test case <%s>, Expected there is an error, but actually not", tc)
		}
		if err.Error() != tp.expectedErrorMessage {
			t.Errorf("For test case <%s>, Expected error message is: <%s>, but actual message is: <%s>", tc, tp.expectedErrorMessage, err.Error())
		}
	}
}

type mockSerializer struct {
	failLoadedFileName        string
	failUnMarshaledStructName string
}

func (s *mockSerializer) ReadFile(filePath string) ([]byte, error) {
	if len(s.failLoadedFileName) != 0 && strings.Contains(filePath, s.failLoadedFileName) {
		return []byte{}, errors.New("read error")
	}
	return []byte{}, nil
}
func (s *mockSerializer) Unmarshal(dataForSerialize []byte, v interface{}) error {
	switch v.(type) {
	case *[]*data.Ticket:
		if s.failUnMarshaledStructName == "tickets" {
			return errors.New("unmarshal error")
		}
	case *[]*data.User:
		if s.failUnMarshaledStructName == "users" {
			return errors.New("unmarshal error")
		}
	case *[]*data.Organization:
		if s.failUnMarshaledStructName == "organizations" {
			return errors.New("unmarshal error")
		}
	}
	return nil
}

func TestPrepareStructMap(t *testing.T) {

	testCases := map[string]struct {
		tickets              []*data.Ticket
		users                []*data.User
		organizations        []*data.Organization
		hasError             bool
		expectedErrorMessage string
		expectedStructMap    map[string]map[string]data.Field
	}{
		"empty ticket list should through an error": {
			tickets:              []*data.Ticket{},
			users:                mockUsers,
			organizations:        mockOrganizations,
			hasError:             true,
			expectedErrorMessage: "The given tickets data is empty",
		},
		"empty users list should through an error": {
			tickets:              mockTickets,
			users:                []*data.User{},
			organizations:        mockOrganizations,
			hasError:             true,
			expectedErrorMessage: "The given users data is empty",
		},
		"empty organizations list should through an error": {
			tickets:              mockTickets,
			users:                mockUsers,
			organizations:        []*data.Organization{},
			hasError:             true,
			expectedErrorMessage: "The given organizations data is empty",
		},
		"valid sources should generate a struct map with all keys lowercase": {
			tickets:       mockTickets,
			users:         mockUsers,
			organizations: mockOrganizations,
			hasError:      false,
			expectedStructMap: map[string]map[string]data.Field{
				"1": map[string]data.Field{
					"id": data.Field{Type: "string", NameWithCase: "ID", ValueMap: map[string][]interface{}{
						"t1": []interface{}{mockTickets[0]},
						"t2": []interface{}{mockTickets[1]},
					}},
					"url": data.Field{Type: "string", NameWithCase: "URL", ValueMap: map[string][]interface{}{
						"http://t1": []interface{}{mockTickets[0]},
						"http://t2": []interface{}{mockTickets[1]},
					}},
					"externalid": data.Field{Type: "string", NameWithCase: "ExternalID", ValueMap: map[string][]interface{}{
						"et1": []interface{}{mockTickets[0]},
						"et2": []interface{}{mockTickets[1]},
					}},
					"createdat": data.Field{Type: "string", NameWithCase: "CreatedAt", ValueMap: map[string][]interface{}{
						"2019-05-11t11:00:01": []interface{}{mockTickets[0]},
						"2019-05-11t11:00:02": []interface{}{mockTickets[1]},
					}},
					"type": data.Field{Type: "string", NameWithCase: "Type", ValueMap: map[string][]interface{}{
						"incident": []interface{}{mockTickets[0], mockTickets[1]},
					}},
					"subject": data.Field{Type: "string", NameWithCase: "Subject", ValueMap: map[string][]interface{}{
						"test1": []interface{}{mockTickets[0]},
						"test2": []interface{}{mockTickets[1]},
					}},
					"description": data.Field{Type: "string", NameWithCase: "Description", ValueMap: map[string][]interface{}{
						"test description": []interface{}{mockTickets[0]},
						"":                 []interface{}{mockTickets[1]},
					}},
					"priority": data.Field{Type: "string", NameWithCase: "Priority", ValueMap: map[string][]interface{}{
						"high": []interface{}{mockTickets[0], mockTickets[1]},
					}},
					"status": data.Field{Type: "string", NameWithCase: "Status", ValueMap: map[string][]interface{}{
						"pending": []interface{}{mockTickets[0], mockTickets[1]},
					}},
					"submitterid": data.Field{Type: "int", NameWithCase: "SubmitterID", ValueMap: map[string][]interface{}{
						"1": []interface{}{mockTickets[0]},
						"2": []interface{}{mockTickets[1]},
					}},
					"assigneeid": data.Field{Type: "int", NameWithCase: "AssigneeID", ValueMap: map[string][]interface{}{
						"1": []interface{}{mockTickets[1]},
						"2": []interface{}{mockTickets[0]},
					}},
					"organizationid": data.Field{Type: "int", NameWithCase: "OrganizationID", ValueMap: map[string][]interface{}{
						"1": []interface{}{mockTickets[0], mockTickets[1]},
					}},
					"tags": data.Field{Type: "[]string", NameWithCase: "Tags", ValueMap: map[string][]interface{}{
						"tag1.1": []interface{}{mockTickets[0]},
						"tag1.2": []interface{}{mockTickets[0]},
						"tag2.1": []interface{}{mockTickets[1]},
						"tag2.2": []interface{}{mockTickets[1]},
					}},
					"hasincidents": data.Field{Type: "bool", NameWithCase: "HasIncidents", ValueMap: map[string][]interface{}{
						"false": []interface{}{mockTickets[0], mockTickets[1]},
					}},
					"dueat": data.Field{Type: "string", NameWithCase: "DueAt", ValueMap: map[string][]interface{}{
						"2019-05-13t11:00:01": []interface{}{mockTickets[0]},
						"2019-05-13t11:00:02": []interface{}{mockTickets[1]},
					}},
					"via": data.Field{Type: "string", NameWithCase: "Via", ValueMap: map[string][]interface{}{
						"web": []interface{}{mockTickets[0], mockTickets[1]},
					}},
				},
				"2": map[string]data.Field{
					"id": data.Field{Type: "int", NameWithCase: "ID", ValueMap: map[string][]interface{}{
						"1": []interface{}{mockUsers[0]},
						"2": []interface{}{mockUsers[1]},
					}},
					"url": data.Field{Type: "string", NameWithCase: "URL", ValueMap: map[string][]interface{}{
						"http://u1": []interface{}{mockUsers[0]},
						"http://u2": []interface{}{mockUsers[1]},
					}},
					"externalid": data.Field{Type: "string", NameWithCase: "ExternalID", ValueMap: map[string][]interface{}{
						"u1": []interface{}{mockUsers[0]},
						"u2": []interface{}{mockUsers[1]},
					}},
					"name": data.Field{Type: "string", NameWithCase: "Name", ValueMap: map[string][]interface{}{
						"test testa": []interface{}{mockUsers[0]},
						"test testb": []interface{}{mockUsers[1]},
					}},
					"alias": data.Field{Type: "string", NameWithCase: "Alias", ValueMap: map[string][]interface{}{
						"user 1": []interface{}{mockUsers[0]},
						"user 2": []interface{}{mockUsers[1]},
					}},
					"createdat": data.Field{Type: "string", NameWithCase: "CreatedAt", ValueMap: map[string][]interface{}{
						"2019-04-11t11:00:01": []interface{}{mockUsers[0]},
						"2019-04-11t11:00:02": []interface{}{mockUsers[1]},
					}},
					"active": data.Field{Type: "bool", NameWithCase: "Active", ValueMap: map[string][]interface{}{
						"true": []interface{}{mockUsers[0], mockUsers[1]},
					}},
					"verified": data.Field{Type: "bool", NameWithCase: "Verified", ValueMap: map[string][]interface{}{
						"true": []interface{}{mockUsers[0], mockUsers[1]},
					}},
					"shared": data.Field{Type: "bool", NameWithCase: "Shared", ValueMap: map[string][]interface{}{
						"true": []interface{}{mockUsers[0], mockUsers[1]},
					}},
					"locale": data.Field{Type: "string", NameWithCase: "Locale", ValueMap: map[string][]interface{}{
						"en-au": []interface{}{mockUsers[0], mockUsers[1]},
					}},
					"timezone": data.Field{Type: "string", NameWithCase: "TimeZone", ValueMap: map[string][]interface{}{
						"australia": []interface{}{mockUsers[0], mockUsers[1]},
					}},
					"lastloginat": data.Field{Type: "string", NameWithCase: "LastLoginAt", ValueMap: map[string][]interface{}{
						"2019-05-12t11:00:01": []interface{}{mockUsers[0]},
						"2019-05-12t11:00:02": []interface{}{mockUsers[1]},
					}},
					"email": data.Field{Type: "string", NameWithCase: "Email", ValueMap: map[string][]interface{}{
						"user1@test.com": []interface{}{mockUsers[0]},
						"user2@test.com": []interface{}{mockUsers[1]},
					}},
					"phone": data.Field{Type: "string", NameWithCase: "Phone", ValueMap: map[string][]interface{}{
						"9991": []interface{}{mockUsers[0]},
						"9992": []interface{}{mockUsers[1]},
					}},
					"signature": data.Field{Type: "string", NameWithCase: "Signature", ValueMap: map[string][]interface{}{
						"":                []interface{}{mockUsers[0]},
						"user signature2": []interface{}{mockUsers[1]},
					}},
					"organizationid": data.Field{Type: "int", NameWithCase: "OrganizationID", ValueMap: map[string][]interface{}{
						"1": []interface{}{mockUsers[0], mockUsers[1]},
					}},
					"tags": data.Field{Type: "[]string", NameWithCase: "Tags", ValueMap: map[string][]interface{}{
						"utag1.1": []interface{}{mockUsers[0]},
						"utag1.2": []interface{}{mockUsers[0]},
						"utag2.1": []interface{}{mockUsers[1]},
						"utag2.2": []interface{}{mockUsers[1]},
					}},
					"suspended": data.Field{Type: "bool", NameWithCase: "Suspended", ValueMap: map[string][]interface{}{
						"false": []interface{}{mockUsers[0], mockUsers[1]},
					}},
					"role": data.Field{Type: "string", NameWithCase: "Role", ValueMap: map[string][]interface{}{
						"admin": []interface{}{mockUsers[0]},
						"user":  []interface{}{mockUsers[1]},
					}},
				},
				"3": map[string]data.Field{
					"id": data.Field{Type: "int", NameWithCase: "ID", ValueMap: map[string][]interface{}{
						"1": []interface{}{mockOrganizations[0]},
					}},
					"url": data.Field{Type: "string", NameWithCase: "URL", ValueMap: map[string][]interface{}{
						"http://org1": []interface{}{mockOrganizations[0]},
					}},
					"externalid": data.Field{Type: "string", NameWithCase: "ExternalID", ValueMap: map[string][]interface{}{
						"o1": []interface{}{mockOrganizations[0]},
					}},
					"name": data.Field{Type: "string", NameWithCase: "Name", ValueMap: map[string][]interface{}{
						"test org1": []interface{}{mockOrganizations[0]},
					}},
					"domainnames": data.Field{Type: "[]string", NameWithCase: "DomainNames", ValueMap: map[string][]interface{}{
						"org1.1.com": []interface{}{mockOrganizations[0]},
						"org1.2.com": []interface{}{mockOrganizations[0]},
					}},
					"createdat": data.Field{Type: "string", NameWithCase: "CreatedAt", ValueMap: map[string][]interface{}{
						"2019-05-01t11:00:00": []interface{}{mockOrganizations[0]},
					}},
					"details": data.Field{Type: "string", NameWithCase: "Details", ValueMap: map[string][]interface{}{
						"details1": []interface{}{mockOrganizations[0]},
					}},
					"sharedtickets": data.Field{Type: "bool", NameWithCase: "SharedTickets", ValueMap: map[string][]interface{}{
						"false": []interface{}{mockOrganizations[0]},
					}},
					"tags": data.Field{Type: "[]string", NameWithCase: "Tags", ValueMap: map[string][]interface{}{
						"otag1.1": []interface{}{mockOrganizations[0]},
						"otag1.2": []interface{}{mockOrganizations[0]},
					}},
				},
			},
		},
	}
	for tc, tp := range testCases {
		mockSerializer := &mockSerializer{}
		dataService := data.NewService(mockSerializer)
		structMap, err := dataService.PrepareStructMap(tp.tickets, tp.users, tp.organizations)
		if tp.hasError {
			if err == nil {
				t.Errorf("For test case <%s>, Expected error returned but Actually not", tc)
			}
			if err.Error() != tp.expectedErrorMessage {
				t.Errorf("For test case <%s>, Expected error message is <%s> but Actual message is <%s>", tc, tp.expectedErrorMessage, err.Error())
			}
		} else {
			if err != nil {
				t.Errorf("For test case <%s>, Expected no error returned but Actually there is", tc)
			}
			if len(tp.expectedStructMap) != len(structMap) {
				t.Errorf("For test case <%s>, Expected struct map has <%v> keys but Actual map has <%v>", tc, len(tp.expectedStructMap), len(structMap))
			}
			for k := range tp.expectedStructMap {
				actualFieldMap, ok := structMap[k]
				if !ok {
					t.Errorf("For test case <%s>, Expected struct map key: <%s> is not a valid key in actual map", tc, k)
				}
				expectedFieldMap := tp.expectedStructMap[k]
				if len(expectedFieldMap) != len(actualFieldMap) {
					t.Errorf("For test case <%s>, the expected fieldMap of structMap key: <%s> has <%v> keys but actual fieldMap has <%v> keys", tc, k, len(expectedFieldMap), len(actualFieldMap))
				}
				for fk := range expectedFieldMap {
					actualField, ok := actualFieldMap[fk]
					if !ok {
						t.Errorf("For test case <%s>, Expected field map key: <%s> is not a valid key in actual field map", tc, fk)
					}
					expectedField := expectedFieldMap[fk]
					if expectedField.NameWithCase != actualField.NameWithCase {
						t.Errorf("For test case <%s>, for struct key <%s>, and field key <%s>, Expected field's nameWithCase is <%s> but Actual field's nameWithCase is <%s>", tc, k, fk, expectedField.NameWithCase, actualField.NameWithCase)
					}
					if expectedField.Type != actualField.Type {
						t.Errorf("For test case <%s>, for struct key <%s>, and field key <%s>, Expected field's type is <%s> but Actual field's type is <%s>", tc, k, fk, expectedField.Type, actualField.Type)
					}
					if len(expectedField.ValueMap) != len(actualField.ValueMap) {
						t.Errorf("For test case <%s>, for struct key <%s>, and field key <%s>, Expected field's value map has <%v> keys but Actual field has <%v>", tc, k, fk, len(expectedField.ValueMap), len(actualField.ValueMap))
					}
					actualValueMap := actualField.ValueMap
					for vk := range expectedField.ValueMap {
						actualPtrList, ok := actualValueMap[vk]
						if !ok {
							t.Errorf("For test case <%s>, for struct key <%s>, and field key <%s>, Expected field's value map has key <%s> but Actual map does not", tc, k, fk, vk)
						}
						expectedPtrList := expectedField.ValueMap[vk]
						if len(expectedPtrList) != len(actualPtrList) {
							t.Errorf("For test case <%s>, for struct key <%s>, field key <%s>, and value key <%s>, Expected value length is <%v> but actual value's length is <%v>", tc, k, fk, vk, len(expectedPtrList), len(actualPtrList))
						}
						for _, ptr := range expectedPtrList {
							isExistInActualList := false
							for _, p := range actualPtrList {
								if ptr == p {
									isExistInActualList = true
									break
								}
							}
							if !isExistInActualList {
								t.Errorf("For test case <%s>, for struct key <%s>, field key <%s>, and value key <%s>, Expected value ptr list contains a ptr not in Actual pr list", tc, k, fk, vk)
							}
						}
					}
				}
			}
		}

	}
}

var (
	mockTickets = []*data.Ticket{
		&data.Ticket{
			ID:             "t1",
			URL:            "http://t1",
			ExternalID:     "et1",
			CreatedAt:      "2019-05-11T11:00:01",
			Type:           "incident",
			Subject:        "Test1",
			Description:    "Test description",
			Priority:       "high",
			Status:         "pending",
			SubmitterID:    1,
			AssigneeID:     2,
			OrganizationID: 1,
			Tags:           []string{"Tag1.1", "Tag1.2"},
			HasIncidents:   false,
			DueAt:          "2019-05-13T11:00:01",
			Via:            "web",
		},
		&data.Ticket{
			ID:             "t2",
			URL:            "http://t2",
			ExternalID:     "et2",
			CreatedAt:      "2019-05-11T11:00:02",
			Type:           "incident",
			Subject:        "test2",
			Description:    "",
			Priority:       "high",
			Status:         "pending",
			SubmitterID:    2,
			AssigneeID:     1,
			OrganizationID: 1,
			Tags:           []string{"Tag2.1", "Tag2.2"},
			HasIncidents:   false,
			DueAt:          "2019-05-13T11:00:02",
			Via:            "web",
		},
	}
	mockUsers = []*data.User{
		&data.User{
			ID:             1,
			URL:            "http://u1",
			ExternalID:     "u1",
			Name:           "Test TestA",
			Alias:          "user 1",
			CreatedAt:      "2019-04-11T11:00:01",
			Active:         true,
			Verified:       true,
			Shared:         true,
			Locale:         "en-AU",
			TimeZone:       "Australia",
			LastLoginAt:    "2019-05-12T11:00:01",
			Email:          "user1@test.com",
			Phone:          "9991",
			Signature:      "",
			OrganizationID: 1,
			Tags:           []string{"utag1.1", "utag1.2"},
			Suspended:      false,
			Role:           "admin",
		},
		&data.User{
			ID:             2,
			URL:            "http://u2",
			ExternalID:     "u2",
			Name:           "Test TestB",
			Alias:          "user 2",
			CreatedAt:      "2019-04-11T11:00:02",
			Active:         true,
			Verified:       true,
			Shared:         true,
			Locale:         "en-AU",
			TimeZone:       "Australia",
			LastLoginAt:    "2019-05-12T11:00:02",
			Email:          "user2@test.com",
			Phone:          "9992",
			Signature:      "user signature2",
			OrganizationID: 1,
			Tags:           []string{"utag2.1", "utag2.2"},
			Suspended:      false,
			Role:           "user",
		},
	}
	mockOrganizations = []*data.Organization{
		&data.Organization{
			ID:            1,
			URL:           "http://org1",
			ExternalID:    "o1",
			Name:          "test org1",
			DomainNames:   []string{"org1.1.com", "org1.2.com"},
			CreatedAt:     "2019-05-01T11:00:00",
			Details:       "details1",
			SharedTickets: false,
			Tags:          []string{"otag1.1", "otag1.2"},
		},
	}
)
