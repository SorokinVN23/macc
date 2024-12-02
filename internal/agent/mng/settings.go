package nmg

type Metric struct {
	MType  MetricType
	MValue interface{}
}

type MetricType string

const (
	MTGauge   MetricType = "gauge"
	MTCounter MetricType = "counter"
)

func getNecessaryMetrics() map[string]MetricType {
	res := make(map[string]MetricType)

	res["Alloc"] = MTGauge
	res["BuckHashSys"] = MTGauge
	res["Frees"] = MTGauge
	res["GCCPUFraction"] = MTGauge
	res["GCSys"] = MTGauge
	res["HeapAlloc"] = MTGauge
	res["HeapIdle"] = MTGauge
	res["HeapInuse"] = MTGauge
	res["HeapObjects"] = MTGauge
	res["HeapReleased"] = MTGauge
	res["HeapSys"] = MTGauge
	res["LastGC"] = MTGauge
	res["Lookups"] = MTGauge
	res["MCacheInuse"] = MTGauge
	res["MCacheSys"] = MTGauge
	res["MSpanInuse"] = MTGauge
	res["MSpanSys"] = MTGauge
	res["Mallocs"] = MTGauge
	res["NextGC"] = MTGauge
	res["NumForcedGC"] = MTGauge
	res["NumGC"] = MTGauge
	res["OtherSys"] = MTGauge
	res["PauseTotalNs"] = MTGauge
	res["StackInuse"] = MTGauge
	res["StackSys"] = MTGauge
	res["Sys"] = MTGauge
	res["TotalAlloc"] = MTGauge
	res["RandomValue"] = MTGauge

	res["PollCount"] = MTCounter

	return res
}

type MetricsProducer interface {
	GetNewMetricsValue(necessaryMetrics map[string]MetricType) (map[string]Metric, error)
}

type Adapter interface {
	GetInt64(mname string, mtype string) (interface{}, error)
	GetFloat64(mname string, mtype string) (interface{}, error)
	SetInt64(mname string, mtype string, mvolume interface{}) error
	SetFloat64(mname string, mtype string, mvolume interface{}) error
}
