package main

import "fmt"

var n int

func init() {
	n = 10
	fmt.Println("Init")
}

func main() {
	fmt.Println("main")
	fmt.Println(n)
}
