package main

import (
	"macc/internal/adapter"
	mng "macc/internal/agent/mng"
	"macc/internal/agent/prd"
	dom "macc/internal/domains"
	"time"
)

func main() {
	settings := &dom.AgentSettings{}
	GetFlags(settings)

	memAd := adapter.NewMemoryAdapter()
	producer := prd.Producer{}
	manager := mng.New(producer, memAd)
	manager.StartCollecting(settings.PollInterval)
	manager.StartSending(settings.Address, settings.ReportInterval)

	time.Sleep(10 * time.Second)
	for !manager.GetCollecting() && !manager.GetSending() {
		time.Sleep(10 * time.Second)
	}
}
