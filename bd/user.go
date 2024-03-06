package bd

import (
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
