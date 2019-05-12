package search_test

import (
	"encoding/json"
	"searchDemo/src/data"
	"searchDemo/src/mock"
	"searchDemo/src/search"
	"testing"
)

func TestStartSearch(t *testing.T) {
	testCases := map[string]struct {
		userInputs           []string
		expectedIsQuit       bool
		expectedHasError     bool
		expectedErrorMessage string
		expectedResults      interface{}
	}{
		"user input invalid value while choosing the search type": {
			userInputs:           []string{"4"},
			expectedHasError:     true,
			expectedErrorMessage: "There is no available search type matched to your selection",
		},
		"user input 'quit' while choosing the search type": {
			userInputs:     []string{"quit"},
			expectedIsQuit: true,
		},
		"user input '1' for search type, then 'quit' for search value": {
			userInputs:     []string{"1", "quit"},
			expectedIsQuit: true,
		},
		"user input '1' for search type, then type invalid search value": {
			userInputs:           []string{"1", "abcd"},
			expectedHasError:     true,
			expectedErrorMessage: "No results returned",
		},
		"user input '1' for search type, then type '1' for search value": {
			userInputs: []string{"1", "1"},
			expectedResults: map[string]interface{}{
				"tickets": []data.TicketForDisplay{
					data.TicketForDisplay{
						Ticket: *mock.MockTickets[0], SubmitterName: mock.MockUsers[0].Name, AssigneeName: mock.MockUsers[1].Name, OrganizationName: mock.MockOrganizations[0].Name,
					},
					data.TicketForDisplay{
						Ticket: *mock.MockTickets[1], SubmitterName: mock.MockUsers[1].Name, AssigneeName: mock.MockUsers[0].Name, OrganizationName: mock.MockOrganizations[0].Name,
					},
				},
				"users": []data.UserForDisplay{
					data.UserForDisplay{
						User: *mock.MockUsers[0], OrganizationName: mock.MockOrganizations[0].Name, SubmittedTicketIDs: []string{mock.MockTickets[0].ID}, AssignedTicketsIDs: []string{mock.MockTickets[1].ID},
					},
					data.UserForDisplay{
						User: *mock.MockUsers[1], OrganizationName: mock.MockOrganizations[0].Name, SubmittedTicketIDs: []string{mock.MockTickets[1].ID}, AssignedTicketsIDs: []string{mock.MockTickets[0].ID},
					},
				},
				"organizations": []data.OrganizationForDisplay{
					data.OrganizationForDisplay{
						Organization: *mock.MockOrganizations[0], UserNames: []string{mock.MockUsers[0].Name, mock.MockUsers[1].Name}, TicketIDs: []string{mock.MockTickets[0].ID, mock.MockTickets[1].ID},
					},
				},
			},
		},
		"user input '1' for search type, then type 't2' for search value": {
			userInputs: []string{"1", "t2"},
			expectedResults: map[string]interface{}{
				"tickets": []data.TicketForDisplay{
					data.TicketForDisplay{
						Ticket: *mock.MockTickets[1], SubmitterName: mock.MockUsers[1].Name, AssigneeName: mock.MockUsers[0].Name, OrganizationName: mock.MockOrganizations[0].Name,
					},
				},
			},
		},
		"user input '2' for search type, then type invalid input for struct type selection": {
			userInputs:           []string{"2", "4"},
			expectedHasError:     true,
			expectedErrorMessage: "No struct found",
		},
		"user input '2' for search type, then type 'quit'": {
			userInputs:     []string{"2", "quit"},
			expectedIsQuit: true,
		},
		"user input '2' for search type, then type '1', then type invalid input for field type selection": {
			userInputs:           []string{"2", "1", "abc"},
			expectedHasError:     true,
			expectedErrorMessage: "No field found",
		},
		"user input '2' for search type, then type '1', then type 'quit'": {
			userInputs:     []string{"2", "1", "quit"},
			expectedIsQuit: true,
		},
		"user input '2' for search type, then type '1', then type 'id', then type invalid input for value search": {
			userInputs:           []string{"2", "1", "id", "aaa"},
			expectedHasError:     true,
			expectedErrorMessage: "No results found",
		},
		"user input '2' for search type, then type '1', then type 'id', then type 'quit'": {
			userInputs:     []string{"2", "1", "id", "quit"},
			expectedIsQuit: true,
		},
		"user input '2' for search type, then type '1', then type 'id', then type 't1'": {
			userInputs: []string{"2", "1", "id", "t1"},
			expectedResults: []data.TicketForDisplay{
				data.TicketForDisplay{
					Ticket: *mock.MockTickets[0], SubmitterName: mock.MockUsers[0].Name, AssigneeName: mock.MockUsers[1].Name, OrganizationName: mock.MockOrganizations[0].Name,
				},
			},
		},
		"user input '2' for search type, then type '1', then type 'status', then type 'pending'": {
			userInputs: []string{"2", "1", "status", "pending"},
			expectedResults: []data.TicketForDisplay{
				data.TicketForDisplay{
					Ticket: *mock.MockTickets[0], SubmitterName: mock.MockUsers[0].Name, AssigneeName: mock.MockUsers[1].Name, OrganizationName: mock.MockOrganizations[0].Name,
				},
				data.TicketForDisplay{
					Ticket: *mock.MockTickets[1], SubmitterName: mock.MockUsers[1].Name, AssigneeName: mock.MockUsers[0].Name, OrganizationName: mock.MockOrganizations[0].Name,
				},
			},
		},
		"user input '2' for search type, then type '1', then type 'description', then type ''": {
			userInputs: []string{"2", "1", "description", ""},
			expectedResults: []data.TicketForDisplay{
				data.TicketForDisplay{
					Ticket: *mock.MockTickets[1], SubmitterName: mock.MockUsers[1].Name, AssigneeName: mock.MockUsers[0].Name, OrganizationName: mock.MockOrganizations[0].Name,
				},
			},
		},
		"user input '2' for search type, then type '2', then type 'active', then type 'true'": {
			userInputs: []string{"2", "2", "active", "true"},
			expectedResults: []data.UserForDisplay{
				data.UserForDisplay{
					User: *mock.MockUsers[0], OrganizationName: mock.MockOrganizations[0].Name, SubmittedTicketIDs: []string{mock.MockTickets[0].ID}, AssignedTicketsIDs: []string{mock.MockTickets[1].ID},
				},
				data.UserForDisplay{
					User: *mock.MockUsers[1], OrganizationName: mock.MockOrganizations[0].Name, SubmittedTicketIDs: []string{mock.MockTickets[1].ID}, AssignedTicketsIDs: []string{mock.MockTickets[0].ID},
				},
			},
		},
		"user input '2' for search type, then type '3', then type 'id', then type '1'": {
			userInputs: []string{"2", "3", "id", "1"},
			expectedResults: []data.OrganizationForDisplay{
				data.OrganizationForDisplay{
					Organization: *mock.MockOrganizations[0], UserNames: []string{mock.MockUsers[0].Name, mock.MockUsers[1].Name}, TicketIDs: []string{mock.MockTickets[0].ID, mock.MockTickets[1].ID},
				},
			},
		},
	}
	for tc, tp := range testCases {
		s := search.NewService(&mockDataServiceForSearch{}, &mockInteractionServiceForSearch{userInputs: tp.userInputs, testCase: tc, t: t})
		s.SetStructMap()
		results, isQuit, err := s.StartSearch()
		if isQuit {
			if tp.expectedIsQuit != isQuit {
				t.Errorf("For test case <%s>, Expected isQuite is <%v>, but Actually is <%v>", tc, tp.expectedIsQuit, !tp.expectedIsQuit)
			}
		} else {
			if err != nil {
				if !tp.expectedHasError {
					t.Errorf("For test case <%s>, Expected there is no error returned,  but Actually is", tc)
				}
				if tp.expectedErrorMessage != err.Error() {
					t.Errorf("For test case <%s>, Expected error message is <%s> but Actual message is <%s>", tc, tp.expectedErrorMessage, err.Error())
				}
			} else {
				if tp.expectedHasError {
					t.Errorf("For test case <%s>, Expected there is an error returned,  but Actually is no", tc)
				}
				eb, _ := json.Marshal(tp.expectedResults)
				ab, _ := json.Marshal(results)
				if string(eb) != string(ab) {
					t.Errorf("For test case <%s>, Expected results are <%s>,  but Actually are <%s>", tc, string(eb), string(ab))
				}

			}
		}

	}
}

type mockInteractionServiceForSearch struct {
	userInputs  []string
	calledTimes int
	testCase    string
	t           *testing.T
}

func (s *mockInteractionServiceForSearch) GetUserInput() (isQuitCommand bool, input string) {
	if s.calledTimes > len(s.userInputs) {
		s.t.Errorf("For test case <%s>, the given user inputs length is invalid", s.testCase)
	}
	input = s.userInputs[s.calledTimes]
	if input == "quit" {
		isQuitCommand = true
	}
	s.calledTimes = s.calledTimes + 1
	return
}

type mockDataServiceForSearch struct{}

func (s *mockDataServiceForSearch) PrepareStructMap(tickets []*data.Ticket, users []*data.User, organizations []*data.Organization) (map[string]map[string]data.Field, error) {
	return mock.MockStructMap, nil
}
func (s *mockDataServiceForSearch) LoadFile() (tickets []*data.Ticket, users []*data.User, organizations []*data.Organization, err error) {
	return
}
