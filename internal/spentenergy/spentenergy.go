package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	speed     = 1.39  // средняя скорость в м/с
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:

// steps int - количество шагов.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	if weight <= 0 || height <= 0 {
		fmt.Println("error: weight or height less or equal 0")
		return 0
	}
	if duration <= 0 {
		fmt.Println("error: duration less or equal 0")
		return 0

	}
	mean := МeanSpeed(steps, duration)
	calories := ((walkingCaloriesWeightMultiplier * weight) + (mean*mean/height)*walkingSpeedHeightMultiplier) * float64(duration) * minInH
	return calories
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	if weight <= 0 {
		fmt.Println("error: weight less or equal 0")
		return 0
	}
	if duration <= 0 {
		fmt.Println("error: duration less or equal 0")
		return 0
	}
	mean := МeanSpeed(steps, duration)
	calories := (runningCaloriesMeanSpeedMultiplier * mean) - runningCaloriesMeanSpeedShift*weight
	return calories

}

// МeanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func МeanSpeed(steps int, duration time.Duration) float64 {
	if duration < 0 {
		return 0
	}
	dist := Distance(steps)
	mean := dist / duration.Hours()
	return mean
}

// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Для расчета дистанции нужно шаги умножить на длину шага lenStep и разделить на mInKm
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
//
// Создайте функцию ниже
func Distance(steps int) float64 {
	dist := (float64(steps) * lenStep) / mInKm
	return dist
}
