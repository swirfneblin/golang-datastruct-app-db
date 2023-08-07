package main

import "fmt"

// func invertSignal(number int) int {
// 	return number * -1
// }

func invertSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	numero := 20
	fmt.Println(numero)
	invertSignalWithPointer(&numero)
	fmt.Println(numero)
}
