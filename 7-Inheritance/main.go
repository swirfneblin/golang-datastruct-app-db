package main

import "fmt"

type pessoa struct {
	nome      string
	sobreNome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)

	e2 := estudante{}
	e2.nome = "Maria"
	e2.sobreNome = "Silva"
	e2.idade = 25
	e2.altura = 165
	e2.curso = "Medicina"
	e2.faculdade = "USP"
	fmt.Println(e2)
}
