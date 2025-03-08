package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength     = 0.65
	KiloCorrection = 1000
)

// создайте структуру DaySteps
type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {

	s := strings.Split(datastring, ",")
	if len(s) != 2 {
		return errors.New("incorrect input data")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return fmt.Errorf("incorrect steps input: %v", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return fmt.Errorf("incorrect duration: %v", err)
	}

	ds.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration.Hours() <= 0 {
		return fmt.Sprintf(""), errors.New("negative duration")
	}
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	distance := (float64(ds.Steps) * StepLength) / KiloCorrection
	return fmt.Sprintf("Количество шагов: %d\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
