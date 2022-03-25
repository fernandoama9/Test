package notificator

import "fmt"

type notificator struct {
	notificators []errorNotificator
	errAux       error
}

func NewNotificator() *notificator {
	return &notificator{}
}

func (n *notificator) Subscribe(notificator errorNotificator) {
	fmt.Println("Suscribiendo notificador con ID: " + notificator.GetID())
	n.notificators = append(n.notificators, notificator)
}

func (n *notificator) Unsubscribe(notificator errorNotificator) {
	fmt.Println("Cancelando suscripcion para notificador con ID: " + notificator.GetID())
	n.notificators = removeFromslice(n.notificators, notificator)
}

func (n *notificator) Notify(ID string, pass string, fecha string, hora string) {
	for _, noti := range n.notificators {
		noti.Notify(ID, pass, fecha, hora)
		//si hubiera error haria errAux = err
	}
}

func removeFromslice(notificatorList []errorNotificator, notificatorLToRemove errorNotificator) []errorNotificator {
	observerListLength := len(notificatorList)
	for i, observer := range notificatorList {
		if notificatorLToRemove.GetID() == observer.GetID() {
			notificatorList[observerListLength-1], notificatorList[i] = notificatorList[i], notificatorList[observerListLength-1]
			return notificatorList[:observerListLength-1]
		}
	}
	return notificatorList
}
