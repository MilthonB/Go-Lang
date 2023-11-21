package main

import (
	"fmt"
)

func main() {
	mensaje := Hola("Milthon")
	print(fmt.Println(mensaje))
}

func Hola(name string) string {

	// var mensaje string
	// := declara y inicializa una variable
	mensaje := fmt.Sprintf("Hola, %v. Bienvenido!", name)

	return mensaje
}
