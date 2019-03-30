package main

import (
	"testing"
)

func checkError(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	checkError(t, got, want)
}

func TestArea(t *testing.T) {

	/*
	匿名结构体。类似：
			type Foo struct {
			shape Shape
			want float64
			}
			areaTests:=[]Foo {
			{Rectangle{12,6}, 72.0},
			{Circle{10}, 314.1592653589793}
			}
	*/
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
