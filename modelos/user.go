package modelos

import "time"

type User struct {
	Id        int
	Name      string
	CreateAtt time.Time
	Status    bool
}

func (usuario *User) AddUser(id int, name string, createatt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreateAtt = createatt
	usuario.Status = status

}
