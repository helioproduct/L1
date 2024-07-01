package main

import (
	"fmt"
	"math"
)

// Point представляет точку в двумерном пространстве с инкапсулированными координатами x и y
type Point struct {
	x, y float64
}

// NewPoint является конструктором для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance рассчитывает расстояние между двумя точками
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(7, 1)

	distance := p1.Distance(p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
