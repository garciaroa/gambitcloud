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
	sentencia += o.Order_UserUUID + "'," + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "," + strconv.Itoa(o.Order_AddId) + ")"
	fmt.Println("bd/order 22 > sentencia " + sentencia)

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

func SelectOrders(user string, fechaDesde string, fechaHasta string, page int, orderId int) ([]models.Orders, error) {
	fmt.Println("Inicia SelectOrders")
	var Orders []models.Orders

	sentencia := "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total From  orders "

	if orderId > 0 {
		sentencia += " WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0
		if page == 0 {
			page = 1
		}
		if page > 1 {
			offset = (10 * (page - 1))
		}
		if len(fechaHasta) == 10 {
			fechaHasta += " 23:59:59"
		}

		var where string
		var whereUser string = " Order_UserUUID = ' " + user + "'"
		if len(fechaDesde) > 0 && len(fechaHasta) > 0 {
			where += " WHERE Order_Date BETWEEN '" + fechaDesde + " 'AND' " + fechaHasta
		}
		if len(where) > 0 {
			where += " WHERE " + whereUser
		}

		limit := " LIMIT 10"
		if offset > 0 {
			limit += " OFFSET " + strconv.Itoa(offset)
		}

		sentencia += where + limit
	}
	fmt.Println("bd/order 89 > sentencia " + sentencia)

	err := DbConnect()
	if err != nil {
		return Orders, err
	}
	defer func() {

		if err := Db.Close(); err != nil {
			fmt.Println("bd/order 98 > Error al intentar cerrar la conexion")
		}
	}()

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	if err != nil {
		return Orders, err
	}
	defer func() {

		if err := Db.Close(); err != nil {
			fmt.Println("bd/order 110 > Error al intentar ejecutar la sentencia")
		}
	}()

	for rows.Next() {
		var Order models.Orders
		var OrderDate sql.NullTime
		var OrderAddId sql.NullInt32
		err := rows.Scan(&Order.Order_Id, &Order.Order_UserUUID, &OrderAddId, &OrderDate, &Order.Order_Total)
		if err != nil {
			return Orders, err
		}

		Order.Order_Date = OrderDate.Time.String()
		Order.Order_AddId = int(OrderAddId.Int32)

		var rowsD *sql.Rows
		sentenciaD := "SELECT OD_Id, OD_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderID = " + strconv.Itoa(Order.Order_Id)
		rowsD, err = Db.Query(sentenciaD)
		if err != nil {
			return Orders, err
		}

		for rowsD.Next() {
			var OD_Id int64
			var OD_ProdId int64
			var OD_Quantity int64
			var OD_Price float64

			err = rowsD.Scan(&OD_Id, &OD_ProdId, &OD_Quantity, &OD_Price)
			if err != nil {
				return Orders, err
			}

			var od models.OrdersDetails
			od.OD_Id = int(OD_Id)
			od.OD_ProId = int(OD_ProdId)
			od.OD_Quantity = int(OD_Quantity)
			od.OD_Price = OD_Price
			Order.OrderDetail = append(Order.OrderDetail, od)
		}
		Orders = append(Orders, Order)
		rowsD.Close()
	}

	fmt.Println("bd/Orders 155 > Ejecucion exitosa")
	return Orders, nil
}
