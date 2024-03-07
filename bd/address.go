package bd

import (
	"fmt"

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
