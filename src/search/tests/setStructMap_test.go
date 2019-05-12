package search_test

import (
	"errors"
	"searchDemo/src/data"
	"searchDemo/src/search"
	"testing"
)

func TestSetStructMap(t *testing.T) {
	testCases := map[string]struct {
		isLoadFileReturnError            bool
		isPrepareStructMapReturnError    bool
		expectedIsPrepareStructMapCalled bool
		expectedHasError                 bool
		expectedErrorMessage             string
	}{
		"Fail to load file": {
			isLoadFileReturnError:            true,
			expectedIsPrepareStructMapCalled: false,
			expectedHasError:                 true,
			expectedErrorMessage:             "error load file",
		},
		"Fail to prepare the struct map": {
			expectedIsPrepareStructMapCalled: true,
			isPrepareStructMapReturnError:    true,
			expectedHasError:                 true,
			expectedErrorMessage:             "error prepare the struct map",
		},
		"Successfully load the struct map": {
			expectedIsPrepareStructMapCalled: true,
			expectedHasError:                 false,
		},
	}
	for tc, tp := range testCases {
		mockDataService := &mockDataService{isLoadFileReturnError: tp.isLoadFileReturnError, isPrepareStructMapReturnError: tp.isPrepareStructMapReturnError}
		s := search.NewService(mockDataService, nil)
		err := s.SetStructMap()
		if err != nil {
			if !tp.expectedHasError {
				t.Errorf("For test case <%s>, Expected there is no error, but actually there is", tc)
			}
			if err.Error() != tp.expectedErrorMessage {
				t.Errorf("For test case <%s>, Expected error message is <%s>, but Actual message is <%s>", tc, tp.expectedErrorMessage, err.Error())
			}
			if mockDataService.IsPrepareStructMapCalled != tp.expectedIsPrepareStructMapCalled {
				t.Errorf("For test case <%s>, Expected dataServie PrepareStructMap func called is <%v>, but Actually is <%v>", tc, tp.expectedIsPrepareStructMapCalled, !tp.expectedIsPrepareStructMapCalled)
			}
		} else {
			savedStructMap := s.GetStructMap()
			if len(savedStructMap) == 0 {
				t.Errorf("For test case <%s>, Expected struct map is saved, but Actually not", tc)
			}
		}
	}
}

type mockDataService struct {
	isLoadFileReturnError         bool
	IsPrepareStructMapCalled      bool
	isPrepareStructMapReturnError bool
}

func (s *mockDataService) LoadFile() (tickets []*data.Ticket, users []*data.User, organizations []*data.Organization, err error) {
	if s.isLoadFileReturnError {
		err = errors.New("error load file")
	}
	return
}

func (s *mockDataService) PrepareStructMap(tickets []*data.Ticket, users []*data.User, organizations []*data.Organization) (map[string]map[string]data.Field, error) {
	s.IsPrepareStructMapCalled = true
	if s.isPrepareStructMapReturnError {
		err := errors.New("error prepare the struct map")
		return nil, err
	}
	return map[string]map[string]data.Field{
		"1": map[string]data.Field{},
	}, nil
}
