package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0
	//while
	fmt.Println("Loops 1 using WHILE")
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	//for
	fmt.Println("Loops 2 using FOR")
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	//for range
	fmt.Println("Loops 3 using FOR RANGE")
	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//for range sem indice
	fmt.Println("Loops 4 using FOR RANGE without index")
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//for range em string
	fmt.Println("Loops 5 using FOR RANGE in string")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	//for range em map
	fmt.Println("Loops 6 using FOR RANGE in map")
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//for infinito
	fmt.Println("Loops 7 using FOR INFINITO")
	for {
		fmt.Println("Loops infinitos...")
		time.Sleep(time.Second)
	}
}
