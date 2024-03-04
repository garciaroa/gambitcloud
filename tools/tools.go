package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}

func ArmoSentencia(s string, fieldName string, typeField string, ValueN int, ValueF float64, ValueS string) string {
	if (typeField == "S" && len(ValueS) == 0) ||
		(typeField == "F" && ValueF == 0) ||
		(typeField == "N" && ValueN == 0) {
		return s
	}

	if !strings.HasSuffix(s, "SET ") {
		s += ", "
	}

	switch typeField {
	case "S": //si es string
		s += fieldName + " = '" + EscapeString(ValueS) + "'"
	case "N": //si es Numerico
		s += fieldName + " = " + strconv.Itoa(ValueN)
	case "F": //si es Flotante
		s += fieldName + " = " + strconv.FormatFloat(ValueF, 'e', -1, 64)
	}

	return s

}
