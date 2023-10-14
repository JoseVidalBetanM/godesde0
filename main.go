package main

import (
	"fmt"

	"github.com/JoseVidalBetanM/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(100)
	fmt.Println(estado)
	fmt.Println(texto)

}
