package main

import (
	"fmt"
	"strings"
)

func rotarSecuencia(secuencia string, cantidadRotaciones int, decision int) {
	//Como la secuencia es string se tendra que realizar un split para asi trabajar con ella dentro
	//de la funcion
	splitSecuencia := strings.Split(secuencia, ",")

	fmt.Println("\nSecuencia original:", splitSecuencia)
	fmt.Println("Cantidad de rotaciones:", cantidadRotaciones)

	//Segun la funcion escogida por el usuario se hara la rotacion, 0 es para izquierda y 1 para derecha

	//En cada if habra un for que permitira llamar la funcion una vez, hasta que la secuencia se rote
	//segun la cantidad de rotaciones solicitadas por el usuario

	if decision == 0 {
		fmt.Println("Direccion:", decision, "= izquierda")
		for contador := 0; contador < cantidadRotaciones; contador++ {
			rotarIzquierda(splitSecuencia) // calling function to perform rotation ont time
		}

	}

	if decision == 1 {
		fmt.Println("Direccion:", decision, "= derecha")
		for contador := 0; contador < cantidadRotaciones; contador++ {
			rotarDerecha(splitSecuencia) // calling function to perform rotation ont time
		}
	}
	fmt.Println("Secuencia final rotada:", splitSecuencia)
}

// Esta funcion permitira rotar una secuencia una vez a la izquierda
func rotarIzquierda(secuencia []string) {
	//Para realizar la rotacion existira un contador que determinara cuantas veces debera rotar
	//uno de los elementos segun la cantidad de rotaciones escogidas por el usuario
	var contador int = 0

	//Temporal permitira gracias al for enviar el resultado de cada una de las rotaciones realizadas
	//a la secuencia
	var temporal string = secuencia[0]

	//En el for se toma en cuenta el largo de la secuencia menos 1, finalizara una vez
	//contador sea menor que el largo de la secuencia menos 1 y se incrementara contador
	for ; contador < len(secuencia)-1; contador++ {
		//La secuencia se modificara en la posicion de contador
		//se asignara el contador mas 1
		secuencia[contador] = secuencia[contador+1]
	}
	//La nueva secuencia de temporal se asignara segun el contador
	secuencia[contador] = temporal
}

// Esta funcion permitira rotar una secuencia una vez a la derecha
func rotarDerecha(secuencia []string) {
	//Para realizar la rotacion existira un contador que determinara cuantas veces debera rotar
	//uno de los elementos segun la cantidad de rotaciones escogidas por el usuario, en este caso
	//para la derecha se debera obtener el largo de la secuencia menos 2
	var contador int = len(secuencia) - 2

	//Temporal permitira gracias al for enviar el resultado de cada una de las rotaciones realizadas
	//a la secuencia, en este caso temporal iniciara en el largo de la secuencia menos 1
	var temporal string = secuencia[len(secuencia)-1]

	//En el for se finalizara cuando contador sea mayor o igual que cero y se disminuira contador
	for ; contador >= 0; contador-- {
		//La secuencia se modificara en la posicion de contador mas 1
		//se asignara el contador
		secuencia[contador+1] = secuencia[contador]
	}
	//La nueva secuencia de temporal se asignara en cero
	secuencia[0] = temporal
}

func main() {

	//Utilizaremos una secuencia de elementos tipo string que sera inmutable en esta seccion
	secuenciaElementos := "a, b, c, d, e, f, g, h"

	//Llamaremos a la funcion del programa
	rotarSecuencia(secuenciaElementos, 3, 0)
}
