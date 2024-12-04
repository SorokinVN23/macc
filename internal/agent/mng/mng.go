package nmg

import (
	"errors"
	"fmt"
	"macc/utils"
	"net/http"
	"strings"
	"time"
)

type Manager struct {
	producer         MetricsProducer
	adapter          Adapter
	necessaryMetrics map[string]MetricType
	collecting       bool
	client           *http.Client
	sending          bool
}

func New(producer MetricsProducer, adapter Adapter) Manager {
	manager := Manager{producer: producer, adapter: adapter}
	manager.necessaryMetrics = getNecessaryMetrics()
	manager.client = &http.Client{}
	return manager
}

func (manager Manager) GetSending() bool {
	return manager.sending
}

func (manager Manager) GetCollecting() bool {
	return manager.collecting
}

func (manager Manager) StartSending(srvurl string, reportInterval int) {
	manager.sending = true
	go func() {
		defer func() {
			manager.sending = false
		}()

		for manager.sending {

			for name, mtype := range manager.necessaryMetrics {
				mvalue, err := manager.getFromStore(name, mtype)
				if err != nil {
					panic(err)
				}

				url1, _ := strings.CutSuffix(srvurl, "/")
				url := fmt.Sprintf("http://%s/update/%s/%s/%v", url1, mtype, name, mvalue)

				resp, err := manager.client.Post(url, "text/plain", nil)
				if err != nil {
					fmt.Printf("%s\n", url)
					fmt.Printf("%+v\n", err)
					continue
				}
				if resp.StatusCode != 200 {
					fmt.Printf("%s\n", url)
					fmt.Printf("resp.StatusCode %v\n", resp.StatusCode)
					continue
				}

				if mtype == MTCounter {
					err = manager.reset(name, mtype)
					if err != nil {
						panic(err)
					}
				}

				//fmt.Printf("%s %v\n", url, resp.StatusCode)
			}

			time.Sleep(time.Duration(reportInterval) * time.Second)
		}
	}()
}

func (manager Manager) StopSending() {
	manager.sending = false
}

func (manager Manager) StopCollecting() {
	manager.collecting = false
}

func (manager Manager) StartCollecting(poolInterval int) {
	manager.collecting = true
	go func() {
		defer func() {
			manager.collecting = false
		}()

		for manager.collecting {
			list, err := manager.producer.GetNewMetricsValue(manager.necessaryMetrics)
			if err != nil {
				panic(err)
			}

			for name, metric := range list {
				err := manager.updateStore(name, metric)
				if err != nil {
					panic(err)
				}
			}

			time.Sleep(time.Duration(poolInterval) * time.Second)
		}
	}()
}

func ConvertByType(mt MetricType, val interface{}) (interface{}, error) {
	switch mt {
	case MTCounter:
		v, err := utils.ConvertToInt64(val)
		return interface{}(v), err
	case MTGauge:
		v, err := utils.ConvertToFloat64(val)
		return interface{}(v), err
	}

	return nil, fmt.Errorf("Invalid MetricType %+v in ConvertByType ", mt)
}

func (manager Manager) getFromStore(name string, mtype MetricType) (interface{}, error) {
	switch mtype {
	case MTCounter:
		currentVolume, err := manager.adapter.GetInt64(name, string(mtype))
		if err != nil {
			return nil, err
		}
		return currentVolume, nil
	case MTGauge:
		currentVolume, err := manager.adapter.GetFloat64(name, string(mtype))
		if err != nil {
			return nil, err
		}
		return currentVolume, nil
	}
	return nil, fmt.Errorf("invalid mtype %+v", mtype)
}

func (manager Manager) reset(name string, mtype MetricType) error {
	switch mtype {
	case MTCounter:
		err := manager.adapter.SetInt64(name, string(mtype), int64(0))
		if err != nil {
			return err
		}
		return nil
	case MTGauge:
		err := manager.adapter.SetFloat64(name, string(mtype), float64(0))
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("invalid mtype %+v", mtype)
}

func (manager Manager) updateStore(name string, metric Metric) error {
	switch metric.MType {
	case MTCounter:
		currentVolume, err := manager.adapter.GetInt64(name, string(metric.MType))
		if err != nil {
			return err
		}
		res, err := utils.Sum(currentVolume, metric.MValue)
		if err != nil {
			return err
		}
		err = manager.adapter.SetInt64(name, string(metric.MType), res)
		if err != nil {
			return err
		}
	case MTGauge:
		err := manager.adapter.SetFloat64(name, string(metric.MType), metric.MValue)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid metric.MType")
	}
	return nil
}
