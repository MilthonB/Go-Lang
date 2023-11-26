package hello

import "fmt"

func main() {

	fmt.Println("Hello world!")

}

func Saludar(name string) string {

	mensaje := " Hola, mi nombre es: " + name

	return mensaje

}
