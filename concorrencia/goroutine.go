package main

import "fmt"

//funcao capaz de executar outras funcoes de modo concorrente

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
func main() {
	go f(0)
	var escreva string
	fmt.Scanln(&escreva)

}
