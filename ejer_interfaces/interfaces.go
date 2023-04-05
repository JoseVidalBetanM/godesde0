package ejer_interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/JoseVidalBetanM/godesde0/interfaces"
)

var num1 int
var num2 int
var num3 int
var nombre string
var err error

func HumanosRespirando(hu interfaces.Humano) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Nombre: ")
	if scanner.Scan() {
		nombre = scanner.Text()

	}

	fmt.Println("Edad: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("hubo un error " + err.Error())
		}
	}

	fmt.Println("Altura: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("hubo un error " + err.Error())
		}
	}

	fmt.Println("Peso(Kg): ")
	if scanner.Scan() {
		num3, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("hubo un error " + err.Error())
		}

	}

	hu.Respirar()
	fmt.Printf("Soy un/a %s, me llamo %s , tengo %d años de edad, mido %d de altura, peso %d y estoy respirando \n", hu.Sexo(), nombre, num1, num2, num3)

}
