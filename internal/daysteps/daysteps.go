package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps represents a record of daily physical activity.
// It stores the number of steps taken, the total duration of activity,
// and personal information about the user.
//
// Fields:
// - Steps: the number of steps completed during the day.
// - Duration: the total time spent being physically active.
// - Personal: embedded personal data of the user (e.g., name, age).
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse parses a comma-separated string into the DaySteps fields.
//
// The expected format of datastring is "steps,duration".
// Example: "1200,2h30m"
//
// Behavior:
//   - Splits the input string by comma.
//   - Converts the first part into an integer (Steps).
//   - Validates that steps > 0.
//   - Parses the second part into a time.Duration (Duration).
//
// Returns an error if:
//   - The input does not contain exactly two parts.
//   - Steps cannot be converted to an integer.
//   - Steps are zero.
//   - Duration cannot be parsed.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("expected 2 parts, got different length")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to conveert steps: %w", err)
	}

	if steps == 0 {
		return errors.New("steps must be greater than 0")
	}

	ds.Steps = steps

	d, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w", err)
	}

	ds.Duration = d

	return nil
}

// ActionInfo generates and returns a formatted string with details
// about the walking activity.
//
// The method uses the DaySteps fields to calculate:
//   - total number of steps;
//   - distance covered (based on height);
//   - calories burned (via WalkingSpentCalories).
//
// It returns a string in the following format:
//
//	Steps: 792
//	Distance: 0.51 km
//	Calories burned: 221.33 kcal
//
// If an error occurs during the calculation (e.g., in WalkingSpentCalories),
// the method returns an empty string along with the error.
func (ds DaySteps) ActionInfo() (string, error) {
	var calories float64
	var err error

	calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("failed to call WalkingSpentCalories: %w", err)
	}

	data := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		spentenergy.Distance(ds.Steps, ds.Height),
		calories,
	)

	return data, nil
}
