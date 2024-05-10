package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

type paraNome []Dados

func (ps paraNome) Len() int {
	return len(ps)

}

func (ps paraNome) Less(i, j int) bool { // item na posição i é menor do que item na posição J
	return ps[i].Nome < ps[j].Nome

}

func (ps paraNome) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func main() {

	criancas := []Dados{
		{"julia", 5},
		{"Joao", 10},
	}

	sort.Sort(paraNome(criancas))
	fmt.Println(criancas)

}
