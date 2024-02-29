package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/garciaroa/gambitcloud/models"
	"github.com/garciaroa/gambitcloud/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza registro de insertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES('" + c.CategName + "','" + "')"
	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	return LastInsertId, nil
}

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza registro de updateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "Update category SET "
	if len(c.CategName) > 0 {
		sentencia += " Categ_name = '" + tools.EscapeString(c.CategName) + "'"
	}
	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia += ", "
		}
		sentencia += "Categ_Path= '" + tools.EscapeString(c.CategPath) + "'"
	}
	sentencia += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecucion Exitosa")
	return nil
}
