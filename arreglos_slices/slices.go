package arreglos_slices

import "fmt"

var tablaS []int = []int{6, 76, 3}
var arreglos [10]int = [10]int{32, 54, 66, 26, 89, 2, 54, 334}

func MuestroSlices() {
	fmt.Println(tablaS)

	porcion := arreglos[3:]
	porcion2 := arreglos[:5]
	porcion3 := arreglos[4:7]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
