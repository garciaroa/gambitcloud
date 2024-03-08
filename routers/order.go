package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/garciaroa/gambitcloud/bd"
	"github.com/garciaroa/gambitcloud/models"
)

func InsertOrder(body string, User string) (int, string) {
	var o models.Orders
	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		return 400, "Error en los datos recibidos"
	}

	o.Order_UserUUID = User

	ok, message := ValidOrder(o)
	if !ok {
		return 400, message
	}
	result, err2 := bd.InsertOrder(o)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de la orden " + err2.Error()
	}

	return 200, "{OrderID:" + strconv.Itoa(int(result)) + "}"

}

func ValidOrder(o models.Orders) (bool, string) {
	if o.Order_Total == 0 {
		return false, "Debe indicar el total de la orden"
	}
	count := 0
	fmt.Println("routers/order 39 > antes del range")
	for _, od := range o.OrderDetail {
		fmt.Println("routers/order 41 > ingresa al range")
		if od.OD_ProId == 0 {
			return false, "Debe Indicar el Id del producto en el detalle de la orden"
		}
		if od.OD_Quantity == 0 {
			return false, "Debe Indicar la cantidad del producto en el detalle de la orden"
		}

		count++

		fmt.Println("routers/order 51 > pasa el count" + strconv.Itoa(count))

	}

	fmt.Println("routers/order 55 > fuera del range" + strconv.Itoa(count))

	if count == 0 {
		return false, "Debe indicar items en la orden"
	}
	return true, ""

}
