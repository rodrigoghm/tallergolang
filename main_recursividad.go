package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"tallergolang/util"
)

func main() {
	cp := len(os.Args)

	if cp < 3 {
		log.Fatal("Error faltan parametros")
	} else {
		b, _ := strconv.Atoi(os.Args[1])
		e, _ := strconv.Atoi(os.Args[2])

		if e < 1 {
			panic("Error exponente indebido...")
		} else {
			r := util.Potencia(uint64(b), uint64(e))
			fmt.Println("Resultado ===> ", r)
		}
	}

	f := float64(util.Potencia(2, 5))
	fmt.Printf("Resultado de f ===> %f\n\n", f)
}
