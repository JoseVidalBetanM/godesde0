package defer_panic

import (
	"fmt"
	"log"
)

func Defen() {
	fmt.Println("primero")
	defer fmt.Println("ultimo")

	fmt.Println("segundo")

}

func Paniquito() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("se genero un PANIC \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Lptm un fallo se encontro un 1")
	}
}
