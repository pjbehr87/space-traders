package stlib

import (
	"encoding/json"
)

type AgentJson struct {
	AccountId       string `json:"accountId"`
	Symbol          string
	Headquarters    string
	Credits         int64
	StartingFaction string `json:"startingFaction"`
}
type Agent struct {
	AccountId       string `json:"accountId"`
	Symbol          string
	Headquarters    *WaypointSymbol
	Credits         int64
	StartingFaction string `json:"startingFaction"`
}

type AgentData struct {
	Agent AgentJson `json:"data"`
}

func (stl *StLib) MyAgent() (Agent, error) {
	resp, err := stl.GetUrl("my/agent")
	if err != nil {
		return Agent{}, err
	}

	agentData := AgentData{}
	err = json.Unmarshal(resp, &agentData)
	if err != nil {
		return Agent{}, err
	}

	return Agent{
		AccountId:       agentData.Agent.AccountId,
		Symbol:          agentData.Agent.Symbol,
		Headquarters:    getWps(agentData.Agent.Headquarters),
		Credits:         agentData.Agent.Credits,
		StartingFaction: agentData.Agent.StartingFaction,
	}, nil
}
