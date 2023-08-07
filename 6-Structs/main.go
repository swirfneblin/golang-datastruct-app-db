package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")

	end := endereco{"Rua dos Bobos", 0}

	var u usuario
	u.nome = "Charles"
	u.idade = 20
	fmt.Println(u)

	u2 := usuario{"Davi", 25, end}
	fmt.Println(u2)

	usuario3 := usuario{idade: 23, endereco: end}
	fmt.Println(usuario3)
}
