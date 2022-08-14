package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var resultadoNombre = l.buscarProducto(nombre)

	reader := bufio.NewReader(os.Stdin)

	if resultadoNombre != -1 {
		fmt.Printf("\nEn el sistema ya esta el producto %s, Desea agregar existencias? Introduzca un 1 para si o un 0 para no.", nombre)
		entrada1, _ := reader.ReadString('\n')           // Leer hasta el separador de salto de línea
		eleccion1 := strings.TrimRight(entrada1, "\r\n") // Remover el salto de línea de la entrada del usuario

		decisionCantidad, err := strconv.Atoi(eleccion1)

		if decisionCantidad == 1 {
			fmt.Printf("\nIntroduzca las nuevas existencias que desea agregar para %s.", nombre)
			entrada2, _ := reader.ReadString('\n')           // Leer hasta el separador de salto de línea
			eleccion2 := strings.TrimRight(entrada2, "\r\n") // Remover el salto de línea de la entrada del usuario

			castingEleccion, err := strconv.Atoi(eleccion2)

			if err != nil {
				// ... handle error
				panic(err)
			}

			if castingEleccion < 1 {
				fmt.Printf("\nSolo se permite agregar almenos una unidad, vuelva a intentarlo")
				return
			}

			(*l)[resultadoNombre].cantidad = (*l)[resultadoNombre].cantidad + castingEleccion
		}

		if decisionCantidad == 0 {
			fmt.Printf("\nNo se agregaron nuevas existencias para %s", nombre)
		}

		fmt.Printf("\nDesea cambiarle el precio al producto %s? Introduzca un 1 para si o un 0 para no.", nombre)
		entrada3, _ := reader.ReadString('\n')           // Leer hasta el separador de salto de línea
		eleccion3 := strings.TrimRight(entrada3, "\r\n") // Remover el salto de línea de la entrada del usuario

		decisionPrecio, err := strconv.Atoi(eleccion3)

		if err != nil {
			// ... handle error
			panic(err)
		}

		if decisionPrecio == 1 {
			fmt.Printf("\nIntroduzca el nuevo precio para %s.", nombre)
			entrada4, _ := reader.ReadString('\n')           // Leer hasta el separador de salto de línea
			eleccion4 := strings.TrimRight(entrada4, "\r\n") // Remover el salto de línea de la entrada del usuario

			castingPrecio, err := strconv.Atoi(eleccion4)

			if err != nil {
				// ... handle error
				panic(err)
			}

			if castingPrecio < 1 {
				fmt.Printf("\nSolo se permite agregar precios positivos, vuelva a intentarlo")
				return
			}

			(*l)[resultadoNombre].precio = castingPrecio

		} else {
			fmt.Printf("\nNo se cambio el nombre para %s", nombre)
		}
		return
	}

	if resultadoNombre == -1 {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})

	}

	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente

}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant

			if (*l)[prod].cantidad == 0 {
				(*l) = append((*l)[:prod], (*l)[prod+1:]...)
			}
		} else {
			fmt.Println("No se puede vender mayor cantidad de productos que los que hay en existencia")
		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	var listaMinima listaProductos
	var contador int = 0
	for {
		if contador == len((*l)) {
			break
		}
		if (*l)[contador].cantidad <= existenciaMinima {

			listaMinima = append(listaMinima, (*l)[contador])
		}
		contador++
	}
	// debe retornar una nueva lista con productos con existencia mínima
	return listaMinima
}

// Esta es nuestra funcion para aumentar el inventario de minimos
func aumentarInventarioDeMinimos(nombre string, listaMinimos listaProductos, lProductos listaProductos) {
	//Buscamos el nombre del producto para luego ser modificado su cantidad tanto en la lista de
	//minimos como en la de productos
	var prodLProductos = lProductos.buscarProducto(nombre)

	//Buscamos el nombre del producto en la lista de minimos
	var prodListaMinimos = -1
	var i int
	for i = 0; i < len(listaMinimos); i++ {
		if (listaMinimos)[i].nombre == nombre {
			prodListaMinimos = i
		}
	}
	//Buscamos el producto en la lista de minimos para ver si se encuentra
	if prodListaMinimos != -1 {
		reader := bufio.NewReader(os.Stdin)

		//Se crea un for para que se vaya agregando la cantidad de unidades necesarias de un producto
		//Hasta que se cumpla la existencia minima
		for {

			//Si la cantidad es igual o mayor a la existencia minima se finaliza
			if (listaMinimos)[prodListaMinimos].cantidad >= existenciaMinima {
				fmt.Printf("\nEl producto %s tiene %d unidades, se cumplio con el minimo establecido de %d unidades.", nombre, (listaMinimos)[prodListaMinimos].cantidad, existenciaMinima)
				break
			}

			//Se le pregunta al usuario si desea agregar unidades o directamente salir, si se da que si
			//se ciclara hasta que se cumple el if anterior.

			fmt.Printf("\nEn el sistema el producto %s tiene %d unidades, incumpliendo con el minimo establecido de %d unidades, Desea agregar existencias? Introduzca un 1 para si o un 0 para no.", nombre, (listaMinimos)[prodListaMinimos].cantidad, existenciaMinima)
			entrada1, _ := reader.ReadString('\n')
			eleccion1 := strings.TrimRight(entrada1, "\r\n")

			decisionCantidad, err := strconv.Atoi(eleccion1)

			if err != nil {
				// ... handle error
				panic(err)
			}

			//Si la decision es 1 se procedera a agregar las unidades tanto en la lista de minimos
			//como en la de productos
			if decisionCantidad == 1 {
				fmt.Printf("\nIntroduzca las nuevas existencias que desea agrega para %s.", nombre)
				entrada2, _ := reader.ReadString('\n')
				eleccion2 := strings.TrimRight(entrada2, "\r\n")

				castingEleccion, err := strconv.Atoi(eleccion2)

				if err != nil {
					// ... handle error
					panic(err)
				}

				if castingEleccion < 1 {
					fmt.Printf("\nSolo se permite agregar almenos una unidad, vuelva a intentarlo")
					return
				}

				(lProductos)[prodLProductos].cantidad = (lProductos)[prodLProductos].cantidad + castingEleccion
				(listaMinimos)[prodListaMinimos].cantidad = (listaMinimos)[prodListaMinimos].cantidad + castingEleccion
			}

			//Si la decision es un 0 se saldra.
			if decisionCantidad == 0 {
				fmt.Printf("\nNo se agregaron nuevas existencias para %s", nombre)
				break
			}
		}
		return
	}

	//Si el producto no se encuentra en la lista de minimos no se continuara.
	if prodListaMinimos == -1 {
		fmt.Printf("\nEl producto %s no esta en la lista de minimos.", nombre)
		return
	}
}

// Para ordenar las listas se usa la libreria "sort"
func (l *listaProductos) ordenarLista(metodoOrdenamiento string) listaProductos {
	//Los productos se ordenaran segun su nombre
	if metodoOrdenamiento == "nombre" {
		sort.Slice((*l), func(i, j int) bool { return (*l)[i].nombre < (*l)[j].nombre })
	}

	//Los productos se ordenaran segun su cantidad de menor a mayor
	if metodoOrdenamiento == "cantidad" {
		sort.Slice((*l), func(i, j int) bool { return (*l)[i].cantidad < (*l)[j].cantidad })
	}

	//Los productos se ordenaran segun su precio de menor a mayor
	if metodoOrdenamiento == "precio" {
		sort.Slice((*l), func(i, j int) bool { return (*l)[i].precio < (*l)[j].precio })
	}

	return (*l)
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	lProductos.venderProducto("arroz", 4)
	fmt.Println(lProductos)
	lProductos.agregarProducto("azucar", 20, 1500)
	fmt.Println(lProductos)
	lProductos.venderProducto("frijoles", 4)
	fmt.Println(lProductos)
	lProductos.venderProducto("leche", 10)

	//Para efectos del ejercicio, solo se imprimiran los datos importantes.

	//Se imprimime la lista de productos
	fmt.Println("\nNuestra lista de productos")
	fmt.Println(lProductos)

	//Producto agregado anteriormente.
	///lProductos.agregarProducto("leche", 20, 1500)

	//Vendiendo un articulo hasta que desaparezca de la lista
	///lProductos.venderProducto("leche", 8)

	//Nuestra lista con productos que no cumplen con la existencia minima.
	var listaMinima listaProductos

	//Llenando la lista con productos que no cumplen con la existencia minima.
	listaMinima = lProductos.listarProductosMínimos()

	//Se imprime la lista.
	fmt.Println("\nNuestra lista de minimos")
	fmt.Println(listaMinima)

	//Se busca un producto que no existe en la lista con productos que no cumplen con la existencia minima.
	aumentarInventarioDeMinimos("lech", listaMinima, lProductos)

	//Se busca un producto que si existe en la lista con productos que no cumplen con la existencia minima.
	aumentarInventarioDeMinimos("leche", listaMinima, lProductos)

	//Volvemos a revisar nuestra lista minima
	listaMinima = lProductos.listarProductosMínimos()

	//Se imprime la lista minima, en este caso, si con el producto leche se llego a la existencia minima no imprimira nada ya que la lista se encuentra vacia.
	fmt.Println("\nNuestra lista de minimos")
	fmt.Println(listaMinima)

	//Se imprime la lista de productos.
	fmt.Println("\nNuestra lista de productos")
	fmt.Println(lProductos)

	//Ordenamos segun la cantidad
	lProductos.ordenarLista("cantidad")

	//Se imprime la lista de productos.
	fmt.Println("\nNuestra lista de productos ordenada por la cantidad menor a mayor")
	fmt.Println(lProductos)

	//Ordenamos segun el precio
	lProductos.ordenarLista("precio")

	//Se imprime la lista de productos.
	fmt.Println("\nNuestra lista de productos ordenada por el precio menor a mayor")
	fmt.Println(lProductos)

	//Ordenamos segun el nombre
	lProductos.ordenarLista("nombre")

	//Se imprime la lista de productos.
	fmt.Println("\nNuestra lista de productos ordenada por el nombre")
	fmt.Println(lProductos)
}

// si ha sobrado tiempo antes del receso, el ejercicio se podría ampliar para que los los productos se almacenen en archivos de texto
// que al inicio se carguen del archivo a la lista y que al final se actualice el archivo con el contenido de la lista

/*
IMPRESION:
[{arroz 15 2500} {frijoles 4 2000} {leche 8 1200} {café 12 4500}]
[{arroz 11 2500} {frijoles 4 2000} {leche 8 1200} {café 12 4500}]
[{arroz 11 2500} {frijoles 4 2000} {leche 8 1200} {café 12 4500} {azucar 20 1500}]
[{arroz 11 2500} {leche 8 1200} {café 12 4500} {azucar 20 1500}]
No se puede vender mayor cantidad de productos que los que hay en existencia

Nuestra lista de productos
[{arroz 11 2500} {leche 8 1200} {café 12 4500} {azucar 20 1500}]

Nuestra lista de minimos
[{leche 8 1200}]

El producto lech no esta en la lista de minimos.
En el sistema el producto leche tiene 8 unidades, incumpliendo con el minimo establecido de 10 unidades, Desea agregar existencias? Introduzca un 1 para si o un 0 para no.1

Introduzca las nuevas existencias que desea agrega para leche.10

El producto leche tiene 18 unidades, se cumplio con el minimo establecido de 10 unidades.
Nuestra lista de minimos
[]

Nuestra lista de productos
[{arroz 11 2500} {leche 18 1200} {café 12 4500} {azucar 20 1500}]

Nuestra lista de productos ordenada por la cantidad menor a mayor
[{arroz 11 2500} {café 12 4500} {leche 18 1200} {azucar 20 1500}]

Nuestra lista de productos ordenada por el precio menor a mayor
[{leche 18 1200} {azucar 20 1500} {arroz 11 2500} {café 12 4500}]

Nuestra lista de productos ordenada por el nombre
[{arroz 11 2500} {azucar 20 1500} {café 12 4500} {leche 18 1200}]
*/
