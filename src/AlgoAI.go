package main



func AlgoAI(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	res := make(map[AgentID]AgentID, len(Elèves))
	for len(Elèves) != 0 {
		Elèves, Universités, res = tourAI(Elèves, Universités, res)
	}
	return res
}

func tourAI(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) ([]Agent, []Agent, map[AgentID]AgentID) {
	choixElèves := DicoChoixElèves(Elèves)
	choixUni := DicoChoixUnis(choixElèves, Universités)
	Elèves, Universités = RemoveTuples(choixUni, Elèves, Universités)
	MergeAgentIDMaps(res, choixUni)
	return Elèves, Universités, res
}

func DicoChoixElèves(Elèves []Agent) map[AgentID][]AgentID {
	choixElèves := make(map[AgentID][]AgentID)

	for _, élève := range Elèves {
		prefUni := élève.Prefs[0]
		choixElèves[prefUni] = append(choixElèves[prefUni], élève.ID)
	}
	return choixElèves
}