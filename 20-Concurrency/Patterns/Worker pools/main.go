package main

func fibonna(n uint) uint {
	if n <= 1 {
		return 1
	}

	return fibonna(n-1) + fibonna(n-2)
}

func main() {
	tarefas := make(chan uint, 45)
	resultados := make(chan uint, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := uint(0); i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := uint(0); i < 45; i++ {
		resultado := <-resultados
		println(resultado)
	}
}

func worker(tarefas <-chan uint, resultados chan<- uint) {
	for numero := range tarefas {
		resultados <- fibonna(numero)
	}
}
