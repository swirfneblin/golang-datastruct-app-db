package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	var array1 [5]string
	fmt.Println(array1)

	array2 := [5]string{"posicao 1", "posicao 2", "posicao 3", "posicao 4", "posicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada" //altera o valor do array2 e do slice2
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice2))
	fmt.Println(reflect.TypeOf(array2))

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho do slice
	fmt.Println(cap(slice3)) //capacidade do slice
}
