package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v %v", e.when, e.what)
}
func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC), "O arquivo desapareceu",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
