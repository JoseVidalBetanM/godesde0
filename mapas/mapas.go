package mapas

import "fmt"

func Mapitas() {
	paises := make(map[string]string)
	// fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	paises["Venezuela"] = "Caracas"

	fmt.Println(paises)
	fmt.Println(paises["Venezuela"])

}

func Campeonato() {
	partidos := map[string]int{
		"Real Madrid": 35,
		"Barcelona":   44,
		"Vinotinto":   100,
		"Chivas":      20,
	}

	fmt.Println(partidos)
	/*for equipo, puntaje := range partidos {
		fmt.Printf("el equipo %s tiene un puntaje de: %d \n", equipo, puntaje)
	}
	*/
	delete(partidos, "Chivas")
	fmt.Println(partidos)

	puntaje, existe := partidos["Vinotinto"]
	fmt.Printf("el punteje capturado es: %d y el equipo exixte %t \n", puntaje, existe)

}
