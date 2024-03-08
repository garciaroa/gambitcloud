package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/garciaroa/gambitcloud/models"
)

func InsertAddress(addr models.Address, User string) error {
	fmt.Println("Comienza el Registro InsertAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer func() {

		if err := Db.Close(); err != nil {
			fmt.Println("bd/address 19 > error al conectar a la Bd")
		}
	}()

	sentencia := "INSERT INTO addresses (Add_UserID, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name)"
	sentencia += "VALUES ('" + User + "','" + addr.AddAddress + "','" + addr.AddCity + "','" + addr.AddState + "','"
	sentencia += addr.AddPostalCode + "','" + addr.AddPhone + "','" + addr.AddTitle + "','" + addr.AddName + " ')"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("bd/address 33 > sentencia" + sentencia)
	fmt.Println("Insert Address 34 > Ejecucion Exitosa")
	return nil
}

func AddressExists(User string, id int) (error, bool) {
	fmt.Println("Comienza AddresExists")

	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()
	sentencia := "SELECT 1 FROM addresses WHERE Add_Id=" + strconv.Itoa(id) + " AND Add_UserId= '" + User + "'"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)
	if err != nil {
		return err, false
	}
	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("Address > ejecucion exitosa " + valor)
	if valor == "1" {
		return nil, true
	}
	return nil, false
}

func UpdateAddress(addr models.Address) error {
	fmt.Println("comienza UpdateAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer func() {
		if err := Db.Close(); err != nil {
			fmt.Println(" bd/address 74 > error al intentar cerrar la conexion DB")
		}
	}()

	sentencia := "UPDATE addresses SET "
	if addr.AddAddress != "" {
		sentencia += "Add_Address = '" + addr.AddAddress + "', "
	}
	if addr.AddCity != "" {
		sentencia += "Add_City = '" + addr.AddCity + "', "
	}
	if addr.AddName != "" {
		sentencia += "Add_Name = '" + addr.AddName + "', "
	}
	if addr.AddPhone != "" {
		sentencia += "Add_Phone = '" + addr.AddPhone + "', "
	}
	if addr.AddPostalCode != "" {
		sentencia += "Add_PostalCode = '" + addr.AddPostalCode + "', "
	}
	if addr.AddState != "" {
		sentencia += "Add_State = '" + addr.AddState + "', "
	}
	if addr.AddTitle != "" {
		sentencia += "Add_Title = '" + addr.AddTitle + "', "
	}
	sentencia, _ = strings.CutSuffix(sentencia, ", ") //corta el ultio caracter ", "
	sentencia += " WHERE Add_Id= " + strconv.Itoa(addr.AddId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("bd/address 110 > sentencia " + sentencia)
	fmt.Println("Update Addres > Ejecucion Exitosa")
	return nil

}

func DeleteAddress(id int) error {
	fmt.Println("inicia Delete Address")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer func() {
		if err := Db.Close(); err != nil {
			fmt.Println("bd/address 126 > Error al cerrar la conexion")
		}
	}()

	sentencia := "DELETE FROM addresses WHERE Add_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println("bd/address 133 > error al ejecutar la sentencia " + err.Error())
		return err
	}

	fmt.Println("bd/address 138 > sentencia " + sentencia)
	fmt.Println("Delete Address > Ejecucion exitosa")
	return nil

}

func SelectAddress(User string) ([]models.Address, error) {
	fmt.Println("Comienza selectaddress")

	addr := []models.Address{}
	err := DbConnect()
	if err != nil {
		return addr, err
	}
	defer func() {
		if err := Db.Close(); err != nil {
			fmt.Println("bd/address 154 > error al intentar cerrar la conexion")
			return
		}
	}()

	sentencia := "SELECT Add_Id, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name, FROM addresses WHERE Add_UserID = '" + User + "'"

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	if err != nil {
		fmt.Println("bd/address 165 > error al ejecutar la sentencia " + err.Error())
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Println("bd/address 169 > error al cerrar la conexion a bd")
			return
		}
	}()

	for rows.Next() {
		var a models.Address
		var addId sql.NullInt16
		var addAddress sql.NullString
		var addCity sql.NullString
		var addState sql.NullString
		var addPostalCode sql.NullString
		var addPhone sql.NullString
		var addTitle sql.NullString
		var addName sql.NullString

		err := rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)
		if err != nil {
			return addr, err
		}

		a.AddId = int(addId.Int16)
		a.AddAddress = addAddress.String
		a.AddCity = addCity.String
		a.AddState = addState.String
		a.AddPostalCode = addPostalCode.String
		a.AddPhone = addPhone.String
		a.AddTitle = addTitle.String
		a.AddName = addName.String
		addr = append(addr, a)
	}

	fmt.Println("Select Address > Ejecucion Exitosa ")
	return addr, nil

}
