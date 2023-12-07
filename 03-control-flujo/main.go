package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Hi world")

	//basicFor()
	fmt.Printf(ifElse(4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	loopFunction()
	switchSentence()
}

func basicFor() {

	var sum int
	// for clasico normal
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum = 0
	// for continue como si de un ciclo while se tratara
	for sum < 100 { // el for continue es un while
		sum += sum
	}

	sum = 0

	// Forever
	//for{} // loop infinito si omites los argumentos

	fmt.Printf("The value sum is: %v", sum)
}

func ifElse(x float64) string {

	if x < 0 {
		// llamas a la funcion para poner en negativo el valor de float
		return ifElse(-x) + "i"
	}

	// Retornas un string format
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, limit float64) float64 {

	// Se puede delcarar las variable dentro del if y luego evaular esa
	// misma variable

	if v := math.Pow(x, n); v < limit {
		return v
	} else { // Todas las variabls declaradas en los argumentos del if son disponibles en el bloque de else
		fmt.Printf("%g >= %g \n", v, limit)
	}

	return limit

}

func switchSentence() {
	fmt.Print("Go run on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}

	t := time.Now()
	switch { // un switch sin argumento es un switch true
	case t.Hour() > 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evenig")
	}

}

// Exercise Loops and Functions
func loopFunction() {

}
