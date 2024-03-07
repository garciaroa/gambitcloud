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
