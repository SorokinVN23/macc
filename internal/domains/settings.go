package domains

type SrvSettings struct {
	Address string
}

type AgentSettings struct {
	Address        string
	ReportInterval int
	PollInterval   int
}
