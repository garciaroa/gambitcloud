package bd

import (
	"database/sql"
	"fmt"

	"github.com/garciaroa/gambitcloud/models"
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
