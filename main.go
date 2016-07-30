package main

import (
	"fmt"
	"os"

	a1 "github.com/curso-go/sintaxe"
	a2 "github.com/curso-go/tipos"
)

func aula1() {
	var word = "Webschool"
	word = a1.Hello(word)
	fmt.Println(word)
}

func aula2(t string) {
	switch t {
	case "int":
		fmt.Println(a2.ShowType(0))
		break

	case "string":
		fmt.Println(a2.ShowType(""))
		break

	case "bool":
		fmt.Println(a2.ShowType(true))
		break
	}
}

func main() {

	var aulas = []string{"Sintaxe"}

	var args = os.Args[1:]
	switch args[0] {
	case "aula1":
		fmt.Println(":: Aula de Sintaxe ::")
		aula1()
	case "aula2":
		fmt.Println(":: Aula de Tipos ::")
		aula2(args[1])
	default:
		fmt.Println("Aulas")
		fmt.Println(aulas)
	}
}
