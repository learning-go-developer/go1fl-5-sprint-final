package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
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

	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}

	t.Steps = steps
	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w", err)
	}

	if duration <= 0 {
		return errors.New("duration must be positive")
	}

	t.Duration = duration

	return nil
}

// ActionInfo generates and returns a formatted string with details about the training.
// It uses the Training struct fields to calculate:
//   - distance (via spentenergy.Distance),
//   - average speed (via spentenergy.MeanSpeed),
//   - calories burned (via RunningSpentCalories or WalkingSpentCalories).
//
// The returned string has the format:
//
//	Training type: Running
//	Duration: 0.75 h.
//	Distance: 10.00 km.
//	Speed: 13.34 km/h
//	Calories burned: 18621.75
//
// If the training type is unknown, the method returns an error
// with the message "unknown training type". Errors are also returned
// if calculations inside the spentenergy package fail.
func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error

	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("failed to call RunningSpentCalories: %w", err)
		}
	} else if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("failed to call WalkingSpentCalories: %w", err)
		}
	} else {
		return "", errors.New("неизвестный тип тренировки")
	}

	data := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		spentenergy.Distance(t.Steps, t.Height),
		spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration),
		calories,
	)

	return data, nil
}
