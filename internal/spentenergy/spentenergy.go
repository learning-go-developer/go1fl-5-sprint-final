package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories calculates the estimated calories spent while walking
// based on the number of steps, the user's weight and height, and the duration.
// It returns an error if steps is zero or duration is non-positive.
// Internally, it calls RunningSpentCalories() to compute the base calorie value,
// then multiplies the result by walkingCaloriesCoefficient to apply the walking
// correction factor.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps == 0 {
		return 0, errors.New("steps must be greater than 0")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate calories: %w", err)
	}

	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories calculates the estimated calories spent while running
// based on the number of steps, the runner's weight and height, and the duration.
// It returns an error if steps is zero or duration is non-positive.
// The calculation uses MeanSpeed() to determine average speed and multiplies
// it by the runner's weight and the duration in minutes, normalized by minInH.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps == 0 {
		return 0, errors.New("steps must be greater than 0")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

// MeanSpeed calculates the average speed in kilometers per hour
// based on the number of steps, the person's height, and the duration.
// If steps is zero or duration is non-positive, it returns 0.
// The function uses Distance() to compute the traveled distance
// and divides it by the duration in hours.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps == 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// Distance calculates the traveled distance in kilometers
// based on the number of steps and the person's height.
// It uses a step length coefficient and converts meters to kilometers.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength

	return distanceMeters / mInKm
}
