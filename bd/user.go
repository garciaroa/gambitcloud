package bd

import (
	"database/sql"
	"fmt"
	"strconv"

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
	var firstName sql.NullString
	var lastName sql.NullString
	//var dateUpg sql.NullTime //sql.NullString //sql.NullTime

	fmt.Println("bd/user 68 > antes del panic ")

	rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &User.UserDateUpd) //&User.UserDateUpd) de bd TIME a Strng en modelo user y es llenado con vacio ""

	fmt.Println("bd/user 70 > rows.scan ")

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	//User.UserDateUpd = dateUpg.Time.String()

	fmt.Println("Select User > Ejecucion Exitosa")
	return User, nil

}

func SelectUsers(Page int) (models.ListUsers, error) {
	fmt.Println("Comienza SelectUsers")

	var lu models.ListUsers
	User := []models.User{}

	err := DbConnect()
	if err != nil {
		return lu, err
	}
	defer func() {
		if err := Db.Close(); err != nil {
			fmt.Println("bd/usuario 61 > " + err.Error())
		}
	}()

	var offset int = (Page * 10) - 10
	var sentencia string
	var sentenciaCount string = "SELECT count(*) as registros FROM users"

	sentencia = "select * from users LIMIT 10"
	if offset > 0 {
		sentencia += " OFFSET " + strconv.Itoa(offset)
	}

	var rowsCount *sql.Rows
	rowsCount, err = Db.Query(sentenciaCount)
	if err != nil {
		return lu, err
	}

	defer func() {
		if err := rowsCount.Close(); err != nil {
			fmt.Println("bd/usuario > " + err.Error())
		}
	}()

	rowsCount.Next()
	var registros int
	rowsCount.Scan(&registros)
	lu.TotalItems = registros

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return lu, err
	}

	for rows.Next() {
		var u models.User
		var firstName sql.NullString
		var lastName sql.NullString
		//var dateUpg sql.NullTime //sql.NullString //sql.NullTime

		fmt.Println("bd/user 140 > select usuarios ")

		rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &u.UserDateUpd) //&User.UserDateUpd) de bd TIME a Strng en modelo user y es llenado con vacio ""

		fmt.Println("bd/user 70 > rows.scan ")

		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		//User.UserDateUpd = dateUpg.Time.String()

		User = append(User, u)

	}

	fmt.Println("Select Users 154 > conexion exitosa")
	lu.Data = User

	return lu, nil

}
