package main

import "fmt"

func main() {
	one()
	two()
}

func one() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(xi)

	for a := range xi {
		fmt.Println("index:", a)
	}

	for a, b := range xi {
		fmt.Printf("xi[%v]: %v\n", a, b)
	}
}

func two() {
	m := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Six":   6,
		"Seven": 7,
	}

	for a := range m {
		fmt.Println(a)
	}

	for a, b := range m {
		fmt.Printf("%v = %v\n", a, b)
	}
}
