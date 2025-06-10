package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	// Проверим, что у нас три части
	if len(parts) != 3 {
		return 0,"", 0, errors.New("parseTraining: the length is less or more than three") 
	}

	// Преобразуем первый элемент в int (количество шагов)
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parseTraining: failed to convert steps: %w", err)
	}
	if steps <= 0 {
		err := errors.New("parseTraining: incorrect number of steps")
		log.Println(err) // Логируем ошибку
		return 0, "", 0, err
	}
	// Преобразуем третий элемент в time.Duration
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parseTraining: failed to parse duration: %w", err)
	}
	if duration.Seconds() < 0 || duration == 0{
		err := errors.New("parseTraining: negative durations are not allowed")
		log.Println(err) // Логируем ошибку
		return 0, "", 0, err
	}
	return steps, parts[1], duration, nil
}


func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// рассчитаем длину шага
	stepLength := ((height * stepLengthCoefficient) * float64(steps)) / mInKm
	return stepLength
}


func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// проверяем что продолжительность больше нуля
	if duration <= 0 { return 0 }

	// получаем часы из продолжительности
	hours := duration.Hours()

	// получаем дистанцию
	distance := distance(steps, height)

	// вычисляем среднюю скорость
	averageSpeed := distance / hours

	return averageSpeed
}


func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, actives, duration, _ := parseTraining(data)
	distance := distance(steps, height)
	meanSpeed := meanSpeed(steps, height, duration)
	switch actives {
	case "Бег":
		spentCalories, err := RunningSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf(
			"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
			actives, duration.Hours(), distance, meanSpeed, spentCalories,
			), err

	case "Ходьба":
		spentCalories, err := WalkingSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf(
			"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
			actives, duration.Hours(), distance, meanSpeed, spentCalories,
			), err

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}


func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверяем входные параметры на корректность
	if steps <= 0 { return 0, errors.New("RunningSpentCalories: the number of steps must be greater than 0") }
	if weight <= 0 { return 0, errors.New("RunningSpentCalories: the weight must be greater than 0") }
	if height <= 0 { return 0, errors.New("RunningSpentCalories: the height must be greater than 0") }
	if duration <= 0 { return 0, errors.New("RunningSpentCalories: the duration must be greater than 0") }

	// рассчитаем среднюю скорость
	meanSpeed := meanSpeed(steps, height, duration)

	// переводим продолжительность в минуты 
	durationInMinutes := duration.Minutes()

	// рассчитываем количество калорий
	amountСalories := (weight * meanSpeed * durationInMinutes) / minInH

	return amountСalories, nil
}


func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// проверяем входные параметры на корректность
	if steps <= 0 { return 0, errors.New("RunningSpentCalories: the number of steps must be greater than 0") }
	if weight <= 0 { return 0, errors.New("RunningSpentCalories: the weight must be greater than 0") }
	if height <= 0 { return 0, errors.New("RunningSpentCalories: the height must be greater than 0") }
	if duration <= 0 { return 0, errors.New("RunningSpentCalories: the duration must be greater than 0") }
	
	// рассчитываем среднюю скорость
	meanSpeed := meanSpeed(steps, height, duration)

	// переводим продолжительность в минуты 
	durationInMinutes := duration.Minutes()

	// рассчитываем количество калорий
	amountСalories := (weight * meanSpeed * durationInMinutes) / minInH

	spentCalories := amountСalories * walkingCaloriesCoefficient

	return spentCalories, nil
}
