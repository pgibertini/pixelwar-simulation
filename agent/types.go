package agent

type RegisterAWRequest struct {
	Address *AgentWorker `json:"address"`
}

type RegisterAMRequest struct {
	Address *AgentManager `json:"address"`
}
