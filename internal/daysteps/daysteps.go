package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"github.com/Yandex-Practicum/tracker/internal/trainings"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
	trainings.Training
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат строки: %s", datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка при парсинге количества шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("count of steps must be > 0")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка при парсинге длительности прогулки: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("count of steps must be > 0")
	}
	if ds.Weight <= 0 {
		return "", errors.New("weight must be > 0")
	}
	if ds.Height <= 0 {
		return "", errors.New("height must be > 0")
	}
	if ds.Duration <= 0 {
		return "", errors.New("продолжительность должна быть положительной")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	var calls float64
	var err error

	// Если тип тренировки не указан, считаем, что это ходьба (как в тестах)
	trainingType := ds.Training.TrainingType
	if trainingType == "" {
		trainingType = "Ходьба"
	}

	switch trainingType {
	case "Бег":
		calls, err = spentenergy.RunningSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка при расчёте калорий для бега: %v", err)
		}
	case "Ходьба":
		calls, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка при расчёте калорий для ходьбы: %v", err)
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
	}

	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calls,
	)

	return info, nil
}
