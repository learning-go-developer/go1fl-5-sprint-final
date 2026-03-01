package daysteps

import (
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
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
	Steps    				int
	Duration 				time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
