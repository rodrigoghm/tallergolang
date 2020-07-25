/********************************************************************************/
/*                                                                              */
/*                              sorteo.go                                       */
/*                                                                              */
/* +++ Descripcion:                                                             */
/*                                                                              */
/*                                                                              */
/*  @ Autor  : Rodrigo G. Higuera M. <rodrigoghm@gmail.com>                     */
/*  @ Fecha  : 2020.07.23 13:38:57                                              */
/*                                                                              */
/* C0pyl3ft - 2020 | Open Source License                                        */
/********************************************************************************/
/********************************************************************************/
/* Ejecucion :                                                                  */
/*                                                                              */
/*                                                                              */
/********************************************************************************/

package main

/**********************************/
/******** Import Packages *********/
/**********************************/
import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*****************************************************/
/******** Definicion de Estructuras - Inicio *********/
/*****************************************************/

/*****************************************************/
/******** Definicion de Estructuras - Fin ************/
/*****************************************************/

/*****************************************************/
/***** Definicion de Constantes Globales - Inicio ****/
/*****************************************************/
const VERSION = "20200723.133857.0001"

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
* @date     2020.07.23 13:38:57
* @version  20200723.133857.0001
* @Company  © Development 2020
*
 */
func main() {
	rand.Seed(time.Now().UnixNano())
	p := len(os.Args)
	fmt.Println(p)
	premio := make([]int, 0)
	index := 0

	if p == 3 {
		min, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}

		fmt.Println("<Min> ==> ", min)
		fmt.Println("<Max> ==> ", max)

		var n int
		for {
			if index == 3 {
				break
			}
			n = rand.Intn(max-min+1) + min
			f := buscar(premio, n)
			if f == false {
				premio = append(premio, n)
				index++
			}
		}

		fmt.Println(premio)
		ShowSlice(premio)

	} else {
		fmt.Println("Error: Debe i")
		panic("Error debe ingresar los parametros min y max <numericos>")
	}
}

func buscar(s []int, v int) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}

	return false
}

func ShowSlice(s []int) {
	for i, item := range s {
		if i == 2 {
			time.Sleep(3 * time.Second)
			fmt.Println(">>> GANADOR(A) <<< ===> ", item)
		} else {
			time.Sleep(2 * time.Second)
			fmt.Println(">>> AL AGUA <<< ===> ", item)
		}
	}
}
