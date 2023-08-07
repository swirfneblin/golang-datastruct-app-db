package main

import "fmt"

func Closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	fmt.Println("Closure")

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := Closure()

	funcaoNova()

}
