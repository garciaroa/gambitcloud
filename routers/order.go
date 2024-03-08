package routers

import (
	"encoding/json"
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
		return 400, "Ocurrio un error al intentar realizar el registro de la orden " + err.Error()
	}

	return 200, "{OrderID:" + strconv.Itoa(int(result)) + "}"

}

func ValidOrder(o models.Orders) (bool, string) {
	if o.Order_Total == 0 {
		return false, "Debe indicar el total de la orden"
	}
	count := 0
	for _, od := range o.OrderDetail {
		if od.OD_ProId == 0 {
			return false, "Debe Indicar el Id del producto en el detalle de la orden"
		}
		if od.OD_Quantity == 0 {
			return false, "Debe Indicar la cantidad del producto en el detalle de la orden"
		}
		count++
	}
	if count == 0 {
		return false, "Debe indicar items en la orden"
	}
	return true, ""

}
