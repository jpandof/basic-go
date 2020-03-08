package main

import "fmt"

func main() {

	///////////////////////////////////////
	// Variables
	// Se pueden definir de varias maneras
	//////////////////////////////////////

	// Explicitamente Sin inicializar
	var i1 int
	var i2 bool
	var i3 string
	//var array = [4]int{9, 8, 7, 6} // array, estructura de datos con tamaño fijado
	//var slice = []int{9, 8, 7, 6}  // slice, es una implementación de array, permite que el tamaño sea modificable
	var mySlice []int // slice, es una implementación de array, permite que el tamaño sea modificable
	var myMap map[string]int

	// Explicitamente inicializadas
	var a string = "inicial" // se puede omitir "string" ya que por defecto al estar entre comillas lo interpreta como string

	// Varias a la vez
	var b1, c1 int = 1, 2
	var (
		b2 bool = false // se puede omitir "bool" ya que por defecto al utilizar la palabra 'true' o 'false' lo interpretra como boolean
		c2      = "rojo"
	)

	// Declaración e inicialización corta de variables.
	myVar1 := 10
	myVar2, myVar3 := "verde", false

	fmt.Printf("i1 = %v \n", i1)
	fmt.Printf("i2 = %v \n", i2)
	fmt.Printf("i3 = %v \n", i3)
	fmt.Printf("mySlice = %v \n", mySlice)
	fmt.Printf("myMap = %v \n", myMap)
	fmt.Printf("a = %v \n", a)
	fmt.Printf("b1 = %v \n", b1)
	fmt.Printf("c1 = %v \n", c1)
	fmt.Printf("b2 = %v \n", b2)
	fmt.Printf("c2 = %v \n", c2)
	fmt.Printf("myVar1 = %v \n", myVar1)
	fmt.Printf("myVar2 = %v \n", myVar2)
	fmt.Printf("myVar3 = %v \n", myVar3)

	fmt.Printf("Is null a slice = %v \n", mySlice == nil)
	fmt.Printf("Is null a map = %v \n", myMap == nil)

	element, existKey := myMap["myElement"]

	fmt.Printf("Is 0 myElement = %v \n", element == 0)
	fmt.Printf("Exist the key myElement = %v \n", existKey)

}
