package main

import (
	"fmt"
	"math"
)

// С маленькой буквы +- инкапсуляция

type Point struct {
	x float64 // private 
	y float64 // private
}

func newPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y} // Конструктор
}

func distance(p1, p2 *Point) float64 {
    dx := p1.x - p2.x
    dy := p1.y - p2.y
    return math.Sqrt(dx*dx + dy*dy) // Теорема Пифагора
}

func main() {
	A := newPoint(10, 6)
	B := newPoint(-1, 12)

	fmt.Println(distance(A, B))
}