package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/javier/santaMariaUser/awsgo"
	"github.com/javier/santaMariaUser/bd"
	"github.com/javier/santaMariaUser/models"
)

// Cognito invoca este trigger al confirmar el registro y también al recuperar
// la contraseña. Solo el registro debe crear el usuario en la base de datos.
const triggerSignUp = "PostConfirmation_ConfirmSignUp"

// Cognito le muestra al usuario el texto del error que devuelve la función,
// por eso nunca se devuelve el detalle interno.
var errRegistro = errors.New("no se pudo completar el registro, inténtalo de nuevo más tarde")

// Resultado del arranque, que se hace una vez por contenedor antes de recibir peticiones
var errInicio error

func main() {
	errInicio = inicializar()
	if errInicio != nil {
		// No se hace panic: el error se devuelve de forma controlada
		// al recibir una petición que realmente necesite la base de datos
		fmt.Println("Error al inicializar, se reintentará en la primera petición: " + errInicio.Error())
	}

	lambda.Start(EjecutoLambda)
}

func inicializar() error {
	if !ValidoParametros() {
		return errors.New("error en los parámetros: debe enviar 'SecretName'")
	}

	if err := awsgo.InicializoAWS(); err != nil {
		return err
	}

	return bd.ReadSecret()
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	if event.TriggerSource != triggerSignUp {
		return event, nil
	}

	if errInicio != nil {
		// El arranque falló (por ejemplo, un fallo transitorio de Secrets Manager): se reintenta
		errInicio = inicializar()
		if errInicio != nil {
			fmt.Println("Error al inicializar: " + errInicio.Error())
			return event, errRegistro
		}
	}

	datos := models.SignUp{
		UserEmail: event.Request.UserAttributes["email"],
		UserUUID:  event.Request.UserAttributes["sub"],
	}

	if datos.UserEmail == "" || datos.UserUUID == "" {
		fmt.Println("El evento no trae los atributos 'email' y 'sub'")
		return event, errRegistro
	}

	if err := bd.SignUp(datos); err != nil {
		fmt.Println("Error en SignUp: " + err.Error())
		return event, errRegistro
	}

	return event, nil
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	return traeParametro
}

// func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	
// 	awsgo.InicializoAWS()

// 	if !ValidoParametros() {
// 		fmt.Println("Error en los parámetros. Debe enviar 'SecretName'")
// 		err := errors.New("error en los parámetros. Debe enviar 'SecretName'")
// 		return event, err
// 	}

// 	var datos models.SignUp

// 	for row, att := range event.Request.UserAttributes {
// 		switch row {
// 		case "email":
// 			datos.UserEmail = att
// 			fmt.Println("Email = " + datos.UserEmail)
// 		case "sub":
// 			datos.UserUUID = att
// 			fmt.Println("Sub = " + datos.UserUUID)
// 		}
// 	}

// 	err := bd.ReadSecret()
// 	if err != nil {
// 		fmt.Println("Error al leer el secret " + err.Error())
// 		return event, err
// 	}

// 	err = bd.SignUp(datos)
// 	return event, err
// }

// func ValidoParametros() bool {
// 	var traeParametro bool
// 	_, traeParametro = os.LookupEnv("SecretName")
// 	return traeParametro
// }