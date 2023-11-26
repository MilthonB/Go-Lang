package main

import (
	"example/hello"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("Saludar: ")
	log.SetFlags(0)

	names := []string{"Milthon", "Milthon2", "Milthon3"}
	mensajes, err := hello.Saludador(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mensajes)

}
