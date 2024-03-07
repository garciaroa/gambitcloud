package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/garciaroa/gambitcloud/bd"
	"github.com/garciaroa/gambitcloud/models"
)

func UpdateUser(body string, User string) (int, string) {
	var t models.User
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.UserFirstName) == 0 && len(t.UserFirstName) == 0 {
		return 400, "Debe especificar el nombre (FirstName) o (lastName)"
	}
	_, encontrado := bd.UserExists(User)
	if !encontrado {
		return 400, "No existe un usuario con ese UUID '" + User + "'"
	}

	err = bd.UpdateUser(t, User)
	if err != nil {
		return 400, "ocurrio un error alintentar realizar la actualizacion del usuario " + User + " > " + err.Error()
	}

	return 200, "UpdateUser OK"

}

func SelectUser(body string, User string) (int, string) {
	_, encontrado := bd.UserExists(User)

	if !encontrado {
		return 400, "No existe un Usuario con ese UUID '" + User + "'"
	}

	row, err := bd.SelectUser(User)

	if err != nil {
		return 400, "Ocurrio un error al intentar realizar el selected del usuario " + User + " > " + err.Error()
	}

	respJson, err := json.Marshal(row)
	if err != nil {
		return 500, "Error al formatear los datos del usuario como JSON"
	}

	return 200, string(respJson)

}

func SelectUsers(body string, User string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var Page int
	if len(request.QueryStringParameters["Page"]) == 0 {
		Page = 1
	} else {
		Page, _ = strconv.Atoi(request.QueryStringParameters["Page"])
	}
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	user, err := bd.SelectUsers(Page)
	if err != nil {
		return 400, "Ocurrio un error al intentar obtener la lista de usuarios 72 > " + err.Error()
	}

	respJson, err := json.Marshal(user)
	if err != nil {
		return 500, "Error al formatear los datos de los usuarios como JSON"
	}

	return 200, string(respJson)

}
