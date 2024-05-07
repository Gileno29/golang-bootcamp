package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 5.3
	x[1] = 1
	x[2] = 3
	x[4] = 5

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total = total + float64(x[i])

	}

	fmt.Println(total / 5)

}
