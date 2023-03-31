package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 0, 88}
var matriz [10][10]int

func MuestroArreglos() {
	tabla[7] = 40
	tabla[2] = 73

	tabla2 := [10]string{"jose", "mateo", "luis", "lea"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[4][4] = 66
	fmt.Println(matriz)
}
