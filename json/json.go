package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Usuarios []User `json:"usuarios"`
}

type User struct {
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Idade  int    `json:"idade"`
	Social Social `json:"rede social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twiter   string `json:"twitter"`
}

func main() {

	jsonFile, err := os.Open("./usuarios.json")

	if err != nil {
		fmt.Println(err)

	}

	fmt.Print("Arquivo aberto com sucesso!!")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var usuarios Users
	json.Unmarshal(byteValue, &usuarios)

	fmt.Println(usuarios)
	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Usuario Tipo: " + usuarios.Usuarios[i].Tipo)
		fmt.Println("Usuario idade: " + strconv.Itoa(usuarios.Usuarios[i].Idade))
		fmt.Println("Usuario Nome: " + usuarios.Usuarios[i].Nome)

	}

}
