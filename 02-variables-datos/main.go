package main

import (
	"fmt"
	"math"
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
}
