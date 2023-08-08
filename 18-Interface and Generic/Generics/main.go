package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println("Fim!", 12, 3, "teste", float32(1231))

	// not recommended
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(123),
	}

	fmt.Println(mapa)
}
