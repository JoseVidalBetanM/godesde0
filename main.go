package main

import (
	"github.com/JoseVidalBetanM/godesde0/files"
)

func main() {

	/*estado, texto := variables.ConviertoaTexto(123)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("esto no en Windows, es", os)
	} else {
		fmt.Println("esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("esto es linux")
	case "Darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejercicios.Ejercicio("150")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.OperaMatematica()

	iteraciones.Ciclos() */

	// fmt. Println(ejercicios.MiTabla())

	// files.GrabaTabla()

	// files.Append()

	files.LeoArchivo()

}
