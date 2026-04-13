package main

import "fmt"

func main() {
	minus(10, 7)
}
func minus(x, y int) {
	left := 0
	right := x

	for left < right {
		mid := (left + right) / 2 //бинарный поиск

		if mid+y < x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Println(left)
}
