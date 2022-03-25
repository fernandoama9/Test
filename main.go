package main

import (
	"fmt"
	"pruebaLDev180222/notificator"
	"pruebaLDev180222/validator"
)

func main() {
	var notificador notificator.INotificator
	txtNotificator := notificator.NewTextfilenotificator("TXT")
	dbNotificator := notificator.NewDbnotificator("DB")
	notificador = notificator.NewNotificator()
	fmt.Println("MAIN - Suscribiendo DB")
	notificador.Subscribe(dbNotificator)
	fmt.Println("MAIN - Suscribiendo Escritor de archivos TXT")
	notificador.Subscribe(txtNotificator)
	var fachada validator.CredentialValidator
	fachada = validator.NewValidatorFacade(notificador, validator.GetInstance())
	fmt.Println("------------MAIN - Iniciando pruebas-------------------")
	// fallo
	fachada.Validate("fallo", "fallo")
	// Exito
	fachada.Validate("prueba", "prueba")
	fmt.Println("MAIN - Unsuscribe Escritor de archivos de texto")
	notificador.Unsubscribe(txtNotificator)
	fmt.Println("------------MAIN - Iniciando pruebas de nuevo-------------------")
	// fallo
	fachada.Validate("fallo", "fallo")
	// Exito
	fachada.Validate("prueba", "prueba")
}
