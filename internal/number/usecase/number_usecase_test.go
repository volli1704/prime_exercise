package usecase_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/volli1704/prime_api/internal/number/usecase"
	"github.com/volli1704/prime_api/pkg/primechecker"
)

type TestCase struct {
	input       []interface{}
	output      []bool
	expectError bool
}

var ucase = usecase.NewNumberUseCase(primechecker.PrimeChecker{})
var testCaseSet = map[string][]TestCase{
	"simple_normal": {
		TestCase{
			[]interface{}{1, 2, 3, 4},
			[]bool{false, true, true, false},
			false,
		},
	},
	"negative": {
		TestCase{
			[]interface{}{-1, -2, -3, -4},
			[]bool{false, false, false, false},
			false,
		},
	},
	"string": {
		TestCase{
			[]interface{}{"1", 2, "5", "99", "11"},
			[]bool{false, true, true, false, true},
			false,
		},
	},
	"invalid_string_with_float": {
		TestCase{
			[]interface{}{"1", 2, "5.4", "99.12", "11"},
			nil,
			true,
		},
	},
	"valid_float": {
		TestCase{
			[]interface{}{"1.0", 2.0, "5.0", "92", "11.0"},
			[]bool{false, true, true, false, true},
			false,
		},
	},
}

func TestFindPrimesForArray(t *testing.T) {
	for name, testCases := range testCaseSet {
		for _, testCase := range testCases {
			res, err := ucase.FindPrimesForArray(testCase.input)

			if err != nil && !testCase.expectError {
				t.Errorf("%s test failed. Unexpected error %v", name, err)
			}

			if err == nil && testCase.expectError {
				t.Errorf("%s test failed. Expected error but nothing raised", name)
			}

			if !reflect.DeepEqual(res, testCase.output) {
				t.Errorf("%s test failed. Expected %v, got %v", name, testCase.output, res)
			}
		}
	}
}

func TestFindPrimesForArrayMessage(t *testing.T) {
	message := "the given input is invalid. Element on index 3 is not a number"
	res, err := ucase.FindPrimesForArray([]interface{}{"1", 2, "hello", 4, "55"})

	if err == nil {
		t.Errorf("valid_error_message test failed. Expected error, got nil")
	}

	if err.Error() != message {
		t.Errorf("valid_error_message test failed. Expected error message %s, got %s", message, err.Error())
	}

	if res != nil {
		t.Errorf("valid_error_message test failed. Expected nil result, got %v", res)
	}
}

func TestHighLoad(t *testing.T) {
	bigInt := int(math.Pow(2, 64))
	testArr := make([]interface{}, 100001)

	for ix := range testArr {
		testArr[ix] = bigInt - ix
	}

	_, err := ucase.FindPrimesForArray(testArr)
	if err != nil {
		t.Errorf("high_load test failed. Expected nil error, got %v", err)
	}
}
