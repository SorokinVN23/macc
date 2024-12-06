package settings

import (
	dom "macc/internal/domains"
	"os"
	"strconv"
)

func GetSettings() *dom.AgentSettings {
	var obj = &dom.AgentSettings{}

	getFlags(obj)
	getEnv(obj)

	return obj
}

func getEnv(obj *dom.AgentSettings) {
	addres, isExist := os.LookupEnv("ADDRES")
	if isExist {
		obj.Address = addres
	}

	r_interval, isExist := os.LookupEnv("REPORT_INTERVAL")
	if isExist {
		i, err := strconv.ParseUint(r_interval, 10, 64)
		if err != nil {
			panic(err)
		}
		obj.ReportInterval = int(i)
	}

	p_interval, isExist := os.LookupEnv("POLL_INTERVAL")
	if isExist {
		i, err := strconv.ParseUint(p_interval, 10, 64)
		if err != nil {
			panic(err)
		}
		obj.PollInterval = int(i)
	}
}
