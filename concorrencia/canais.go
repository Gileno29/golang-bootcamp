package main

import (
	"fmt"
	"time"
)

//modo de duas gorroutines sincronizarem e se comunicarem

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
