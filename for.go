/********************************************************************************/
/*                                                                              */
/*                              for.go                                          */
/*                                                                              */
/* +++ Descripcion:                                                             */
/*                                                                              */
/*                                                                              */
/*  @ Autor  : Rodrigo G. Higuera M. <rodrigoghm@gmail.com>                     */
/*  @ Fecha  : 2020.07.23 23:21:31                                              */
/*                                                                              */
/* C0pyl3ft - 2020 | Open Source License                                        */
/********************************************************************************/
/********************************************************************************/
/* Ejecucion :                                                                  */
/*                                                                              */
/*                                                                              */
/********************************************************************************/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

/**********************************/
/******** Import Packages *********/
/**********************************/

/*****************************************************/
/******** Definicion de Estructuras - Inicio *********/
/*****************************************************/

/*****************************************************/
/******** Definicion de Estructuras - Fin ************/
/*****************************************************/

/*****************************************************/
/***** Definicion de Constantes Globales - Inicio ****/
/*****************************************************/
const VERSION = "20200723.232131.0001"

/*****************************************************/
/***** Definicion de Constantes Globales - Fin *******/
/*****************************************************/

/*****************************************************/
/****** Definicion de Variables Globales - Inicio ****/
/*****************************************************/

/*****************************************************/
/****** Definicion de Variables Globales - Fin ******/
/*****************************************************/

/*
*
* @author   Rodrigo G. Higuera M. <rodrigoghm@gmail.com>
* @date     2020.07.23 23:21:31
* @version  20200723.232131.0001
* @Company  © Development 2020
*
 */
func main() {
	cp := len(os.Args)
	if cp < 2 {
		log.Fatal("Cantidad de parametros incorrectos")
		panic("Entre en panico!")
	}

	// FOR three component
	fmt.Println("for three component")
	fmt.Println("====================")
	n, _ := strconv.Atoi(os.Args[1])
	for i := n; i >= 1; i-- {
		fmt.Println("Valor de i ===> ", i)
	}

	// FOR parametros opcionales
	fmt.Println("for params optional")
	fmt.Println("====================")
	j, _ := strconv.Atoi(os.Args[1])
	for j >= 1 {
		fmt.Println("Valor de <j> ===> ", j)
		j--
	}

	//For = while
	fmt.Println("for como while")
	fmt.Println("====================")
	for n >= 1 {
		fmt.Println("Valor de <n> ===> ", n)
		n--
	}

	//For = while
	fmt.Println("for infinito")
	fmt.Println("====================")
	x := 500
	for {
		fmt.Println("::: ciclo infinito :::")
		if x == 0 {
			break
		}
		x--
	}

	// exit a loop
	fmt.Println("exit a loop")
	fmt.Println("====================")
	sum := 0
	for y := 1; y <= 5; y++ {
		if y%2 != 0 {
			continue
		}
		sum += y
	}

	log.Println("Valor de la suma ===> ", sum)

	//for con range.
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, val := range v {
		go fmt.Printf("Valor [%d]\n", val)
	}

	time.Sleep(5 * time.Second)

}
