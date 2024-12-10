package nmg

import (
	"errors"
	"fmt"
	"macc/utils"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Manager struct {
	producer         MetricsProducer
	adapter          Adapter
	necessaryMetrics map[string]MetricType
	client           *http.Client
	Collecting       *chan struct{}
	Sending          *chan struct{}
	WaitGroup        *sync.WaitGroup
}

func New(producer MetricsProducer, adapter Adapter) Manager {
	manager := Manager{producer: producer, adapter: adapter}
	manager.necessaryMetrics = getNecessaryMetrics()
	manager.client = &http.Client{}
	manager.WaitGroup = &sync.WaitGroup{}
	return manager
}

func (manager Manager) StartSending(srvurl string, reportInterval int) {
	stopChan := make(chan struct{})
	manager.WaitGroup.Add(1)

	go func() {
		defer func() {
			close(stopChan)
			manager.WaitGroup.Done()
		}()

		for {
			select {
			case <-stopChan:
				return
			default:
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

					resp.Body.Close()
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
				}

				time.Sleep(time.Duration(reportInterval) * time.Second)
			}
		}
	}()
	manager.Sending = &stopChan
}

func (manager Manager) StopSending() {
	close(*manager.Sending)
}

func (manager Manager) StopCollecting() {
	close(*manager.Collecting)
}

func (manager Manager) StartCollecting(poolInterval int) {
	stopChan := make(chan struct{})
	manager.WaitGroup.Add(1)

	go func() {
		defer func() {
			close(stopChan)
			manager.WaitGroup.Done()
		}()

		for {
			select {
			case <-stopChan:
				return
			default:
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
		}
	}()
	manager.Collecting = &stopChan
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

	return nil, fmt.Errorf("invalid MetricType %+v in ConvertByType ", mt)
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
