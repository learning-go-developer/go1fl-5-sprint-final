package personaldata

import "fmt"

// Personal represents a person's basic information.
// It contains the Name, Weight, and Height fields.
// Name is a string, Weight and Height are stored as float64 values.
type Personal struct {
	Name 	string 	// Person's name
	Weight 	float64 // Person's weight in kilograms
	Height 	float64 // Person's height in meters
}

// Print outputs the Personal struct data to standard output.
// It prints the Name, Weight, and Height fields in a formatted way.
// The weight and height are displayed with two decimal places.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
