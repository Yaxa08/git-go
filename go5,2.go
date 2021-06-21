package main

import "fmt"

func fibbonachi(n uint) uint {

	fn := make(map[uint]uint)
	var i uint
	for i = 0; i <= n; i++ {
		var f uint
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	fmt.Print("ввидите число:")
	var a uint
	fmt.Scan(&a)
	fmt.Println(fibbonachi(a))
}
