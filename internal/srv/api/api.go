package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"macc/internal/adapter"
	"macc/internal/srv/mng"
)

var ad mng.Adapter = adapter.NewMemoryAdapter()
var manager mng.Manager = mng.NewManager(ad)
var tpl = template.Must(template.ParseFiles("../../static/list.html"))

func Update(w http.ResponseWriter, r *http.Request) {
	/* if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
	} */

	mtype := chi.URLParam(r, "mtype")
	mname := chi.URLParam(r, "mname")
	mStringValue := chi.URLParam(r, "mvalue")
	var err error
	var description *mng.Description
	var mvalue interface{} = nil

	description, err = mng.GetDescription(mtype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch description.ValueType {
	case mng.ValueInt64:
		value, err := strconv.ParseInt(mStringValue, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mvalue = interface{}(value)
	case mng.ValueFloat64:
		value, err := strconv.ParseFloat(mStringValue, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mvalue = interface{}(value)
	default:
	}

	metric, err := mng.NewMetric(mtype, mvalue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := manager.Update(mname, *metric)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf("Metrics value = %+v", res)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusBadRequest)
	}

	list, err := manager.GetList()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, list)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
