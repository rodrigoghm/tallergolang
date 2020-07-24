package main

import (
	"fmt"
	_ "strconv"
)

const VERSION string = "VALOR CONSTANTE"

func main() {
	a := 16
	b := 4
	resul, res := division(a, b)
	fmt.Printf("Resultado ===> [%d] || Resto ===> [%d]\n", resul, res)
	resul, res = divi(a, b)
	fmt.Printf("Resultado ===> [%d] || Resto ===> [%d]\n", resul, res)
	fmt.Println(VERSION)
}

func division(dividendo, divisor int) (int, int) {
	var resto int
	r := dividendo / divisor
	resto = dividendo % divisor
	return r, resto
}

func divi(dividendo, divisor int) (r int, resto int) {
	r = dividendo / divisor
	resto = dividendo % divisor
	return
}
