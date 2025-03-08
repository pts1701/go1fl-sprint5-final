package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")
	if len(s) != 3 {
		return errors.New("incorrect input data")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return fmt.Errorf("incorrect steps input: %v", err)
	}
	t.Steps = steps

	activity := s[1]
	if activity != "Бег" && activity != "Ходьба" {
		return errors.New("unknown training type")
	}
	t.TrainingType = activity

	duration, err := time.ParseDuration(s[2])
	if err != nil {
		return fmt.Errorf("incorrect duration: %v", err)
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {

	if t.Duration.Hours() < 0 {
		return fmt.Sprintf(""), errors.New("negative duration")
	}

	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return fmt.Sprintf("Неизвестный вид тренировки"), errors.New("unknown training type")
	}

	dist := spentenergy.Distance(t.Steps)

	speed := spentenergy.МeanSpeed(t.Steps, t.Duration)

	var calories float64

	switch t.TrainingType {

	case "Бег":
		calories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)

	case "Ходьба":
		calories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f. ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), dist, speed, calories), nil

}
