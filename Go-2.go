package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&a)

	fmt.Println("диаметр:", 2*math.Sqrt(a/math.Pi))
	fmt.Println("Длина окружности", (2*math.Sqrt(a/math.Pi))*math.Pi)

}
