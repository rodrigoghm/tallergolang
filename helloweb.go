/********************************************************************************/
/*                                                                              */
/*                              helloweb.go                                     */
/*                                                                              */
/* +++ Descripcion:                                                             */
/*                                                                              */
/*                                                                              */
/*  @ Autor  : Rodrigo G. Higuera M. <rodrigoghm@gmail.com>                     */
/*  @ Fecha  : 2020.07.25 14:09:10                                              */
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
	"io"
	"net/http"
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
const VERSION = "20200725.14910.0001"

/*****************************************************/
/***** Definicion de Constantes Globales - Fin *******/
/*****************************************************/

/*****************************************************/
/****** Definicion de Variables Globales - Inicio ****/
/*****************************************************/

/*****************************************************/
/****** Definicion de Variables Globales - Fin ******/
/*****************************************************/

func callHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ejecutandose...")
	io.WriteString(w, "::: Hola mundo WEB con Golang ::: Taller Sodimac :::")
}

/*
*
* @author   Rodrigo G. Higuera M. <rodrigoghm@gmail.com>
* @date     2020.07.25 14:09:10
* @version  20200725.14910.0001
* @Company  © Development 2020
*
 */
func main() {

	http.HandleFunc("/", callHello)
	http.ListenAndServe(":8004", nil)
}
