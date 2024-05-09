//defer: escalona as funções
package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda")
}
func main() {
	defer dia1()
	dia2()

}
