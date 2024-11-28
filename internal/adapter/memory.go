package adapter

import "errors"

type MemoryAdapter struct {
	forInt64   map[string]map[string]int64
	forFloat64 map[string]map[string]float64
}

func NewMemoryAdapter() MemoryAdapter {
	ad := MemoryAdapter{map[string]map[string]int64{}, map[string]map[string]float64{}}
	return ad
}

func (adapter MemoryAdapter) GetInt64(mname string, mtype string) (interface{}, error) {
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
