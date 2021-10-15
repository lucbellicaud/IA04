package main

import (
	"errors"
	"math/rand"
)

type AgentID string

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []AgentID
}

func NewAgent(id AgentID, name string, prefs []AgentID) Agent {
	return Agent{id, name, prefs}
}

func Equal(ag1 Agent, ag2 Agent) bool {
	return ag1.ID == ag2.ID
}

func GetAgent(id AgentID, pool []Agent) (Agent, error) {
	for _, v := range pool {
		if v.ID == id {
			return v, nil
		}
	}
	return Agent{}, errors.New("Agent %s not found" + string(id))
}

func (ag Agent) Rank(ag1 Agent) (int, error) {
	for i, v := range ag.Prefs {
		if v == ag1.ID {
			return i, nil
		}
	}
	return -1, errors.New("Agent %s not found" + string(ag1.ID))
}

func (ag Agent) Prefers(ag1 Agent, ag2 Agent) (bool, error) {
	rank1, err1 := ag.Rank(ag1)
	rank2, err2 := ag.Rank(ag2)

	if err1 != nil {
		return false, errors.New("Agent %s not found" + string(ag1.ID))
	}

	if err2 != nil {
		return false, errors.New("Agent %s not found" + string(ag2.ID))
	}
	return rank1 < rank2, nil

}

func RandomPrefs(ids []AgentID) (res []AgentID) {
	res = make([]AgentID, len(ids))
	copy(res, ids)
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	return
}
