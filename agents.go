package wazuhapi

import (
	"encoding/json"
)

type AgentsStatus struct {
	Active       int
	Disconnected int
	Total        int
	Synced       int
	NotSynced    int
}

func (w *WazuhAPI) GetAgents() (agents AgentsStatus, err error) {
	resp, err := w.Call("/overview/agents", "GET", "")
	if err != nil {
		return AgentsStatus{}, err
	}

	type Response struct {
		Data struct {
			AgentStatus struct {
				Connection struct {
					Active       int `json:"active"`
					Disconnected int `json:"disconnected"`
					Total        int `json:"total"`
				}
				Configuration struct {
					Synced    int `json:"synced"`
					NotSynced int `json:"not_synced"`
				}
			} `json:"agent_status"`
		}
	}

	var response Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return AgentsStatus{}, err
	}

	agents = AgentsStatus{
		Active:       response.Data.AgentStatus.Connection.Active,
		Disconnected: response.Data.AgentStatus.Connection.Disconnected,
		Total:        response.Data.AgentStatus.Connection.Total,
		Synced:       response.Data.AgentStatus.Configuration.Synced,
		NotSynced:    response.Data.AgentStatus.Configuration.NotSynced,
	}

	return agents, nil
}
