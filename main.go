package main

import (
	"fmt"

	"github.com/JoseVidalBetanM/godesde0/goroutines"
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

	// files.LeoArchivo()

	// funciones.Calculos()

	// funciones.LlamarClosure()

	// funciones.Exponencia(2)

	// arreglos_slices.MuestroArreglos()

	// arreglos_slices.MuestroSlices()

	// arreglos_slices.Capacidad()

	// mapas.Mapitas()

	// mapas.Campeonato()

	// users.AltaUsuario()

	/*
		Pedro := new(modelos.Hombre)
		e.HumanosRespirando(Pedro)

		Maria := new(modelos.Mujer)
		e.HumanosRespirando(Maria)
	*/

	// defer_panic.Paniquito()

	canal1 := make(chan bool)
	go goroutines.Velocito("Jose Vidal", canal1)

	fmt.Println("estoy aqui")
	var x string
	fmt.Scanln(&x)

	<-canal1

}
