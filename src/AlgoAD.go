package main

func AlgoAD(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	res := make(map[AgentID]AgentID, len(Elèves))
	round := 0
	change := true
	for change {
		Elèves, Universités, res, change = tourAD(Elèves, Universités, res, round)
	}
	return res
}

func tourAD(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID, round int) ([]Agent, []Agent, map[AgentID]AgentID, bool) {
	var change bool
	choixElèves := DicoChoixElèves(Elèves)
	choixUni := DicoChoixUnis(choixElèves, Universités)
	Elèves, change = SupprimerNonAcceptés(Elèves, choixElèves, choixUni)
	MergeAgentIDMaps(res, choixUni)
	return Elèves, Universités, res, change
}

func SupprimerNonAcceptés(Elèves []Agent, choixElèves map[AgentID][]AgentID, choixUni map[AgentID]AgentID) ([]Agent, bool) {
	change := false
	for i, élève := range Elèves {
		if _, ok := choixUni[élève.ID]; !ok { // Si l'élève n'a pas d'université attribuée #rip
			Elèves[i].Prefs = élève.Prefs[1:]
			change = true
		}

	}
	return Elèves, change
}