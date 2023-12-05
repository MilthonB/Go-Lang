package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

func main() {

	fmt.Println("Welcome to the playground")
	fmt.Println("The time is: ", time.Now())

	var nombre string = "Hi iam"

	fmt.Printf("Now you have %g problem. \n", math.Sqrt(7))
	fmt.Printf("Now you have %v problem. \n", true)
	fmt.Printf("Now you have %v problem. \n", nombre)
	fmt.Printf("Now you have %#v problem. \n", nombre)
	fmt.Printf("Now you have %q  problem. \n", nombre)

	//The default format for %v is:
	/*
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	*/

	x, y := suma(2, 4)

	fmt.Printf("The value x is: %v \nThe value y is: %v \n", x, y)
	fmt.Println(split(17))

	variables()
	basicType()
	valueZero()
	typeConvertion()
	inferenceType()
}

func variables() {
	// declaracionde variables como una lista y al final va el tipo de la variable
	var c, python, java bool // Cuando no lo declaras con un valor tiene el valor de false por defaut
	var i int                // Cuando no lo declaras con un valor tiene el valor de 0 por default
	var x, j int = 2, 4      // Declaracion de vairbales inicializadas
	a, b := 6, 8             // Variables declaraas de forma corta y rapida
	fmt.Printf("Lenguajes: %v, %v, %v, %v, %v, %v, %v, %v", c, python, java, i, x, j, a, b)
}

func suma(x, y int) (int, int) { // se pueden retornar variis valores
	return x * 2, y * 2
}

func split(sum int) (x, y int) { // Si en el metodo esta definido variable de retorno, estas se enviaran cuando declares la palabra reservada return

	x = sum * 4 / 9
	y = sum - x
	return
}

func basicType() {

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("\nType: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)

}

func valueZero() {
	var i int     // valor por default 0
	var f float64 // valor por default 0
	var b bool    // valor por defaultfalse
	var s string  // valor por default ""

	fmt.Printf("Value Zero: %v %v %v %q \n", i, f, b, s)
}

func typeConvertion() {
	var x, y int = 2, 4
	var f float64 = float64(x*x + y*y)
	var z int = int(f)
	fmt.Println(x, y, f, z)
}

func inferenceType() {
	// Las variables si no son declaradas con su tipo de variable estas toman por inferencia el valor que tiene a sue derecha

	var j int = 1
	var i = j

	// Pueden ser character, string, boolean o numeric value
	const World = "Mexico" // las cosntantes son con mayusculas
	const Truth = true

	fmt.Println(World, Truth, i, j)

}
