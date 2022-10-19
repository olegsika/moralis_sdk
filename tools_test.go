package moralissdk

import (
	"errors"
	"reflect"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type TestStruct struct {
		Name string
	}
	var A = TestStruct{Name: "A"}

	tCases := []struct {
		name           string
		entryStruct    interface{}
		expectedResult map[string]interface{}
		expectedErr    error
	}{
		{
			name:           "Should be success",
			entryStruct:    A,
			expectedResult: map[string]interface{}{"Name": "A"},
		},
		{
			name:        "Should be error. Entry not a struct",
			entryStruct: 1,
			expectedErr: errors.New("json: cannot unmarshal number into Go value of type map[string]interface {}"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			res, err := StructToMap(tc.entryStruct)

			if !IsErrorEqual(err, tc.expectedErr) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
				return
			}

			if !reflect.DeepEqual(res, tc.expectedResult) {
				tt.Errorf("expected result = %v but actual result = %v", tc.expectedResult, res)
				return
			}
		})
	}
}

func TestInListString(t *testing.T) {
	tCases := []struct {
		name           string
		expectedResult bool
		entry          string
		list           []string
	}{
		{
			name:           "Test success",
			expectedResult: true,
			entry:          "test",
			list:           []string{"test", "test_2"},
		},
		{
			name:           "Test not success",
			expectedResult: false,
			entry:          "test_1",
			list:           []string{"test", "test_2"},
		},
		{
			name:           "Test empty entry",
			expectedResult: false,
			entry:          "",
			list:           []string{"test", "test_2"},
		},
		{
			name:           "Test empty list",
			expectedResult: false,
			entry:          "",
			list:           []string{},
		},
		{
			name:           "Test empty list and not empty entry",
			expectedResult: false,
			entry:          "test",
			list:           []string{},
		},
		{
			name:           "Test list nil",
			expectedResult: false,
			entry:          "",
			list:           nil,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			res := InListString(tc.entry, tc.list)

			if res != tc.expectedResult {
				tt.Errorf("expected result %v, but actual result %v", tc.expectedResult, res)
				return
			}
		})
	}
}

func TestIsErrorEqual(t *testing.T) {
	var (
		aErr = errors.New("test")
		bErr = aErr
	)

	tCases := []struct {
		name           string
		firstErr       error
		secondErr      error
		expectedResult bool
	}{
		{
			name:           "should be true",
			firstErr:       aErr,
			secondErr:      bErr,
			expectedResult: true,
		},
		{
			name:           "should be false",
			firstErr:       aErr,
			secondErr:      errors.New("qweqweqweqwe"),
			expectedResult: false,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			res := IsErrorEqual(tc.firstErr, tc.secondErr)
			if res != tc.expectedResult {
				tt.Errorf("expected result = %v but actual result = %v", tc.expectedResult, res)
			}
		})
	}
}
