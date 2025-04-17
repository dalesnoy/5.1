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

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	list := strings.Split(datastring, ",")
	if len(list) != 3 {
		return errors.New("incorrect number of arguments")
	}

	Steps, err := strconv.Atoi(list[0])
	if err != nil {
		return err
	}
	if Steps <= 0 {
		return errors.New("incorrect number of steps")
	}
	t.Steps = Steps

	TrainingType := strings.TrimSpace(list[1])
	t.TrainingType = TrainingType // Убрали проверку типа здесь

	duration, err := time.ParseDuration(strings.TrimSpace(list[2]))
	if err != nil {
		return fmt.Errorf("error while parsing duration")
	}
	if duration <= 0 {
		return errors.New("duration must be > 0")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var err error
	var calls float64
	var activity string
	dist := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		calls, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		activity = "Бег"
	case "Ходьба":
		calls, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		activity = "Ходьба"
	default:
		return "", fmt.Errorf("неизвестный тип тренировки %s", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("error counting callories")
	}

	info := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		activity,
		t.Duration.Hours(),
		dist,
		meanSpeed,
		calls,
	)

	return info, nil
}
