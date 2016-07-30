package main

import (
	"fmt"
	"os"

	a1 "github.com/curso-go/sintaxe"
)

func aula1() {
	var word = "Webschool"
	word = a1.Hello(word)
	fmt.Println(word)
}

func main() {

	var aulas = []string{"Sintaxe"}

	var args = os.Args[1:]
	switch args[0] {
	case "aula1":
		fmt.Println(":: Aula de Sintaxe ::")
		aula1()
	default:
		fmt.Println("Aulas")
		fmt.Println(aulas)
	}
}
