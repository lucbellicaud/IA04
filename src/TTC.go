package main

import "fmt"

func TTC(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	res := make(map[AgentID]AgentID, len(Elèves))
	for len(Elèves) != 0 {
		liste_choix := GetChoicesTTC(Elèves, Universités)
		Elèves, Universités, res = TortoiseAndHare(liste_choix, Elèves, Universités, res, 0)
	}

	return res
}

func GetChoicesTTC(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
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

func TortoiseAndHare(graph map[AgentID]AgentID, Elèves []Agent, Universités []Agent, res map[AgentID]AgentID, i int) ([]Agent, []Agent, map[AgentID]AgentID) {
	fmt.Println("Tour ",i)
	fmt.Println("graph reçu :\n",graph)
	var tortoise AgentID
	var hare AgentID

	for _,élève := range Elèves { //Initialiser le lièvre et la tortue
		if _, ok := graph[élève.ID]; ok { // Si l'élève n'a pas d'université attribuée #rip
			tortoise = graph[élève.ID]
			hare = graph[graph[élève.ID]]
		}
	}

	fmt.Println("Tortue à l'état initial", tortoise, "Lièvre à l'état initial", hare)

	for hare != tortoise  { //Cherche un cycle
		tortoise = graph[tortoise]
		fmt.Println(hare)
		hare = graph[graph[hare]]
	}


	temp:=graph[hare]
	
	for temp!=hare{
		temp=graph[temp]
	}

	to_remove := make(map[AgentID]AgentID)
	to_remove[tortoise] = graph[tortoise]
	élève := graph[graph[tortoise]]
	uni := graph[élève]

	for élève != tortoise {
		to_remove[élève] = uni
		élève = graph[uni]
		uni = graph[élève]
	}

	Elèves, Universités = RemoveTuples(to_remove, Elèves, Universités)
	res = MergeAgentIDMaps(res, to_remove)
	graph = CleanGraph(to_remove,graph)

	if len(graph)==0{
		return Elèves,Universités,res
	}

	return TortoiseAndHare(graph, Elèves, Universités, res, i+1)
}

func CleanGraph(to_remove map[AgentID]AgentID, graph map[AgentID]AgentID) map[AgentID]AgentID {
	for élève,uni :=range to_remove{
		for i,v := range graph{
			if élève==i || uni==v || élève==v || uni==i{
				delete(graph,i)
				to_remove[i]=v
				return CleanGraph(to_remove,graph)
			}
		}
	}
	return graph
}