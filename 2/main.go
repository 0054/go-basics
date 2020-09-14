package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Radius = %d cm.\n", radius)
	fmt.Printf("circle Area = %f32 \n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("radius must not be negative")
	}
	return float32(radius) * float32(radius) * pi, nil
}

func main() {
	printCircleArea(2)
	x := 10
	p := &x
	// использую символ & перед именем переменной мы обращаемся к адресу памяти в котором хранится значение переменной
	// используя *p мы обращаемся к значению которое находится по адресу p
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(p)
	fmt.Println(*p)
}
