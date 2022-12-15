package functions

import (
	"log"
	"strconv"
)

func Impedancia(resistencia_linha string, reatancia_linha string, impedancia_atual_str string) complex128 {
	var resistencia, _ = strconv.ParseFloat(resistencia_linha, 64)
	var reatancia, _ = strconv.ParseFloat(reatancia_linha, 64)
	var impedancia_atual, _ = strconv.ParseComplex(impedancia_atual_str, 128)

	impedancia := complex(resistencia, reatancia)

	if impedancia_atual != 0 {
		impedancia = (impedancia * impedancia_atual) / (impedancia + impedancia_atual)
	}

	return impedancia
}

func StringToFloat(grandeza_str string) float64 {
	grandeza, _ := strconv.ParseFloat(grandeza_str, 64)

	return grandeza
}

func ErrorValidade(err error) {
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
