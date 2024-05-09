package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if _, err := io.WriteString(os.Stdout, "Hello world"); err != nil {
		log.Printf("%+v\n", err)
	}

	mesage := []byte("Hello, Golphers!")
	err := ioutil.WriteFile("Hello", mesage, 0644)

	if err != nil {

	}

	fmt.Printf("%s", bytes.Title([]byte("her royal highness")))

}
