package main

import "fmt"

func recuperarExecucaoFuncao() {
	fmt.Println("Tentativa de recuperar a execução da função")
	if r := recover(); r != nil {
		fmt.Println("Função recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucaoFuncao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MEDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println("Panic e Recover")

	fmt.Println(alunoEstaAprovado(3, 6))
	fmt.Println("Pos execução!")

}
