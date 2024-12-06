package settings

import (
	dom "macc/internal/domains"
	"os"
)

func GetSettings() *dom.SrvSettings {
	var obj = &dom.SrvSettings{}

	getFlags(obj)
	getEnv(obj)

	return obj
}

func getEnv(obj *dom.SrvSettings) {
	addres, isExist := os.LookupEnv("ADDRES")
	if isExist {
		obj.Address = addres
	}
}
