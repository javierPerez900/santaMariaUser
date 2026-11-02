package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/javier/santaMariaUser/models"
	"github.com/javier/santaMariaUser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	// Idempotente: Cognito puede reintentar la invocación, y un reintento
	// con el mismo usuario no debe fallar (requiere que User_UUID sea único).
	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE User_UUID = User_UUID"

	_, err = Db.Exec(sentencia, sig.UserEmail, sig.UserUUID, tools.FechaMySQL())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución Exitosa")
	return nil
}