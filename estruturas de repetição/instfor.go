package main

import "fmt"

func main() {

	//exemplos de for

	fmt.Println("============EXEMPLOS DE DE LACO FOR================")
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//Exemplos de if

	fmt.Println("============EXEMPLOS DE DE LACO IF================")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

}
