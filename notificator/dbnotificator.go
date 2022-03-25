package notificator

import "fmt"

type dbnotificator struct {
	aux string
	id string
	//iria la dependencia de la BD
}

func NewDbnotificator(id string) *dbnotificator {
	return &dbnotificator{id: id}
}

func (d *dbnotificator) Notify(ID string, pass string, fecha string, hora string) {
	fmt.Println("almacenando en la BD ID: " + ID +", pass: "+ pass +", fecha: "+ fecha+ ", hora: " + hora)
	d.aux = ID+pass+fecha+hora
	fmt.Println("almacenado en la BD")
}

func (d *dbnotificator) GetID() string {
	return d.id
}
