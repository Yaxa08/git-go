package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	var e int = a % 100 % 10
	var d int = (a%100 - e) / 10
	var c int = (a - d - e) / 100

	fmt.Println("сотни: ", c)
	fmt.Println("Десятки: ", d)
	fmt.Println("Еденицы: ", e)

}
