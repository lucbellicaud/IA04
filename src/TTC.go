package main

import (
	"fmt"
)

func TTC(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	res := make(map[AgentID]AgentID, len(Elèves))
	for len(Elèves) != 0 {
		liste_choix := GetChoicesTTC(Elèves, Universités)
		Elèves, Universités, res = TortoiseAndHare(liste_choix, Elèves, Universités, res)
	}

	return res
}

func GetChoicesTTC(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	choix := make(map[AgentID]AgentID)
	for _, eleve := range Elèves {
		UniPref := eleve.Prefs[0]
		choix[eleve.ID] = UniPref
	}
	for _, uni := range Universités {
		ElevePref := uni.Prefs[0]
		choix[uni.ID] = ElevePref
	}
	fmt.Println(choix)
	return choix
}

func TortoiseAndHare(graph map[AgentID]AgentID, Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) ([]Agent, []Agent, map[AgentID]AgentID) {
	tortoise := graph[Elèves[0].ID]
	hare := graph[graph[Elèves[0].ID]]
		

	for hare != tortoise { //Cherche un cycle
		tortoise = graph[tortoise]
		hare = graph[graph[hare]]
	}

	to_remove := make(map[AgentID]AgentID)
	to_remove[tortoise] = graph[tortoise]
	eleve := graph[graph[tortoise]]
	uni := graph[eleve]

	for eleve != tortoise {
		to_remove[eleve] = uni
		eleve = graph[uni]
		uni = graph[eleve]
	}
	fmt.Println("To remove",to_remove)
	Elèves, Universités = RemoveTuples(to_remove, Elèves, Universités)
	res = MergeAgentIDMaps(res, to_remove)

	return Elèves, Universités, res
}


