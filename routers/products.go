package routers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/garciaroa/gambitcloud/bd"
	"github.com/garciaroa/gambitcloud/models"
)

func InsertProducto(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el nombre (Title) del producto"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro del producto " + t.ProdTitle + ">"
	}
	return 200, "{ProductID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateProduct(body string, User string, id int) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el update del producto" + strconv.Itoa(id) + ">" + err2.Error()
	}

	return 200, "Update ok"
}

func DeleteProduct(User string, id int) (int, string) {
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteProduct(id)
	if err != nil {
		return 400, "Ocurrio un error al intentar realizar el DELETE  del producto" + strconv.Itoa(id) + ">" + err.Error()
	}

	return 200, "Delete ok"

}
func SelectProduct(request events.APIGatewayV2HTTPRequest) (int, string) {
	var t models.Product
	//var err error
	var page, pageSize int
	var orderType, orderField string

	param := request.QueryStringParameters
	page, _ = strconv.Atoi(param["page"])
	pageSize, _ = strconv.Atoi(param["pageSize"])
	orderType = param["orderType"]   //D= desc - (A o Nil) = ASC
	orderField = param["orderField"] // ('I' por Id), ('T' Title), ('D' Description), ('F' Created At) ('P' Price), ('C' CategId), ('S' Stock)

	if !strings.Contains("ITDFPCS", orderField) {
		orderField = ""
	}

	var choice string

	if len(param["prodId"]) > 0 {
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}
	if len(param["search"]) > 0 {
		choice = "S"
		t.ProdSearch = param["search"]
	}
	if len(param["categId"]) > 0 {
		choice = "C"
		t.ProdCategoryId, _ = strconv.Atoi(param["categId"])
	}
	if len(param["slug"]) > 0 {
		choice = "U"
		t.ProdPath, _ = param["slug"]
	}
	if len(param["slugCateg"]) > 0 {
		choice = "K"
		t.ProdCategPath = param["slugCateg"]
	}
	fmt.Println(param)
	result, err2 := bd.SelectProduct(t, choice, page, pageSize, orderType, orderField)
	if err2 != nil {
		return 400, "selectProduct > Ocurrio un error al intentar capturar los resultados de la busqueda de tipo '" + choice + "' en products"
	}

	product, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "selectProduct > Ocurrio un error al intentar convertir en JSON la busqueda de Productos"
	}

	return 200, string(product)
}

func UpdateStock(body string, User string, id int) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateStock(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el update del Stock" + strconv.Itoa(id) + ">" + err2.Error()
	}

	return 200, "Update Stock ok"
}
