package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/garciaroa/gambitcloud/bd"
	"github.com/garciaroa/gambitcloud/models"
)

func InsertOrder(body string, User string) (int, string) {
	var o models.Orders
	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
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
	fmt.Println("routers/order 39 > antes del range" + strconv.Itoa(len(o.OrderDetail)))
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

func SelectOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string) {

	var fechaDesde, fechaHasta string
	var orderId, page int

	if len(request.QueryStringParameters["fechaDesde"]) > 0 {
		fechaDesde = request.QueryStringParameters["fechadesde"]
	}
	if len(request.QueryStringParameters["fechaHasta"]) > 0 {
		fechaDesde = request.QueryStringParameters["fechaHasta"]
	}
	if len(request.QueryStringParameters["page"]) > 0 {
		fechaDesde = request.QueryStringParameters["page"]
	}
	if len(request.QueryStringParameters["orderId"]) > 0 {
		fechaDesde = request.QueryStringParameters["orderId"]
	}
	result, err2 := bd.SelectOrders(user, fechaDesde, fechaHasta, page, orderId)
	if err2 != nil {
		return 400, "ocurrio un error al intentar capturar los registros de ordenes del " + fechaDesde + " al " + fechaHasta
	}

	fmt.Println("router/order > 86 ")
	orders, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON el registro de Orden"
	}

	return 200, string(orders)
}
