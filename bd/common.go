package bd

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/garciaroa/gambitcloud/models"
	"github.com/garciaroa/gambitcloud/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB
var DsnGlobal string

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambitcloud"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	DsnGlobal = dsn
	fmt.Println("ConnStr - token - dbUser" + dsn + authToken + dbUser)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza UserIsAdmin")
	err := DbConnect()
	if err != nil {
		return false, SecretModel.Username + "-" + SecretModel.Password + "-" + SecretModel.Host + "-" + SecretModel.Engine + "-" + SecretModel.DbClusterIdentifier + "-" + strconv.Itoa(SecretModel.Port)
	}

	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0 "
	fmt.Println(sentencia)
	rows, err := Db.Query(sentencia)
	if err != nil {
		return false, err.Error()
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("UserIsAdmin > Ejecucion exitosa - valor devuelto" + valor)
	if valor == "1" {
		return true, ""
	}

	return false, "user is not Admin"
}
