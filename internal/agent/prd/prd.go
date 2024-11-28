package prd

import (
	mng "macc/internal/agent/mng"
	"reflect"
	"runtime"
)

type Producer struct{}

func (p Producer) GetNewMetricsValue(necessaryMetrics map[string]mng.MetricType) map[string]mng.Metric {
	res := make(map[string]mng.Metric)
	ms := getMemStatus()

	for name, mtype := range necessaryMetrics {
		mvalue, exist := ms[name]
		if exist {
			res[name] = mng.Metric{MType: mtype, MValue: mvalue}
		} else if name == "PollCount" {
			res[name] = mng.Metric{MType: mtype, MValue: int64(1)}
		} else if name == "RandomValue" {
			res[name] = mng.Metric{MType: mtype, MValue: 1}
		}
	}

	return res
}

func getMemStatus() map[string]interface{} {
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)
	v := reflect.ValueOf(ms)

	count := v.NumField()
	res := make(map[string]interface{})

	for i := 0; i < count; i++ {
		v1 := v.Field(i)
		t := v.Type()
		res[t.Field(i).Name] = v1.Interface()
	}

	return res
}
