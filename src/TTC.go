package main

import "fmt"

func TTC(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	fmt.Println("Début de l'ago TTC")
	res := make(map[AgentID]AgentID, len(Elèves))
	for len(Elèves) != 0 {
		liste_choix := InitGraphTTC(Elèves, Universités)
		Elèves, Universités, res = tourTTC(liste_choix, Elèves, Universités, res, 0)
	}

	fmt.Print("Fin de l'ago TTC\n\n")
	return res
}

func InitGraphTTC(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	choix := make(map[AgentID]AgentID)
	for _, élève := range Elèves {
		UniPref := élève.Prefs[0]
		choix[élève.ID] = UniPref
	}
	for _, uni := range Universités {
		ElevePref := uni.Prefs[0]
		choix[uni.ID] = ElevePref
	}
	return choix
}

func tourTTC(graph map[AgentID]AgentID, Elèves []Agent, Universités []Agent, res map[AgentID]AgentID, i int) ([]Agent, []Agent, map[AgentID]AgentID) {
	var tortue AgentID
	var lièvre AgentID

	for _,élève := range Elèves { //Initialiser le lièvre et la tortue
		if _, ok := graph[élève.ID]; ok { // Si l'élève n'a pas d'université attribuée #rip
			tortue = graph[élève.ID]
			lièvre = graph[graph[élève.ID]]
		}
	}

	for lièvre != tortue  { //Cherche un cycle
		tortue = graph[tortue]
		lièvre = graph[graph[lièvre]]
	}

	temp:=graph[lièvre]
	
	for temp!=lièvre{
		temp=graph[temp]
	}

	to_remove := make(map[AgentID]AgentID)
	to_remove[tortue] = graph[tortue]
	élève := graph[graph[tortue]]
	uni := graph[élève]

	for élève != tortue {
		to_remove[élève] = uni
		élève = graph[uni]
		uni = graph[élève]
	}

	Elèves, Universités = SupprimerTuples(to_remove, Elèves, Universités)
	res = ConcaténerAgentIDDic(res, to_remove)
	graph = NettoyerGraph(to_remove,graph)

	if len(graph)==0{
		return Elèves,Universités,res
	}else{
		return tourTTC(graph, Elèves, Universités, res, i+1)
	}

	
}

func NettoyerGraph(to_remove map[AgentID]AgentID, graph map[AgentID]AgentID) map[AgentID]AgentID {
	for élève,uni :=range to_remove{
		for i,v := range graph{
			if élève==i || uni==v || élève==v || uni==i{
				delete(graph,i)
				to_remove[i]=v
				return NettoyerGraph(to_remove,graph)
			}
		}
	}
	return graph
}