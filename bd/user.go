package bd

import (
	"database/sql"
	"fmt"

	"github.com/garciaroa/gambitcloud/models"
	"github.com/garciaroa/gambitcloud/tools"
)

func UpdateUser(UField models.User, User string) error {
	fmt.Println("Comienza Update")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()
	sentencia := "UPDATE users SET "
	coma := ""
	if len(UField.UserFirstName) > 0 {
		coma = ","
		sentencia += "User_FirstName = '" + UField.UserFirstName + "'"
	}
	if len(UField.UserLastName) > 0 {
		sentencia += coma + "User_LastName = '" + UField.UserFirstName + "'"
	}
	sentencia += ", User_DateUpg = '" + tools.FechaMySQL() + "' WHERE User_UUID='" + User + "'"
	fmt.Println("usuario 27> sentencia :.. " + sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("actualizacion usuario 35> Ejecucion exitosa")
	return nil
}

func SelectUser(UserId string) (models.User, error) {
	fmt.Println("Comienza SelecUser")
	User := models.User{}

	err := DbConnect()
	if err != nil {
		return User, err
	}
	defer Db.Close()

	sentencia := "SELECT * FROM users WHERE User_UUID = '" + UserId + "'"

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	if err != nil {
		fmt.Println("bd/usuario 54" + err.Error())
		return User, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Println("bd/usuario 61" + err.Error())
		}
	}()

	rows.Next()
	/*var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullString //sql.NullTime*/

	fmt.Println("bd/user 68 > antes del panic ")

	rows.Scan(&User.UserUUID, &User.UserEmail, &User.UserFirstName, &User.UserLastName, &User.UserStatus, &User.UserDateAdd, &User.UserDateUpd)

	fmt.Println("bd/user 70 > rows.scan ")
	/*
		User.UserFirstName = firstName.String
		User.UserLastName = lastName.String
		User.UserDateUpd = dateUpg.String //.Time.String()*/

	fmt.Println("Select User > Ejecucion Exitosa")
	return User, nil

}
