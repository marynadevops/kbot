package events

import (
	"fmt"
	"regexp"
	"time"
)

type EventParser struct{}

func (parser *EventParser) ParseFirstDate(input string) (time.Time, error) {
	// Regular expression pattern to match date formats like "MM/DD/YYYY" or "DD-MM-YYYY"
	datePattern := `(0[1-9]|1[0-2])/(0[1-9]|[12]\d|3[01])/\d{4}|(0[1-9]|[12]\d|3[01])-(0[1-9]|1[0-2])-\d{4}`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(datePattern)

	// Find the first match in the input string
	match := regex.FindString(input)

	if match == "" {
		return time.Time{}, fmt.Errorf("no date found in the input string")
	}

	// Parse the matched date string into a time.Time value
	parsedDate, err := time.Parse("01/02/2006", match)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse the date: %v", err)
	}

	return parsedDate, nil
}
/*
func main() {
	parser := EventParser{}

	// Example usage
	input := "The event will take place on 06/30/2023 at 9:00 AM"
	parsedDate, err := parser.ParseFirstDate(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Date:", parsedDate)
	}
}
*/