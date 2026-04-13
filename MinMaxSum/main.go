package main

import (
	"fmt"
)

// -------- MIN (без if) --------
//сравниваем первый элемент со всеми, если есть меньше берем его 
func myMin(arr []int) int {
	m := arr[0]
	for _, x := range arr[1:] { 
		m = (m + x - abs(m-x)) / 2
	}
	return m
}

// -------- MAX (без if) --------
//аналагично минимуму
func myMax(arr []int) int {
	m := arr[0]
	for _, x := range arr[1:] {
		m = (m + x + abs(m-x)) / 2
	}
	return m
}

// -------- SUM --------
func mySum(arr []int) int {
	s := 0
	for _, x := range arr {
		s += x
	}
	return s
}

// -------- PROD --------
func myProd(arr []int) int {
	p := 1
	for _, x := range arr {
		p *= x
	}
	return p
}

// -------- FILTER (оставим чётные) --------
func myFilter(arr []int) []int {
	var res []int
	for _, x := range arr {
		if x%2 == 0 {
			res = append(res, x)
		}
	}
	return res
}

// -------- ABS (своя, без math) --------
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// -------- MAIN --------
func main() {
	arr := []int{5, -2, 10, 3, 8}

	fmt.Println("Array:", arr)
	fmt.Println("Min:", myMin(arr))
	fmt.Println("Max:", myMax(arr))
	fmt.Println("Sum:", mySum(arr))
	fmt.Println("Prod:", myProd(arr))
	fmt.Println("Filter (even):", myFilter(arr))
}
