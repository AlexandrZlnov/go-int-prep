package main

import (
	"fmt"
)

func createTempAdjuster() (func(val float64) float64, float64) {
	baseTemp := 90.0

	adjusterTemp := func(val float64) float64 {
		baseTemp += val

		return baseTemp
	}

	return adjusterTemp, baseTemp
}

func main() {
	tAdjustment, bTemp := createTempAdjuster()

	fmt.Println("Базовая температура = ", bTemp)

	fmt.Println("Отрегулируем температуру на 3 градуса. Темп = ", tAdjustment(3))
	fmt.Println("Отрегулируем температуру на 6 градуса. Темп = ", tAdjustment(6))
	fmt.Println("Отрегулируем температуру на -4 градуса. Темп = ", tAdjustment(-4))

}

// Вывод:
// Базовая температура =  90
// Отрегулируем температуру на 3 градуса. Темп =  93
// Отрегулируем температуру на 6 градуса. Темп =  99
// Отрегулируем температуру на -4 градуса. Темп =  95
