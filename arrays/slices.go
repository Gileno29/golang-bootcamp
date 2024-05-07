package main

import "fmt"

func main() {

	fatia := make([]float64, 5)
	fmt.Println(fatia)

	//SEGUNDA FORMA
	arr := [7]float64{1, 2, 3, 45, 6, 7}
	fatia02 := arr[0:4]

	//Acrescentar elemento no slice
	fatia03 := append(fatia02, 3, 2, 3)

	fmt.Println(fatia02, fatia03)

	//Copiando fatias
	fatia04 := make([]float64, 2)

	copy(fatia04, fatia03)

	fmt.Println(fatia03, fatia04)

}
