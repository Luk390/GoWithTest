package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{height: 10.2, width: 2.0}
		want := 20.4 + 4.0
		checkPerimeter(t, rectangle, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{radius: 2.0}
		want := 4.0 * math.Pi
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v, got %g want %g", shape, got, want)
		}
	}
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{height: 10, width: 2,}, hasArea: 20,},
		{name: "Circle", shape: Circle{radius: 2,}, hasArea: 6.283185307179586,},
		{name: "Triangle", shape: Triangle{base: 3, height: 2}, hasArea: 3,},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.hasArea
			checkArea(t, tt.shape, want)
		})
	}


	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{10.2, 2.0}
	// 	want := 20.4
	// 	checkArea(t, rectangle, want)
	// })
	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{radius: 2.0}
	// 	want := 2.0 * math.Pi
	// 	checkArea(t, circle, want)
	// })

}
