package main

func TTC(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	res := make(map[AgentID]AgentID, len(Elèves))
	for len(Elèves)!=0{
		choixEleves := GetElevesChoicesTTC(Elèves, Universités)
		prefsUni := GetUniChoicesTTC(choixEleves,Elèves,Universités)
		cycles := FindCycles(prefsUni)
		res,eleves_to_remove,unis_to_remove := RemoveCycles(prefsUni)
		Elèves,Universités := RemoveFromLists(Elèves,Universités,eleves_to_remove,unis_to_remove)

	}
	


	return res
}

func GetElevesChoicesTTC(Elèves []Agent, Universités []Agent) map[AgentID][]AgentID {
	choixEleves:= make(map[AgentID][]AgentID)
	for eleve := range Elèves{
		prefUni := eleve.Prefs[0]
		choixEleves[prefUni] = append(choixEleves[prefUni], eleve.ID)
	}
	return choixEleves
}

func GetUniChoicesTTC(choixEleves map[AgentID][]AgentID, Elèves []Agent, Universités []Agent){
	
}