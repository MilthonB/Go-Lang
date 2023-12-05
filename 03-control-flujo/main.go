package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hi world")

	basicFor()

}

func basicFor() {

	var sum int

	for i := 0; i < 10; i++ {
		sum += i
	}

	sum = 0
	// for continue
	for sum < 100 { // el for continue es un while
		sum += sum
	}

	sum = 0

	// Forever
	//for{} // loop infinito si omites los argumentos

	fmt.Printf("The value sum is: %v", sum)
}
