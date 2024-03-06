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
