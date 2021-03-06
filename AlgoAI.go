package main

import "fmt"

func AlgoAI(Elèves []Agent, Universités []Agent) map[AgentID]AgentID {
	//Algorithme de tour par tour. Chaque tour, les élèves donnent leur choix préféré, et les universités prennent leur préféré.
	res := make(map[AgentID]AgentID, len(Elèves))
	for round := 0; len(Universités) != 0; round++ { //Boucle tour par tour
		fmt.Println("Tour : ", round+1)
		new := tourAI(Elèves, Universités, round)
		fmt.Println("Nouvelles affectations du tour", new)
		for uni, eleve := range new {
			res[uni] = eleve
			Elèves = Supprimer(Elèves, Trouver(Elèves, eleve))
			Universités = Supprimer(Universités, Trouver(Universités, uni))
		}
		fmt.Println("Liste à jour sur le tour ", round, " ", res)
	}

	return res
}

func tourAI(Elèves []Agent, Universités []Agent, round int) map[AgentID]AgentID {
	//Cette fonction retourne une map avec les choix validés à chaque tour
	choixEleves := make(map[AgentID][]AgentID)

	//Prend les choix préférés des élèves à ce tour
	for _, eleve := range Elèves {
		prefUni := RetournerFavoriSelonTour(eleve.Prefs, Universités, round)
		choixEleves[prefUni] = append(choixEleves[prefUni], eleve.ID)
	}

	fmt.Println("Choix des élèves : ", choixEleves)
	choixUni := make(map[AgentID]AgentID)

	for uni, eleves := range choixEleves { // L'université prend son élève préféré
		choixUni[uni] = eleves[0]
		if len(eleves) >= 2 {
			for _, eleve := range eleves[1:] {
				ag, _ := GetAgent(uni, Universités)
				neweleve, _ := GetAgent(eleve, Elèves)
				oldeleve, _ := GetAgent(choixUni[uni], Elèves)
				fmt.Println("Comparaison de ", oldeleve, " avec ", neweleve)
				pref, _ := ag.Prefers(neweleve, oldeleve)
				if pref {
					fmt.Println("Remplacement de ", oldeleve, " par ", neweleve)
					choixUni[uni] = eleve
				}
			}
		}
	}

	return choixUni
}

func RetournerFavoriSelonTour(prefListe []AgentID, Universités []Agent, round int) (uni AgentID) {
	//Cette fonction retourne le premier choix qui n'est pas déjà pris.
	for i := round; i <= len(prefListe[round]); round++ {
		for _, uni := range Universités {
			if uni.ID == prefListe[round] {
				return uni.ID
			}
		}

	}
	return AgentID("")
}

func Trouver(a []Agent, x AgentID) int {
	for i, n := range a {
		if x == n.ID {
			return i
		}
	}
	return len(a)
}

func Trouver_ID(eleve AgentID, liste_eleves []AgentID) int {
	for i, val := range liste_eleves {
		if val == eleve {
			return i
		}
	}
	return len(liste_eleves)
}

func Supprimer(a []Agent, i int) []Agent {
	if i == len(a) {
		return a[:i]
	}
	return append(a[:i], a[i+1:]...)
}
