package search_test

import (
	"searchDemo/src/search"
	"testing"
)

func TestRequestNewSearch(t *testing.T) {
	testCases := map[string]struct {
		input                      string
		expectedIsRequestNewSearch bool
	}{
		"user input quit": {
			input: "quit",
			expectedIsRequestNewSearch: false,
		},
		"user input n": {
			input: "n",
			expectedIsRequestNewSearch: false,
		},
		"user input N": {
			input: "N",
			expectedIsRequestNewSearch: true,
		},
	}
	for tc, tp := range testCases {
		s := search.NewService(nil, &mockInteractionService{tp.input})
		isRequestForNewSearch := s.RequestNewSearch()
		if tp.expectedIsRequestNewSearch != isRequestForNewSearch {
			t.Errorf("For test case <%s>, Expected isRequestForNewSearch is <%v> but Actual is <%v>", tc, tp.expectedIsRequestNewSearch, !tp.expectedIsRequestNewSearch)
		}
	}
}

type mockInteractionService struct {
	userInput string
}

func (s *mockInteractionService) GetUserInput() (isQuitCommand bool, input string) {
	input = s.userInput
	if input == "quit" {
		isQuitCommand = true
	}
	return
}
