package mng

import (
	"fmt"
)

var metricsTypes = map[string]*Description{
	"gauge":   {ValueFloat64, OperationReplace},
	"counter": {ValueInt64, OperationAdd},
}

type Description struct {
	ValueType     string
	OperationType string
}

const (
	ValueInt64   = "int64"
	ValueFloat64 = "float64"
)

const (
	OperationReplace = "replace"
	OperationAdd     = "add"
)

type Metric struct {
	metricType    string
	operationType string
	valueType     string
	value         interface{}
}

func (metric Metric) GetMetricType() string {
	return metric.metricType
}

func (metric Metric) GetOperationType() string {
	return metric.operationType
}

func (metric Metric) GetValueType() string {
	return metric.valueType
}

func (metric Metric) GetValue() interface{} {
	return metric.value
}

func NewMetric(mtype string, value interface{}) (*Metric, error) {
	description, err := GetDescription(mtype)
	if err != nil {
		return nil, err
	}

	_, err = GetDescription(mtype)
	if err != nil {
		return nil, err
	}

	metric := Metric{mtype, description.OperationType, description.ValueType, value}
	return &metric, nil
}

func GetDescription(mtype string) (*Description, error) {
	description, isExist := metricsTypes[mtype]
	if !isExist {
		return nil, fmt.Errorf("invalid metrics type %s", mtype)
	}
	return description, nil
}

func GetValueType(value interface{}) (string, error) {
	var valueType string
	switch t := value.(type) {
	case int64:
		valueType = ValueInt64
	case float64:
		valueType = ValueFloat64
	default:
		return "", fmt.Errorf("invalid value type %t", t)
	}
	return valueType, nil
}
