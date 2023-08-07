package main

import (
	"errors"
	"fmt"
)

func diaDaSemana(numero int8) (string, error) {
	switch numero {
	case 1:
		return "Domingo", nil
	case 2:
		return "Segunda", nil
	case 3:
		return "Terça", nil
	case 4:
		return "Quarta", nil
	case 5:
		return "Quinta", nil
	case 6:
		return "Sexta", nil
	case 7:
		return "Sabado", nil
	default:
		return "", errors.New("numero invalido")
	}
}

func diaDaSemana2(numero int8) string {
	diaDaSemana := ""
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Numero invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia, err := diaDaSemana(12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dia)
	}

	dia = diaDaSemana2(2)
	fmt.Println(dia)
}
