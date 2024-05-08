//funcoes também são chamadas de procedimento ou sub-rotina
//parte código
//ela pega um dado de entrada e dá um dado de saída
package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0

	for _, valor := range lista {
		total += valor
	}

	return total / float64(len(lista))
}

func main() { //Programa que calcula a média de uma sala
	// lista de notas

	lista := []float64{98, 93, 77, 82}

	//total := 0.0

	fmt.Printf("Media: %.1f", media(lista))

}
