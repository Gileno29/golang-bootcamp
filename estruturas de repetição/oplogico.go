package main

import "fmt"

func main() {
	x := 2

	if x == 2 || x == 3 {
		fmt.Println("Sim, x é 2 ou 3!")
	} else {
		fmt.Println("X não nem 2 nem 3")
	}

}
