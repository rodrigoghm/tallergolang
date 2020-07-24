package main

import "fmt"

func main() {
	XX()
	fmt.Println("After call to XX().")
}

func XX() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in XX() => ", r)
		}
	}()
	fmt.Println("Calling YY().")
	YY(0)
	fmt.Println("Returned normally from g.")
}

func YY(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in YY() => ", i)
	fmt.Println("Printing in YY()", i)
	YY(i + 1)
}
