package main

import (
	"fmt"
	"module/utils"

	"github.com/badoux/checkmail"
)

func main() {
	println("Hello World, modulo 01")
	utils.Write()
	utils.Write2()

	err := checkmail.ValidateFormat("charles@gmail")
	fmt.Println(err)

}
