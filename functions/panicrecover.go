//panic: ero do programador/erro de execução
//recover: ela interrompe o panic

package main

import "fmt"

func main() {

	defer func() {
		x := recover()

		fmt.Println(x)
	}()

	panic("Panico!!")

}
