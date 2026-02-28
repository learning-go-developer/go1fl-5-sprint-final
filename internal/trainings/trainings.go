package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

// Training represents a workout session.
// It contains the number of steps, the type of training (running or walking),
// the duration of the workout, and embedded personal data about the user.
// By embedding personaldata.Personal, Training also includes the user's
// name, weight, height, and inherits the Print() method for displaying
// this information.
type Training struct {
	Steps                 int           // number of steps during training
	TrainingType          string        // type of training: "running" or "walking"
	Duration              time.Duration // duration of the workout
	personaldata.Personal               // embedded struct with Name, Weight, Height
}

// Parse takes a comma-separated string and fills the Training fields.
// The expected format is "steps,trainingType,duration".
// - steps: integer greater than 0
// - trainingType: string (e.g., "run" or "walk")
// - duration: valid time.Duration string (e.g., "30m", "1h15m")
// If the input is invalid or cannot be converted, Parse returns an error.
// On success, it sets the Steps, TrainingType, and Duration fields of Training.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("expected 3 parts, got different length")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to conveert steps: %w", err)
	}

	if steps == 0 {
		return errors.New("steps must be greater than 0")
	}

	t.Steps = steps
	t.TrainingType = parts[1]

	d, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w", err)
	}

	t.Duration = d

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
