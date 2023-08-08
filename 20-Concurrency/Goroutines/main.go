package main

import "time"

func main() {
	go escrever("Olá Mundo!")
	go escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		println(texto)
		time.Sleep(time.Second)
	}
}