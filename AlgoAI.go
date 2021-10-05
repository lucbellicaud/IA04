package main

func AlgoAI(Elèves []Agent, Universités []Agent) {
	//Algorithme de tour par tour. Chaque tour, les élèves donnent leur choix préféré, et les universités prennent leur préféré.
	res := make(map[AgentID]AgentID, len(Elèves))
	for round := 0; true; round++ { //Boucle tour par tour
		res := retournerChoix(Elèves, Universités, round, res)
	}
}

//eleve.Prefs[temp] Université

func retournerChoix(Elèves []Agent, Universités []Agent, round int, res map[AgentID]AgentID) map[AgentID]AgentID {
	//Cette fonction retourne une map avec les choix validés à chaque tour
	choixEleves := make(map[Agent][]Agent)

	//Prend les choix préférés des élèves à ce tour
	for _, eleve := range Elèves {
		prefUni := ReturnFirst(eleve.Prefs, Universités, round)
		choixEleves[eleve] = append(choixEleves[prefUni], eleve)
	}

	choixUni := make(map[Agent]Agent)

	for uni, eleves := range choixEleves { // L'université prend son élève préféré
		choixUni[uni] = eleves[0]
		for _, eleve := range eleves {
			if uni.Prefs(eleve, choixUni[uni]) {
				choixUni[uni] = eleve
			}
		}
	}

}

func ReturnFirst(prefListe []AgentID, Universités []Agent, round int) (uni Agent) {
	//Cette fonction retourne le premier choix qui n'est pas déjà pris.
	for i := round; i <= len(prefListe[round]); round++ {
		for _, uni := range Universités {
			if uni.ID == prefListe[round] {
				return uni
			}
		}

	}

}
