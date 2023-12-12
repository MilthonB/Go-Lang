package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	//fmt.Println("Hi world")

	//basicFor()
	//fmt.Printf(ifElse(4))
	//fmt.Println(
	//	pow(3, 2, 10),
	//		pow(3, 3, 20),
	//	)

	//	loopFunction()
	//	switchSentence()
	//  dDefer()
	//  pPointer()
	//	sStruct()
	//	aArray()
	//	mMap()
	// aAppen()
	rRange()
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

func dDefer() {

	defer fmt.Println("1") // Cuarto defer detiene la linea y se brinca a la siguiente hasta completar el ciclo de
	defer fmt.Println("2") // Tercero defer detiene la linea y se brinca a la siguiente hasta completar el ciclo de
	defer fmt.Println("3") // Segundo defer detiene la linea y se brinca a la siguiente hasta completar el ciclo de
	defer fmt.Println("4") // Primero defer detiene la linea y se brinca a la siguiente hasta completar el ciclo de

	for i := 0; i < 10; i++ {
		fmt.Println("numero: ", i)
	}

	fmt.Println("World")

}

func pPointer() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	*p = *p / 21
	fmt.Println(j)

}

func sStruct() {
	data := dataPersonal{ // los vcalores pueden ser inferidos por el lenguague si no pones todos los valores de la estructura entonces este los toma por default
		name:    "Milthon",
		apllido: "Borboa",
		edad:    30,
	}

	fmt.Println(data)

}

type dataPersonal struct {
	name    string
	apllido string
	edad    int
}

func aArray() {

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	prime := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(prime)

	var sslice []int = prime[1:4]
	fmt.Println(sslice)

	names := [4]string{
		"Milthon",
		"Orlando",
		"Felicia",
		"JD",
	}

	fmt.Println(names)

	e := names[0:2]
	b := names[1:3]
	fmt.Println(e, b)

	b[0] = "XXXX" // se pasa como referencia los valores del array al array principal
	fmt.Println(b)
	fmt.Println(names)

	s := []struct {
		i int
		j int
	}{
		{i: 1, j: 2},
		{i: 2, j: 3},
		{i: 3, j: 4},
		{i: 4, j: 5},
		{i: 5, j: 6},
	}

	fmt.Println(s)

	//	var arrayliteral [10]string // Esto es un array literal

	//	var arrayliteral2 []string // Esto es un arreglo no literal

	// Tipos de slices [0:1], [:1] [1:]

	sss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice1 := sss[1:]
	fmt.Println(slice1)

	slice2 := sss[:2]
	fmt.Println(slice2)

	slice3 := sss[:1]
	fmt.Println(slice3)

	var slice4 []int
	fmt.Println(len(slice4), cap(slice4), slice4)

	if slice4 == nil {
		fmt.Println("nil!!")
	}

}

func mMap() {

	// var a = [10]int la capacidad es de 10
	// var a = [10]int {1,2,3,4,5} la capacidad es de 10 pero tiene de 4 elementos len()

	//a := make([]int, 0, 5) // ( el tipo de lo que vas crear, la longuitud, y la capacidad )
	a := make([]int, 5) // ( el tipo de lo que vas crear, la longuitud, y la capacidad ) len por default es 5 elementos de valores 0
	printSlice("a", a)

	b := make([]int, 0, 5) // el len es 0 no tiene ningun valor
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func aAppen() {

	var a []int
	printSlice1(a)

	// append(arreglo, elemento)
	a = append(a, 1)
	printSlice1(a)

	a = append(a, 2, 3, 4, 5, 6)
	printSlice1(a)

}

func rRange() {

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// range for ( i = index, v = value )
	for i, v := range pow {
		fmt.Println("2**%d = %d", i, v)
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice1(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
