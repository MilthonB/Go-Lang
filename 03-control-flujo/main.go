package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hi world")

	//basicFor()
	fmt.Printf(ifElse(4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
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
	}

	return limit

}
