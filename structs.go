/********************************************************************************/
/*                                                                              */
/*                              structs.go                                      */
/*                                                                              */
/* +++ Descripcion:                                                             */
/*                                                                              */
/*                                                                              */
/*  @ Autor  : Rodrigo G. Higuera M. <rodrigoghm@gmail.com>                     */
/*  @ Fecha  : 2020.07.25 13:37:17                                              */
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
	"time"
)

/*****************************************************/
/******** Definicion de Estructuras - Inicio *********/
/*****************************************************/
type Cuadrado struct {
	Lado float64
}

type Triangulo struct {
	Base, Altura float64
}

/*****************************************************/
/******** Definicion de Estructuras - Fin ************/
/*****************************************************/

/*****************************************************/
/***** Definicion de Constantes Globales - Inicio ****/
/*****************************************************/
const VERSION = "20200725.133717.0001"

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
* @date     2020.07.25 13:37:17
* @version  20200725.133717.0001
* @Company  © Development 2020
*
 */
func main() {
	rand.Seed(time.Now().UnixNano())
	min := 2.0
	max := 16.0
	b := rand.Float64()*(max-min) + min
	h := rand.Float64()*(max-min) + min

	C1 := Cuadrado{b}
	T1 := Triangulo{
		Base:   b,
		Altura: h,
	}

	fmt.Println("Calculo del area de un cuadrado:")
	fmt.Println("======================================")
	fmt.Println("Parametro Lado ===> ", b)
	fmt.Println("Resultado del area ===> ", getAreaCuadrado(C1.Lado))
	fmt.Println("Resultado del area [Con metodos] ===> ", C1.getAreaC())

	fmt.Println("")
	fmt.Println("Calculo del area de un triangulo:")
	fmt.Println("======================================")
	fmt.Println("Parametro Base ===> ", b)
	fmt.Println("Parametro Altura ===> ", h)
	fmt.Println("Resultado del area ===> ", getAreaTriangulo(T1.Base, T1.Altura))
	fmt.Println("Resultado del area [Con metodos] ===> ", T1.getAreaT())
	fmt.Println("")
}

func (c Cuadrado) getAreaC() float64 {
	return c.Lado * c.Lado
}

func (t Triangulo) getAreaT() float64 {
	return t.Base * t.Altura * 0.5
}

func getAreaCuadrado(b float64) float64 {
	return b * b
}

func getAreaTriangulo(b, h float64) float64 {
	return 0.5 * b * h
}
