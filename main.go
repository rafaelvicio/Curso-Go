package main

import "fmt"

func Veriadicas(e ...int) {
	fmt.Println(e)
}

func main() {
	Veriadicas(1, 2, 3, 4, 5, 6, 7, 8, 9)
}
