package api

import (
	"net/http"
	"strconv"
	"strings"

	"macc/internal/adapter"
	"macc/internal/srv/mng"
)

var ad mng.Adapter = adapter.NewMemoryAdapter()
var manager mng.Manager = mng.NewManager(ad)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/update/", update)

	return mux
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
	}

	path := r.URL.Path
	path = strings.Replace(path, "/update/", "", 1)
	parts := strings.Split(path, "/")

	var err error
	var mname string
	var mtype string
	var description *mng.Description
	var mvalue interface{} = nil

	for i, s := range parts {
		switch i {
		case 0:
			mtype = s
			description, err = mng.GetDescription(mtype)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case 1:
			mname = s
		case 2:
			switch description.ValueType {
			case mng.ValueInt64:
				value, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				mvalue = interface{}(value)
			case mng.ValueFloat64:
				value, err := strconv.ParseFloat(s, 64)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				mvalue = interface{}(value)
			default:
			}
		}
	}

	if mvalue == nil {
		http.Error(w, "Expected URL pattern /update/MetricsType/MerticsName/MetricsValue", http.StatusNotFound)
		return
	}

	metric, err := mng.NewMetric(mtype, mvalue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = manager.Update(mname, *metric)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
