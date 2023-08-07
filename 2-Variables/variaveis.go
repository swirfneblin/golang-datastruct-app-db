package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1, variavel2)

	variavel1, variavel2 = variavel2, variavel1
	fmt.Println(variavel1, variavel2)

}
