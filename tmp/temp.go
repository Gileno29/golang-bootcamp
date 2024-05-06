package main

import "fmt"

//declaracao d avariavel CONST do ponto de ebulicao da agua em F
const ebulicaoF float64=212.0 

func main(){

	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32)*5/9 //conversao de F para C

	fmt.Println("A temperatura de ebulicao da agaua em °F", tempF)
	fmt.Println("A temperatura de ebulição da agua em °C = ", tempC)







}
