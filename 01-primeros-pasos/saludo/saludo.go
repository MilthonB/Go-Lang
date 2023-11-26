package main

import (
	"example/hello"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("Saludar: ")
	log.SetFlags(0)

	mensaje, err := hello.Saludar("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mensaje)

}
