package main

import (
	"macc/internal/adapter"
	mng "macc/internal/agent/mng"
	"macc/internal/agent/prd"
	"time"
)

func main() {
	memAd := adapter.NewMemoryAdapter()
	producer := prd.Producer{}
	manager := mng.New(producer, memAd)
	manager.StartCollecting()
	manager.StartSending("http://localhost:8080")

	time.Sleep(10 * time.Second)
	for !manager.GetCollecting() && !manager.GetSending() {
		time.Sleep(10 * time.Second)
	}
}
