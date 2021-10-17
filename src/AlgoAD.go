package main

import "fmt"

func AlgoAD(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	fmt.Println("Début de l'ago AD")
	res := make(map[AgentID]AgentID, len(Elèves))
	change := true
	for change {
		Elèves, Universités, res, change = tourAD(Elèves, Universités, res)
	}
	fmt.Print("Fin de l'ago AD\n\n")
	return res
}

func tourAD(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) ([]Agent, []Agent, map[AgentID]AgentID, bool) {
	var change bool
	choixElèves := DicoChoixElèves(Elèves)
	choixUni := DicoChoixUnis(choixElèves, Universités)
	Elèves, change = SupprimerNonAcceptés(Elèves, choixElèves, choixUni)
	ConcaténerAgentIDDic(res, choixUni)
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