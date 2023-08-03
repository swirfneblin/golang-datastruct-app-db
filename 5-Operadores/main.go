package main

import "fmt"

func main() {
	// Operadores aritmeticos
	// + - / * %
	// Operadores de atribuição
	// = += -= *= /= %=
	// Operadores de comparação
	// == != > >= < <=
	// Operadores Logicos
	// && || !
	// Operadores Unarios
	// ++ --
	// Operadores Ternarios
	// ?:

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	var variavel1 string = "String 1"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

}
