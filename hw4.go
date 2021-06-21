package main

import (
	"fmt"
)

func sortSlice(arr []int) []int {
	for a := 1; a < len(arr); a++ {
		for b := a; b > 0 && arr[b-1] > arr[b]; b-- {
			arr[b-1], arr[b] = arr[b], arr[b-1]
		}
	}
	return arr
}
func main() {
	fmt.Println(sortSlice([]int{8, 10, 159, 1503, 3, 9}))
}
