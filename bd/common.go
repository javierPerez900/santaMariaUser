package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/javier/santaMariaUser/models"
	"github.com/javier/santaMariaUser/secretm"
)

var SecretModel models.SecretRDSJson
var Db *sql.DB

func ReadSecret() error{
	secret, err := secretm.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		return err
	}

	SecretModel = secret
	return nil
}

func DbConnect() error {
	db, err := sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return err
	}

	Db = db
	fmt.Println("Conexión exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndPoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndPoint = claves.Host
	dbName = "santa_maria"

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndPoint, dbName)
}