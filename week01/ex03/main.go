package main

import "fmt"

const (
	a = iota
	b
)

func main() {
	switch a := 1; {
	case a >= 0:
		fmt.Println("a i true")
		fallthrough
	case a > 200:
		fmt.Println("a is false")
	}

	stringSlice1 := []string{"1", "2", "3", "4"}
	fmt.Println(stringSlice1[:2])
}
