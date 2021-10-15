package main

import "fmt"

func AlgoAD(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	//Algorithme de tour par tour. Chaque tour, les élèves donnent leur choix préféré, et les universités prennent leur préféré.
	res := make(map[AgentID]AgentID, len(Elèves))
	done := false

	for round := 0; done == false; round++ { //Boucle tour par tour
		res, done, Elèves, Universités = tourAD(Elèves, Universités, res)
		fmt.Println("Liste à jour sur le tour ", round, " ", res)
		fmt.Println(done)
	}

	return res
}

func tourAD(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) (map[AgentID]AgentID, bool, []Agent, []Agent) {
	done := true
	choixEleves := make(map[AgentID][]AgentID)

	//Prend les choix préférés des élèves à ce tour
	for _, eleve := range Elèves {
		prefUni := eleve.Prefs[0]
		choixEleves[prefUni] = append(choixEleves[prefUni], eleve.ID)
	}
	fmt.Println("Choix des élèves", choixEleves)

	for uni, eleves := range choixEleves { // L'université prend son élève préféré
		for _, eleve := range eleves {
			_, ok := res[uni]
			if !ok {
				res[uni] = eleve
			} else {
				ag, _ := GetAgent(uni, Universités)
				neweleve, _ := GetAgent(eleve, Elèves)
				oldeleve, _ := GetAgent(res[uni], Elèves)
				if !Equal(neweleve, oldeleve) {
					done = false
					fmt.Println("Comparaison de ", oldeleve, " avec ", neweleve)
					pref, _ := ag.Prefers(neweleve, oldeleve)
					if pref {
						Elèves = SupprimerEleve(Elèves, oldeleve)
						fmt.Println("Remplacement de ", oldeleve, " par ", neweleve)
						res[uni] = eleve
					} else {
						Elèves = SupprimerEleve(Elèves, neweleve)
						fmt.Println("Pas de remplacement de ", oldeleve, " par ", neweleve)
					}
				}

			}

		}
	}
	fmt.Println(Elèves)
	return res, done, Elèves, Universités
}

func SupprimerEleve(pool []Agent, ag Agent) []Agent {
	//fmt.Println("Avant", pool)
	j := -1
	for i, v := range pool {
		if Equal(v, ag) {
			j = i
			break
		}
	}
	pool[j].Prefs = pool[j].Prefs[1:]
	return pool

}
