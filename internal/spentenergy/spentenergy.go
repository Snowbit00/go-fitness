package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Ошибки проверки входных параметров.
var (
	ErrInvalidSteps    = errors.New("количество шагов должно быть больше нуля")
	ErrInvalidWeight   = errors.New("вес должен быть больше нуля")
	ErrInvalidHeight   = errors.New("рост должен быть больше нуля")
	ErrInvalidDuration = errors.New("продолжительность должна быть больше нуля")
)

// validate проверяет, что все параметры активности положительные.
func validate(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return ErrInvalidSteps
	case weight <= 0:
		return ErrInvalidWeight
	case height <= 0:
		return ErrInvalidHeight
	case duration <= 0:
		return ErrInvalidDuration
	}
	return nil
}

// spentCalories считает калории по формуле (вес * средняя скорость * минуты) / минут в часе.
func spentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}
	speed := MeanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH, nil
}

// WalkingSpentCalories возвращает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := spentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories возвращает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	return spentCalories(steps, weight, height, duration)
}

// MeanSpeed возвращает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// Distance возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
