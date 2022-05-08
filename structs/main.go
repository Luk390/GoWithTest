package main

import "math"

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base float64
	height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	perimeter := (r.width + r.height) * 2
	return perimeter
}

func (c Circle) Perimeter() float64 {
	perimiter := math.Pi * (c.radius * 2)
	return perimiter
}

func(t Triangle) Perimeter() float64 {
	return 0
}

func (r Rectangle) Area() float64 {
	area := r.height * r.width
	return area
}

func (c Circle) Area() float64 {
	area := c.radius * math.Pi
	return area
}

func (t Triangle) Area() float64 {
	area := (t.base * t.height) / 2
	return area+1
}

