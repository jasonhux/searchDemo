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
}

type mockSerializer struct {
	failLoadedFileName        string
	failUnMarshaledStructName string
}

func (s *mockSerializer) ReadFile(filePath string) ([]byte, error) {
	if strings.Contains(filePath, s.failLoadedFileName) {
		return []byte{}, errors.New("read error")
	}
	return []byte{}, nil
}
func (s *mockSerializer) Unmarshal(dataForSerialize []byte, v interface{}) error {
	switch v.(type) {
	case []*data.Ticket:
		if s.failUnMarshaledStructName == "tickets" {
			return errors.New("unmarshal error")
		}
	case []*data.User:
		if s.failUnMarshaledStructName == "users" {
			return errors.New("unmarshal error")
		}
	case []*data.Organization:
		if s.failUnMarshaledStructName == "organizations" {
			return errors.New("unmarshal error")
		}
	}
	return nil
}
