package main

import "fmt"

func soma(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}
	return total

}
func main() {
	x := soma(1, 2, 3)

	fmt.Println(x)
}
