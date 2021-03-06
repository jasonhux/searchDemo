package data_test

import (
	"errors"
	"searchDemo/src/data"
	"searchDemo/src/mock"
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
			users:                mock.MockUsers,
			organizations:        mock.MockOrganizations,
			hasError:             true,
			expectedErrorMessage: "The given tickets data is empty",
		},
		"empty users list should through an error": {
			tickets:              mock.MockTickets,
			users:                []*data.User{},
			organizations:        mock.MockOrganizations,
			hasError:             true,
			expectedErrorMessage: "The given users data is empty",
		},
		"empty organizations list should through an error": {
			tickets:              mock.MockTickets,
			users:                mock.MockUsers,
			organizations:        []*data.Organization{},
			hasError:             true,
			expectedErrorMessage: "The given organizations data is empty",
		},
		"valid sources should generate a struct map with all keys lowercase": {
			tickets:           mock.MockTickets,
			users:             mock.MockUsers,
			organizations:     mock.MockOrganizations,
			hasError:          false,
			expectedStructMap: mock.MockStructMap,
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
