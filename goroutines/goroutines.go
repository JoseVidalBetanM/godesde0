package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func Velocito(nombre string, canal1 chan bool) {
	palabra := strings.Split(nombre, "")
	for _, letra := range palabra {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	canal1 <- true
}
