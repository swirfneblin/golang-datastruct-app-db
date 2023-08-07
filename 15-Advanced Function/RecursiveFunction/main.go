package main

import "fmt"

func fibonna(n uint) uint {
	if n <= 1 {
		return 1
	}

	return fibonna(n-1) + fibonna(n-2)
}

func main() {
	fmt.Println("Recursividade")
	posicao := uint(10)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonna(i))
	}

}
