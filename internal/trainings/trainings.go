package trainings

import (
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

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
