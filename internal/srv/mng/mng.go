package mng

import "errors"

type Adapter interface {
	GetInt64(mname string, mtype string) (interface{}, error)
	GetFloat64(mname string, mtype string) (interface{}, error)
	SetInt64(mname string, mtype string, mvolume interface{}) error
	SetFloat64(mname string, mtype string, mvolume interface{}) error
}

type Manager struct {
	adapter Adapter
}

func NewManager(adapter Adapter) Manager {
	manager := Manager{adapter}
	return manager
}

func (manager Manager) Update(name string, metric Metric) error {
	var funcGet func(mname string, mtype string) (interface{}, error)
	var funcSet func(mname string, mtype string, mvolume interface{}) error

	switch metric.GetValueType() {
	case ValueInt64:
		funcGet = manager.adapter.GetInt64
		funcSet = manager.adapter.SetInt64
	case ValueFloat64:
		funcGet = manager.adapter.GetFloat64
		funcSet = manager.adapter.SetFloat64
	default:
		return errors.New("invalid metric value type")
	}

	mt := metric.GetMetricType()
	ot := metric.GetOperationType()
	value := metric.GetValue()
	switch ot {
	case OperationReplace:
		err := funcSet(name, mt, value)
		if err != nil {
			return err
		}
	case OperationAdd:
		currentVolume, err := funcGet(name, mt)
		if err != nil {
			return err
		}
		res, err := sum(currentVolume, value)
		if err != nil {
			return err
		}
		err = funcSet(name, mt, res)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid metric operation type")
	}

	return nil
}

func sum(v1 interface{}, v2 interface{}) (interface{}, error) {
	switch v1.(type) {
	case int64:
		res := v1.(int64) + v2.(int64)
		return interface{}(res), nil
	case float64:
		res := v1.(float64) + v2.(float64)
		return interface{}(res), nil
	}

	return nil, errors.New("sum operation is not supported")
}
