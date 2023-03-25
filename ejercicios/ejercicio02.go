package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func MiTabla() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("numero para la tabla: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)

	}

	return texto

}
