package adapter

import (
	"errors"
	dom "macc/internal/domains"
	"sync"
)

type MemoryAdapter struct {
	mutex      *sync.Mutex
	forInt64   map[string]map[string]int64
	forFloat64 map[string]map[string]float64
}

func NewMemoryAdapter() MemoryAdapter {
	ad := MemoryAdapter{
		mutex:      &sync.Mutex{},
		forInt64:   map[string]map[string]int64{},
		forFloat64: map[string]map[string]float64{},
	}
	return ad
}

func (adapter MemoryAdapter) GetList() ([]dom.Metric, error) {
	adapter.mutex.Lock()
	defer adapter.mutex.Unlock()

	list := make([]dom.Metric, 0)

	for name, data := range adapter.forInt64 {
		for mtype, mvalue := range data {
			list = append(list, dom.Metric{MName: name, MType: mtype, MValue: mvalue})
		}
	}

	for name, data := range adapter.forFloat64 {
		for mtype, mvalue := range data {
			list = append(list, dom.Metric{MName: name, MType: mtype, MValue: mvalue})
		}
	}

	return list, nil
}

func (adapter MemoryAdapter) GetInt64(mname string, mtype string) (interface{}, error) {
	adapter.mutex.Lock()
	defer adapter.mutex.Unlock()

	var volume int64 = 0
	typsForName, isExist := adapter.forInt64[mname]
	if !isExist {
		return interface{}(volume), nil
	}
	volume, isExist = typsForName[mtype]
	if !isExist {
		return interface{}(volume), nil
	}
	return interface{}(volume), nil
}

func (adapter MemoryAdapter) GetFloat64(mname string, mtype string) (interface{}, error) {
	adapter.mutex.Lock()
	defer adapter.mutex.Unlock()

	var volume float64 = 0
	typsForName, isExist := adapter.forFloat64[mname]
	if !isExist {
		return interface{}(volume), nil
	}
	volume, isExist = typsForName[mtype]
	if !isExist {
		return interface{}(volume), nil
	}
	return interface{}(volume), nil
}

func (adapter MemoryAdapter) SetInt64(mname string, mtype string, mvolume interface{}) error {
	adapter.mutex.Lock()
	defer adapter.mutex.Unlock()

	volume, success := mvolume.(int64)
	if !success {
		return errors.New("invalid type mvolume for SetInt64")
	}
	typsForName, isExist := adapter.forInt64[mname]
	if !isExist {
		typsForName = map[string]int64{}
		adapter.forInt64[mname] = typsForName
	}
	typsForName[mtype] = volume
	return nil
}

func (adapter MemoryAdapter) SetFloat64(mname string, mtype string, mvolume interface{}) error {
	adapter.mutex.Lock()
	defer adapter.mutex.Unlock()

	volume, success := mvolume.(float64)
	if !success {
		return errors.New("invalid type mvolume for SetFloat64")
	}
	typsForName, isExist := adapter.forFloat64[mname]
	if !isExist {
		typsForName = map[string]float64{}
		adapter.forFloat64[mname] = typsForName
	}
	typsForName[mtype] = volume
	return nil
}
