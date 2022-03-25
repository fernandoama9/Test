package notificator

import "fmt"

type textfilenotificator struct {
	id string
	aux string
	//dependencias de escribir en archivos
}

func NewTextfilenotificator(id string) *textfilenotificator {
	return &textfilenotificator{id: id}
}

func (t *textfilenotificator) Notify(ID string, pass string, fecha string, hora string) {
	fmt.Println("almacenando en archivo ID: " + ID +", pass: "+ pass +", fecha: "+ fecha+ ", hora: " + hora)
	t.aux = ID+pass+fecha+hora
	fmt.Println("almacenado en archivo")
}

func (t *textfilenotificator) GetID() string {
	return t.id
}
