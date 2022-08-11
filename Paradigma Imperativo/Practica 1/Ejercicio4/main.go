package main

import "fmt"

// Nuestra estructura para el programa
type calzado struct {
	marca    string
	precio   int
	talla    int
	cantidad int
}

// Nuestro slice para el calzado
var lista_calzado_slice []calzado

// Funcion para crear el calzado
func crearCalzado(m string, p int, t int, c int) calzado {
	cal := calzado{marca: m, precio: p, talla: t, cantidad: c}
	return cal
}

// Funcion para agregar el calzado
func agregarCalzado(c calzado) {
	lista_calzado_slice = append(lista_calzado_slice, c)
}

// Esta funcion creara el nuevo calzado por medio de filtros
func nuevoCalzado(m string, p int, t int, c int) {
	var contador int = 0

	//Este for revisara de calzado en calzado para verificar si ya se habia agregado anteriormente uno
	//del mismo modelo o marca y talla para evitar la duplicacion de datos
	for {
		if contador == len(lista_calzado_slice) {
			break
		}
		if lista_calzado_slice[contador].marca == m && lista_calzado_slice[contador].talla == t {
			fmt.Printf("\nYa el calzado: %s de talla: %d habia sido agregado anteriormente, agrega uno nuevo.", m, t)
			return
		}
		contador++
	}

	//Este if verificara si la talla del calzado no es menor que 34 o mayor que 44, si se cumplen ambos
	//se podra agregar el calzado
	if t < 34 || t > 44 {
		fmt.Printf("\nEl calzado: %s de talla: %d no puede ser agregado, solo se permiten tallas de 34 al 44", m, t)
	} else {
		agregarCalzado(crearCalzado(m, p, t, c))
	}
}

// Funcion para imprimir el calzado del slice
func recorrerSlice(slice []calzado) {
	var contador int = 0
	fmt.Printf("\n\nNuestro calzado:")
	for {
		if contador == len(slice) {
			fmt.Printf("\n")
			break
		}
		fmt.Printf("\nMarca: %s Precio: %d Talla: %d Cantidad: %d",
			lista_calzado_slice[contador].marca, lista_calzado_slice[contador].precio,
			lista_calzado_slice[contador].talla,
			lista_calzado_slice[contador].cantidad)
		contador++
	}
}

// Funcion para vender calzado
func venderCalzado(m string, t int, cantidadVendida int) {
	var contador int = 0

	//Gracias a un for y un contador podemos recorrer el slice para verificar si existe un calzado
	//por medio de su marca y talla, a la vez se puede ver si tiene suficientes unidades para vender
	for {
		//Si el if se cumple es que no se encuentra el calzado
		if contador == len(lista_calzado_slice) {
			fmt.Printf("\nEl calzado: %s de talla: %d no se encuentra en la base de datos, no se puede vender.", m, t)
			break
		}
		//En este if se comprobara si existe la marca y la talla que se desea vender
		if lista_calzado_slice[contador].marca == m && lista_calzado_slice[contador].talla == t {
			//Si este if se cumple no se puede vender debido a que no hay unidades disponibles
			if lista_calzado_slice[contador].cantidad == 0 {
				fmt.Printf("\nEl stock del calzado: %s de talla: %d es 0, no se puede vender.", m, t)
				break
			}

			//en caso de que haya almenos una unidad se procedera a verificar si se pueden vender
			//segun la cantidad solicitada
			var comprobarCalzado int = lista_calzado_slice[contador].cantidad - cantidadVendida

			//En este if de comprobarCalzado si el resultado es negativo no se podra vender
			if comprobarCalzado < 0 {
				fmt.Printf("\nEl stock del calzado: %s de talla: %d es menor a la cantidad que se desea comprar, no se puede vender.", m, t)
				break
			} else {

				//en caso contrario se actualizara el calzado
				lista_calzado_slice[contador].cantidad = comprobarCalzado

				fmt.Printf("\nEl precio del calzado: %s de talla: %d es de %d colones, su total por %d par(es) es de %d colones.", m, t, lista_calzado_slice[contador].precio, cantidadVendida, cantidadVendida*lista_calzado_slice[contador].precio)
				break
			}
		}
		contador++
	}
}

// Funcion para agregar nuevas unidades de calzado
func agregarUnidadesCalzado(m string, t int, nuevaCantidad int) {
	var contador int = 0

	//Gracias a un for y un contador podemos recorrer el slice para verificar si existe un calzado
	//por medio de su marca y talla
	for {
		//Si el if se cumple es que no se encuentra el calzado
		if contador == len(lista_calzado_slice) {
			fmt.Printf("\nEl calzado: %s de talla: %d no se encuentra en la base de datos, agregue primero el calzado para agregar nuevas unidades.", m, t)
			break
		}
		//En este if se comprobara si existe la marca y la talla que se desea agregar unidades
		//si existe se procedera a agregar las nuevas unidades
		if lista_calzado_slice[contador].marca == m && lista_calzado_slice[contador].talla == t {
			lista_calzado_slice[contador].cantidad = nuevaCantidad
			fmt.Printf("\nEl nuevo stock del calzado: %s de talla: %d es de %d unidad(es).", m, t, lista_calzado_slice[contador].cantidad)
			break
		}
		contador++
	}
}

func main() {
	//Cabe aclarar que para agregar el calzado se debe usar la funcion nuevoCalzado, ya que aqui se
	//encuentran los filtros para evitar la duplicacion de datos

	//Agregamos nuestro calzado
	nuevoCalzado("Prada", 3000, 34, 2)
	nuevoCalzado("Nike", 700, 44, 4)
	nuevoCalzado("Toms", 3444, 42, 6)
	nuevoCalzado("Reebok", 5788, 35, 7)
	nuevoCalzado("Converse", 67544, 36, 3)
	nuevoCalzado("Kickers", 54556, 37, 4)
	nuevoCalzado("Timberland", 35565, 39, 7)
	nuevoCalzado("Puma", 3455, 43, 40)
	nuevoCalzado("Lotto", 5433, 43, 41)

	//Pueden existir marcas o modelos iguales, pero, solo de diferente talla
	nuevoCalzado("Adidas", 6000, 40, 1)
	nuevoCalzado("Adidas", 3500, 42, 5)
	nuevoCalzado("Adidas", 4500, 44, 6)

	//En este caso el programa no agregara este calzado al slice ya que ya fue agregado anteriormente
	nuevoCalzado("Adidas", 5343, 40, 4)
	nuevoCalzado("Adidas", 6545, 42, 4)
	nuevoCalzado("Adidas", 68734, 44, 3)

	//En este caso el programa no agregara este calzado al slice, ya que las tallas no se aceptan
	nuevoCalzado("Adidas", 8453, 33, 4)
	nuevoCalzado("Adidas", 6756, 20, 4)
	nuevoCalzado("Adidas", 7464, 45, 4)

	//Se ve el calzado existente
	recorrerSlice(lista_calzado_slice)

	//Vendemos calzado hasta se quede en 0 de stock
	venderCalzado("Adidas", 44, 6)
	venderCalzado("Prada", 34, 2)
	venderCalzado("Converse", 36, 3)

	//Se ve el calzado existente despues de la venta
	recorrerSlice(lista_calzado_slice)

	//Tratamos de vender este mismo calzado pero sin existencias
	venderCalzado("Adidas", 44, 6)
	venderCalzado("Prada", 34, 2)
	venderCalzado("Converse", 36, 3)

	//Tratamos de vender este calzado que no existe en la base de datos
	venderCalzado("New Balance", 44, 1)
	venderCalzado("Victoria", 33, 4)
	venderCalzado("Lodi", 36, 6)

	//Actualizamos el nuevo calzado
	agregarUnidadesCalzado("Adidas", 44, 2)
	agregarUnidadesCalzado("Prada", 34, 5)
	agregarUnidadesCalzado("Converse", 36, 4)

	//Vendemos una unidad de este mismo calzado
	venderCalzado("Adidas", 44, 1)
	venderCalzado("Prada", 34, 1)
	venderCalzado("Converse", 36, 1)

	//Se ve el calzado existente despues de agregar nuevas unidades
	recorrerSlice(lista_calzado_slice)
}

/*
RESULTADO:

Ya el calzado: Adidas de talla: 40 habia sido agregado anteriormente, agrega uno nuevo.
Ya el calzado: Adidas de talla: 42 habia sido agregado anteriormente, agrega uno nuevo.
Ya el calzado: Adidas de talla: 44 habia sido agregado anteriormente, agrega uno nuevo.
El calzado: Adidas de talla: 33 no puede ser agregado, solo se permiten tallas de 34 al 44
El calzado: Adidas de talla: 20 no puede ser agregado, solo se permiten tallas de 34 al 44
El calzado: Adidas de talla: 45 no puede ser agregado, solo se permiten tallas de 34 al 44

Nuestro calzado:
Marca: Prada Precio: 3000 Talla: 34 Cantidad: 2
Marca: Nike Precio: 700 Talla: 44 Cantidad: 4
Marca: Toms Precio: 3444 Talla: 42 Cantidad: 6
Marca: Reebok Precio: 5788 Talla: 35 Cantidad: 7
Marca: Converse Precio: 67544 Talla: 36 Cantidad: 3
Marca: Kickers Precio: 54556 Talla: 37 Cantidad: 4
Marca: Timberland Precio: 35565 Talla: 39 Cantidad: 7
Marca: Puma Precio: 3455 Talla: 43 Cantidad: 40
Marca: Lotto Precio: 5433 Talla: 43 Cantidad: 41
Marca: Adidas Precio: 6000 Talla: 40 Cantidad: 1
Marca: Adidas Precio: 3500 Talla: 42 Cantidad: 5
Marca: Adidas Precio: 4500 Talla: 44 Cantidad: 6

El precio del calzado: Adidas de talla: 44 es de 4500 colones, su total por 6 par(es) es de 27000 colones.
El precio del calzado: Prada de talla: 34 es de 3000 colones, su total por 2 par(es) es de 6000 colones.
El precio del calzado: Converse de talla: 36 es de 67544 colones, su total por 3 par(es) es de 202632 colones.

Nuestro calzado:
Marca: Prada Precio: 3000 Talla: 34 Cantidad: 0
Marca: Nike Precio: 700 Talla: 44 Cantidad: 4
Marca: Toms Precio: 3444 Talla: 42 Cantidad: 6
Marca: Reebok Precio: 5788 Talla: 35 Cantidad: 7
Marca: Converse Precio: 67544 Talla: 36 Cantidad: 0
Marca: Kickers Precio: 54556 Talla: 37 Cantidad: 4
Marca: Timberland Precio: 35565 Talla: 39 Cantidad: 7
Marca: Puma Precio: 3455 Talla: 43 Cantidad: 40
Marca: Lotto Precio: 5433 Talla: 43 Cantidad: 41
Marca: Adidas Precio: 6000 Talla: 40 Cantidad: 1
Marca: Adidas Precio: 3500 Talla: 42 Cantidad: 5
Marca: Adidas Precio: 4500 Talla: 44 Cantidad: 0

El stock del calzado: Adidas de talla: 44 es 0, no se puede vender.
El stock del calzado: Prada de talla: 34 es 0, no se puede vender.
El stock del calzado: Converse de talla: 36 es 0, no se puede vender.
El calzado: New Balance de talla: 44 no se encuentra en la base de datos, no se puede vender.
El calzado: Victoria de talla: 33 no se encuentra en la base de datos, no se puede vender.
El calzado: Lodi de talla: 36 no se encuentra en la base de datos, no se puede vender.
El nuevo stock del calzado: Adidas de talla: 44 es de 2 unidad(es).
El nuevo stock del calzado: Prada de talla: 34 es de 5 unidad(es).
El nuevo stock del calzado: Converse de talla: 36 es de 4 unidad(es).
El precio del calzado: Adidas de talla: 44 es de 4500 colones, su total por 1 par(es) es de 4500 colones.
El precio del calzado: Prada de talla: 34 es de 3000 colones, su total por 1 par(es) es de 3000 colones.
El precio del calzado: Converse de talla: 36 es de 67544 colones, su total por 1 par(es) es de 67544 colones.

Nuestro calzado:
Marca: Prada Precio: 3000 Talla: 34 Cantidad: 4
Marca: Nike Precio: 700 Talla: 44 Cantidad: 4
Marca: Toms Precio: 3444 Talla: 42 Cantidad: 6
Marca: Reebok Precio: 5788 Talla: 35 Cantidad: 7
Marca: Converse Precio: 67544 Talla: 36 Cantidad: 3
Marca: Kickers Precio: 54556 Talla: 37 Cantidad: 4
Marca: Timberland Precio: 35565 Talla: 39 Cantidad: 7
Marca: Puma Precio: 3455 Talla: 43 Cantidad: 40
Marca: Lotto Precio: 5433 Talla: 43 Cantidad: 41
Marca: Adidas Precio: 6000 Talla: 40 Cantidad: 1
Marca: Adidas Precio: 3500 Talla: 42 Cantidad: 5
Marca: Adidas Precio: 4500 Talla: 44 Cantidad: 1

*/
