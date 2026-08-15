package bd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/javier/santaMariaUser/models"
	"github.com/javier/santaMariaUser/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error{
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println( err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println( err.Error())
		return err
	}

	fmt.Println("Conexión exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndPoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndPoint = claves.Host
	dbName = "santa_maria"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndPoint, dbName)
	fmt.Println(dsn)
	return dsn
}