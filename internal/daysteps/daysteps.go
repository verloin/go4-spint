package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slc := strings.Split(data, ",")

	if len(slc) != 2 {
		return 0, 0, errors.New("parseDuration: the length is less than or more than two")
	}
	steps, err := strconv.Atoi(slc[0])
	if err != nil {
		return 0, 0, fmt.Errorf("parseDuration: failed to convert steps: %w", err)
	}
	if steps == 0 {
		return 0, 0, errors.New("parseDuration: the number of steps is zero")
	}

	s := strings.ReplaceAll(slc[1], "min", "m")
	duration, err := time.ParseDuration(s)
	if err != nil {
		return 0, 0, fmt.Errorf("parseDuration: failed to parse duration: %w", err)
	}
	return 0, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Errorf("failed to convert steps: %w", err)
		return ""
	}
	if steps == 0 {	return "" }
	distance := (float64(steps) * stepLength) / mInKm
	calories := WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.1f км.\nВы сожгли %.1f ккал.", steps, distance, calories)

}
