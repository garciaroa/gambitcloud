package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/garciaroa/gambitcloud/models"
)

// /InsertOrder inserta la orden
func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Comienza Registro Oders")
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) VALUES('"
	sentencia += o.Order_UserUUID + "'" + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "," + strconv.Itoa(o.Order_AddId) + ")"

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

	for _, od := range o.OrderDetail {
		sentencia = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId))
		sentencia += "," + strconv.Itoa(od.OD_ProId) + "," + strconv.Itoa(od.OD_Quantity) + "," + strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + ")"

		fmt.Println("bd/order 38 > sentencia " + sentencia)

		_, err = Db.Exec(sentencia)
		if err != nil {
			fmt.Println("error al ejecutar la senencia " + sentencia)
			return 0, err
		}
	}

	fmt.Println("Insert Order > Ejecucion Exitosa")
	return LastInsertId, nil
}
