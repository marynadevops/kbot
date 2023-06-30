package events

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseFirstDate(t *testing.T) {
	parser := EventParser{}

	// Test cases
	testCases := []struct {
		input  string
		output string
		err    error
	}{
		// Valid date case
		{
			input:  "Event on 06/30/2023 at 10:00 AM",
			output: "06/30/2023",
			err:    nil,
		},
		// Valid date case
		{
			input:  "Meeting at 2023-06-16 14:30",
			output: "2023-06-16",
			err:    nil,
		},
		// Invalid date case
		{
			input:  "Invalid date format",
			output: "",
			err:    errors.New("Invalid date format"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			// Call the ParseFirstDate function with the input
			result, err := parser.ParseFirstDate(tc.input)
			// parsedDate, err := parser.ParseFirstDate(testCase.input)

			// Check if the error matches the expected error
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error(), "Unexpected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}

			// Check if the output matches the expected output
			output, _ := time.Parse("01/02/2006", tc.output)
			assert.Equal(t, output, result, "Unexpected date")
		})
	}
}

/*
package events

import (
	// "fmt"
	// "regexp"
	"time"

	"testing"
)

func TestEventParser_ParseFirstDate(t *testing.T) {
	parser := EventParser{}

	testCases := []struct {
		input         string
		expectedDate  string
		expectedError bool
	}{
		{"The event will take place on 06/30/2023 at 9:00 AM", "06/30/2023", false},
		// {"Please join us on 12-15-2022 for the conference.", "12-15-2022", false},
		{"No date in this input", "", true},
		{"Invalid date format: 2023/06/30", "", true},
	}

	for _, testCase := range testCases {
		parsedDate, err := parser.ParseFirstDate(testCase.input)

		if testCase.expectedError {
			if err == nil {
				t.Errorf("Expected an error, but got nil for input: %s", testCase.input)
			} else {
				t.Log("Test case passed!")
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error, but got: %v for input: %s", err, testCase.input)
			} else {
				expectedDate, _ := time.Parse("01/02/2006", testCase.expectedDate)
				if !parsedDate.Equal(expectedDate) {
					t.Errorf("Expected date: %s, but got: %s for input: %s", expectedDate, parsedDate, testCase.input)
				}
			}
		}
	}
}
*/
