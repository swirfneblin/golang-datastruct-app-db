package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	var str string = "texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)

	var boleano bool
	fmt.Println(boleano)

	var err = errors.New("Erro interno")
	if err != nil {
		fmt.Println(err)
	}

}
