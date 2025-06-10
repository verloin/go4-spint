package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	// Проверим, что у нас две части
	if len(slc) != 2 {
		err := errors.New("parsePackage: incorrect length")
		log.Println(err) // Логируем ошибку
		return 0, 0, err // Возвращаем ошибку
	}
	// Преобразуем первый элемент в int (количество шагов)
	steps, err := strconv.Atoi(slc[0])
	if err != nil {
		return 0, 0, fmt.Errorf("parsePackage: failed to convert steps: %w", err)
	}
	// Проверяем, что количество шагов больше 0
	if steps <= 0 {
		err := errors.New("parsePackage: incorrect number of steps")
		log.Println(err) // Логируем ошибку
		return 0, 0, err
	}
	// Преобразуем второй элемент в time.Duration
	duration, err := time.ParseDuration(slc[1])
	if err != nil {
		return 0, 0, fmt.Errorf("parsePackage: failed to parse duration: %w", err)
	}
	// Проверка на отрицательную длительность времени
	if duration.Seconds() < 0 || duration == 0{
		err := errors.New("parsePackage: negative durations are not allowed")
		log.Println(err) // Логируем ошибку
		return 0, 0, err
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {	return "" }
	distance := (float64(steps) * stepLength) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}