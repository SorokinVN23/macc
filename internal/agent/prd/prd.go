package prd

import (
	mng "macc/internal/agent/mng"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

var myRand *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	myRand = rand.New(source)
}

type Producer struct{}

func (p Producer) GetNewMetricsValue(necessaryMetrics map[string]mng.MetricType) (map[string]mng.Metric, error) {
	res := make(map[string]mng.Metric)
	ms := getMemStatus()

	for name, mtype := range necessaryMetrics {
		mvalue, exist := ms[name]
		if exist {
			val, err := mng.ConvertByType(mtype, mvalue)
			if err != nil {
				return nil, err
			}
			res[name] = mng.Metric{MType: mtype, MValue: val}
		} else if name == "PollCount" {
			res[name] = mng.Metric{MType: mtype, MValue: interface{}(int64(1))}
		} else if name == "RandomValue" {
			randomFloat := myRand.Float64()
			res[name] = mng.Metric{MType: mtype, MValue: interface{}(randomFloat)}
		}
	}

	return res, nil
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
