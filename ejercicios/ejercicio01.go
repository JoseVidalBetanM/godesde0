package ejercicios

import (
	"strconv"
)

func Ejercicio(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "hubo un error " + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
