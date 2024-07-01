package main

import (
	"fmt"
)

// groupTemperatures группирует температуры с шагом в 10 градусов и возвращает мапу групп
func groupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temps {
		group := int(temp) / 10 * 10
		// Добавляем температуру в соответствующую группу
		groups[group] = append(groups[group], temp)
	}
	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroups := groupTemperatures(temperatures)
	for key, values := range tempGroups {
		fmt.Printf("%d: %v\n", key, values)
	}
}
