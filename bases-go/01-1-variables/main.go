package main

import "fmt"

func main() {
	fmt.Println("Variables")

	/*
		Valores cero de las variables
		Numeros = 0
		String = 0
		Punteros = null
		Interfaces = null
	*/

	var myIntVar int
	myIntVar = -12
	fmt.Println("My int var is:", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("My uint var is:", myUintVar)

	var myStringVar string
	myStringVar = "I'm a string"
	fmt.Println("My string var is:", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My bool var is:", myBoolVar)

	newIntVar := 34
	fmt.Println("my new int var:", newIntVar)

	// Constantes

	const myFirstConst = "a12"
	const myIntConst int = 1289
	const myStringConst string = "aaa"
	fmt.Printf("string: %s, int: %d, string: %s \n", myFirstConst, myIntConst, myStringConst)

	//Verbos variables -> %s - string, %f - float, %d - int, %v - default value, %T - type value

	myStringMultiLine := `Hola
	mundo!
	Soy multilinea`

	fmt.Printf("%s \n", myStringMultiLine)

	fmt.Println("\n Cast de Variables")
}
