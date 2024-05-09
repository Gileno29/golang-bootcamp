package main

import "fmt"

func inicial(x *int) {
	*x = 0
}

func main() {
	x := 5
	inicial(&x)
	fmt.Println(x)
}
