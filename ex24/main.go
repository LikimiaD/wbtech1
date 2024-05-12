package main

import (
	"fmt"
	"math"
)

type Point struct {
	//? Структура, с информацией о расположении X и Y
	X float64
	Y float64
}

// ? Функция создания структуры Point
func newPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

// ? Функция расчета расстояние от точки p(p1) до p2
func (p *Point) distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p.X, 2) + math.Pow(p2.Y-p.Y, 2))
}

func main() {
	x := newPoint(-2.3, 4)                                         //? Инициализация точки X
	y := newPoint(8.5, 0.7)                                        //? Инициализация точки Y
	fmt.Printf("Расстоение между т. X и Y -> %f\n", x.distance(y)) //? Вычисление расстояния
}
