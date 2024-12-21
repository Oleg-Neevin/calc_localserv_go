package calculation

import (
	"errors"
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "priority+",
			expression:     "2+(2*(2+3))",
			expectedResult: 12,
		},
		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:        "simple",
			expression:  "1+1*",
			expectedErr: ErrInvalidExpression,
		},
		{
			name:        "priority",
			expression:  "(2+2*2",
			expectedErr: ErrInvalidParentheses,
		},
		{
			name:        "priority",
			expression:  "2+2*2)-",
			expectedErr: ErrInvalidParentheses,
		},
		{
			name:        "/",
			expression:  "2/0",
			expectedErr: ErrDivisionByZero,
		},
		{
			name:        "hello world",
			expression:  "hello world 2+2*2",
			expectedErr: ErrInvalidExpression,
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Calc(testCase.expression)
			if !errors.Is(err, testCase.expectedErr) {
				t.Fatalf("successful case %s want error %s, got %s", testCase.expression, testCase.expectedErr, err)
			}
			if val != 0.0 {
				t.Fatalf("%f should be equal 0.0", val)
			}
		})
	}

}
