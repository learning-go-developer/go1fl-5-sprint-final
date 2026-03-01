package actioninfo

import (
	"log"
)

// DataParser defines the behavior for types that can parse input
// data and provide a formatted activity summary.
//
// Implementations of this interface must provide:
//   - Parse(): to process and validate input data;
//   - ActionInfo(): to generate a formatted string with details
//     about the activity (e.g., steps, distance, calories).
type DataParser interface {
	Parse(input string) error
	ActionInfo() (string, error)
}

// Info processes a dataset of input strings using a DataParser.
//
// For each element in the dataset, the function:
//   - Calls Parse on the DataParser to process the input string;
//   - Logs an error if parsing fails;
//   - Calls ActionInfo on the DataParser to generate a workout summary;
//   - Logs an error if ActionInfo fails.
//
// The function does not return a value; it relies on logging to report
// errors and output summaries.
func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("failed to call Parse: %v", err)
		}

		dp.ActionInfo()
		if err != nil {
			log.Printf("failed to call WalkingSpentCalories: %v", err)
		}
	}
}
