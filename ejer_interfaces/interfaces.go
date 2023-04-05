package ejer_interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/JoseVidalBetanM/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	scanner := bufio.NewScanner(os.Stdin)
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())

	fmt.Println("Edad: ")
	if scanner.Scan() {
		num1, _ := strconv.Atoi(scanner.Text())
		fmt.Println(num1)
	}

	fmt.Println("Altura: ")
	if scanner.Scan() {
		num2, _ := strconv.Atoi(scanner.Text())
		fmt.Println(num2)
	}

	fmt.Println("Peso(Kg): ")
	if scanner.Scan() {
		num3, _ := strconv.Atoi(scanner.Text())
		fmt.Println(num3)

	}

}
