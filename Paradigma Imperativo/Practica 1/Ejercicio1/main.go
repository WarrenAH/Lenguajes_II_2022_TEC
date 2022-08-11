package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func escanerTexto(textoEscrito string) int {

	//Para realizar el escaner de texto se tomo de ejemplo como lo hace Word, este no nos detecta
	//los saltos de linea o las tildes como caracteres, cada palabra se representa con un espacio, etc

	print("\nSu texto ingresado: \n\n")

	print(textoEscrito)
	//Encontrar los caracteres:
	//Se detectan los saltos de linea y las tildes ya que no son considerados caracteres para luego restarlos
	//entre la cantidad de caracteres totales del texto
	caracteresFinales := len(textoEscrito) - strings.Count(textoEscrito, "á") - strings.Count(textoEscrito, "é") - strings.Count(textoEscrito, "í") -
		strings.Count(textoEscrito, "ó") - strings.Count(textoEscrito, "ú") - strings.Count(textoEscrito, "Á") - strings.Count(textoEscrito, "É") -
		strings.Count(textoEscrito, "Í") - strings.Count(textoEscrito, "Ó") - strings.Count(textoEscrito, "Ú") - strings.Count(textoEscrito, "\n")

	print("\n")
	fmt.Printf("\nCantidad de caracteres (incluyendo espacios): %d\n", (caracteresFinales))

	//Encontrar las palabras:
	//Ignoramos los los caracteres sin espacio para contar las palabras.
	ignorarEnTexto := regexp.MustCompile(`[\S]+`)

	//Se encuentran las palabras.
	resultadoFinal := ignorarEnTexto.FindAllString(textoEscrito, -1)

	fmt.Printf("Cantidad de palabras: %d\n", len(resultadoFinal))

	//Encontrar las lineas
	//se escanea el texto
	escanear := bufio.NewScanner(strings.NewReader(textoEscrito))

	//Se utiliza la funcion split
	escanear.Split(bufio.ScanLines)

	//Se lleva un contador de las lineas.
	contadorLineas := 0

	//Se utiliza un for para escanear el texto
	for escanear.Scan() {
		contadorLineas++
	}

	fmt.Printf("Cantidad de lineas: %d\n", (contadorLineas))

	return 0
}

func main() {
	//agregaTexto("To be or not to be, that is the question. \n Vivimos en una sociedad.")
	//print("To be or not to be, that is the question.\nVivimos en una sociedad.")
	escanerTexto("La programación es el proceso de crear un conjunto de instrucciones que le dicen a una " +
		"\ncomputadora como realizar algún tipo de tarea. Pero no solo la acción de escribir \nun código para que la " +
		"computadora o el software lo ejecute. Incluye, además, todas las tareas \nnecesarias para que el código " +
		"funcione correctamente y cumpla el objetivo para el cual se \nescribió. \nEn la actualidad, la noción de " +
		"programación se encuentra muy asociada a la creación de \naplicaciones de informática y videojuegos. En este " +
		"sentido, es el proceso por el cual una \npersona desarrolla un programa, valiéndose de una herramienta que le " +
		"permita escribir el \ncódigo (el cual puede estar en uno o varios lenguajes, como C++, Java y Python, entre " +
		"\nmuchos otros) y de otra que sea capaz de traducirlo a lo que se conoce como lenguaje de \nmáquina, que " +
		"puede comprender el microprocesador.")
}

/*
RESULTADO:
Su texto ingresado:

La programación es el proceso de crear un conjunto de instrucciones que le dicen a una
computadora como realizar algún tipo de tarea. Pero no solo la acción de escribir
un código para que la computadora o el software lo ejecute. Incluye, además, todas las tareas
necesarias para que el código funcione correctamente y cumpla el objetivo para el cual se
escribió.
En la actualidad, la noción de programación se encuentra muy asociada a la creación de
aplicaciones de informática y videojuegos. En este sentido, es el proceso por el cual una
persona desarrolla un programa, valiéndose de una herramienta que le permita escribir el
código (el cual puede estar en uno o varios lenguajes, como C++, Java y Python, entre
muchos otros) y de otra que sea capaz de traducirlo a lo que se conoce como lenguaje de
máquina, que puede comprender el microprocesador.

Cantidad de caracteres (incluyendo espacios): 852
Cantidad de palabras: 145
Cantidad de lineas: 11
*/
