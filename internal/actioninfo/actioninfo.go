package actioninfo

// DataParser defines the behavior for types that can parse input
// data and provide a formatted activity summary.
//
// Implementations of this interface must provide:
//   - Parse(): to process and validate input data;
//   - ActionInfo(): to generate a formatted string with details
//     about the activity (e.g., steps, distance, calories).
type DataParser interface {
	Parse()
	ActionInfo()
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
}
