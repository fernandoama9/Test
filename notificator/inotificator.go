package notificator

type INotificator interface {
	Subscribe(notificator errorNotificator)
	Unsubscribe(notificator errorNotificator)
	Notify(ID string, pass string, fecha string, hora string)
}

type errorNotificator interface {
	Notify(ID string, pass string, fecha string, hora string)
	GetID() string //para hacer uso
}
