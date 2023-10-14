package main

import (
	"fmt"
	"runtime"
	//"github.com/JoseVidalBetanM/godesde0/variables"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(100)
	fmt.Println(estado)
	fmt.Println(texto)*/
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("esto no es windows")
	} else {
		fmt.Println("esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux:":
		fmt.Println("esto es linux")
	case "Darwin":
		fmt.Println("esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

}
