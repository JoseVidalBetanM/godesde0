package users

import (
	"fmt"
	"time"

	"github.com/JoseVidalBetanM/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(13, "Pablo", time.Now(), true)
	fmt.Println(u)
}
