package routers

import (
	"encoding/json"

	"github.com/garciaroa/gambitcloud/bd"
	"github.com/garciaroa/gambitcloud/models"
)

func InsertAddress(body string, User string) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.AddAddress == "" {
		return 400, "Debe Especificar el Address"
	}

	if t.AddName == "" {
		return 400, "Debe Especificar el Name"
	}

	if t.AddTitle == "" {
		return 400, "Debe Especificar el Title"
	}

	if t.AddCity == "" {
		return 400, "Debe Especificar el City"
	}

	if t.AddPhone == "" {
		return 400, "Debe Especificar el Phone"
	}

	if t.AddPostalCode == "" {
		return 400, "Debe Especificar el PostalCode"
	}

	err = bd.InsertAddress(t, User)
	if err != nil {
		return 400, "routers/address 43 > Ocurrio un error al intentar realizar el registro del address para el ID de usuario " + User + ">" + err.Error()
	}

	return 200, "InsertAddres OK"
}

func Updateaddress(body string, user string, id int) (int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}
	t.AddId = id
	var encontrado bool
	err, encontrado = bd.AddressExists(user, id)
	if !encontrado {
		if err != nil {
			return 400, "Error al intentar buscar address para el usuario" + user + ">" + err.Error()
		}
		return 400, "Nose encuentra un registro de ID de Usuario asociado a esa ID  de address"
	}

	err = bd.UpdateAddress(t)
	if err != nil {
		return 400, "Ocurrio un errror al intentar realizar la actualizacion del address para el ID de usuario " + user + " > " + err.Error()
	}

	return 200, "Updateaddress Ok"
}

func DeleteAddress(User string, id int) (int, string) {
	err, encontrado := bd.AddressExists(User, id)
	if !encontrado {
		if err != nil {
			return 400, " Erroral intentar buscar Address para el usuario " + User + " > " + err.Error()
		}
		return 400, "No se encuentra un registro de ID de Usuario asociado a esa Id de adddress"
	}

	err = bd.DeleteAddress(id)
	if err != nil {
		return 400, "Ocurrio un error al intentar borrar una direccion de usuario '" + User + "' > " + err.Error()
	}

	return 200, "DeleteAddress OK"

}

func SelectAddress(User string) (int, string) {
	addr, err := bd.SelectAddress(User)
	if err != nil {
		return 400, "Ocurrio un error al intentar obtener la lista de direcciones del usuario " + User + " > " + err.Error()
	}

	respJson, err := json.Marshal(addr)
	if err != nil {
		return 500, "Error al formatear los datos de las addresses como JSON"
	}
	return 200, string(respJson)

}
