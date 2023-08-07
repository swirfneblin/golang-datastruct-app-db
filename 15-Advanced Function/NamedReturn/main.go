package main

import "fmt"

func calculosMatematicos(n1 int8, n2 int8) (soma int8, subtracao int8) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
