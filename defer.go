package main

import (
	"fmt"
)

func main() {
	a := 10
	defer llamada(a)

	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	a = 7
}

func llamada(i int) {
	fmt.Println(i)
}
