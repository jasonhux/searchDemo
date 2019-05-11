package interaction_test

import (
	"searchDemo/src/interaction"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	testCases := map[string]struct {
		userInput             string
		expectedIsQuitCommand bool
		expectedInputValue    string
	}{
		"user type 'quit'": {
			userInput:             "quit",
			expectedIsQuitCommand: true,
			expectedInputValue:    "quit",
		},
		"user type 'test'": {
			userInput:             "test",
			expectedIsQuitCommand: false,
			expectedInputValue:    "test",
		},
	}

	for tc, tp := range testCases {
		interactionService := interaction.NewService(&mockScanner{tp.userInput})
		isQuitCommand, input := interactionService.GetUserInput()
		if isQuitCommand != tp.expectedIsQuitCommand {
			t.Errorf("For test case <%s>, Expected isQuitCommand is: <%v>, but actual value is: <%v>", tc, tp.expectedIsQuitCommand, !tp.expectedIsQuitCommand)
		}
		if input != tp.expectedInputValue {
			t.Errorf("For test case <%s>, Expected input value is: <%s>, but actual value is: <%s>", tc, tp.expectedInputValue, input)
		}
	}
}

type mockScanner struct {
	inputValue string
}

func (s *mockScanner) Scan() bool {
	return true
}
func (s *mockScanner) Text() string {
	return s.inputValue
}
