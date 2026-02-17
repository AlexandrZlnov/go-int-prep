// Пример реализации паттерна - Bridge
// Bridge (Мост) — это структурный паттерн, который разделяет
// абстракцию и её реализацию так, чтобы их можно было изменять независимо друг от друга.
// Вместо жёсткого наследования создаётся связь через интерфейс: абстракция содержит
// ссылку на объект реализации и делегирует ему часть работы. Это позволяет
// комбинировать разные типы абстракций с разными реализациями без взрыва количества классов.
// Обычно реализуется через интерфейс реализации и структуру абстракции, которая хранит
// ссылку на этот интерфейс.

// Абстракция (Circle) содержит ссылку на реализацию (Renderer).

package main

import (
	"fmt"
)

// Интерфейс реализации
type Renderer interface {
	RenderCircle(float64)
	RenderSquare(int64)
}

// Конкретные реализации
type VectorRenderer struct {
}

func (vr *VectorRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing Circle as Vector with Radius = %.2f\n", radius)
}

func (vr *VectorRenderer) RenderSquare(side int64) {
	fmt.Printf("Drawing Square as Vector with Side = %d\n", side)
}

type RasterRenderer struct {
}

func (rr *RasterRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing Circle as Pixels with Radius = %.2f\n", radius)
}

func (rr *RasterRenderer) RenderSquare(side int64) {
	fmt.Printf("Drawing Square as Pixels with Side = %d\n", side)
}

// Абстракция - верхнеуровневая
type Shape interface {
	Draw()
}

// Уточненные абстракции
type Circle struct {
	renderer Renderer
	radius   float64
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

type Square struct {
	renderer Renderer
	side     int64
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.side)
}

// Клиент
func main() {
	vector := &VectorRenderer{}
	raster := &RasterRenderer{}

	// Круг, нарисованный растром
	circle1 := &Circle{
		renderer: vector,
		radius:   10,
	}
	circle1.Draw()

	// Круг, нарисованный векторно
	circle2 := &Circle{
		renderer: raster,
		radius:   15,
	}
	circle2.Draw()

	// Квадрат, нарисованный векторно
	square1 := &Square{
		renderer: vector,
		side:     20,
	}
	square1.Draw()

	// Квадрат, нарисованный растром
	square2 := &Square{
		renderer: raster,
		side:     22,
	}
	square2.Draw()
}
