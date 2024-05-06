package main

import "fmt"

var a int = 20
var b string = "Nome"

func main() {
	var c float64 = float64(a)
	var d = c
	fmt.Println("o valor de C é:  o valor de B é: ", c, d)
}
