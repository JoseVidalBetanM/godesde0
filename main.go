package main

import (
	"fmt"

	"github.com/JoseVidalBetanM/godesde0/variables"
)

func main() {

	estado, texto := variables.ConviertoaTexto(1500)
	fmt.Println(estado)
	fmt.Println(texto)

}
