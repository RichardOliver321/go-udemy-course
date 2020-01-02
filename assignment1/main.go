package main

import "fmt"

func main() {
	s := []int{}

	for i := range [10]int{} {
		s = append(s, i)
	}

	for i := range s {
		fmt.Println(i, isEvenOrOdd(s[i]))
	}
}

func isEvenOrOdd(value int) string {
	if value%2 == 0 {
		return " is even"
	}
	return " is odd"
}
