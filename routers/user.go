package routers

import (
	"encoding/json"

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
