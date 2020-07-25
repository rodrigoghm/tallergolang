package main

import (
	"fmt"
	_ "strconv"
	"tallergolang/util"
)

const VERSION string = "VALOR CONSTANTE"

func main() {
	a := 16
	b := 4
	resul, res := util.Division(a, b)
	fmt.Printf("Resultado ===> [%d] || Resto ===> [%d]\n", resul, res)
	resul, res = util.Divi(a, b)
	fmt.Printf("Resultado ===> [%d] || Resto ===> [%d]\n", resul, res)
	fmt.Println(VERSION)
}
