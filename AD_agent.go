package main

import (
	"fmt"
	"time"
)

type RequestEleve struct {
    idEleve AgentID
    choixUni AgentID
    c        chan bool
}

type RequestUni struct {
	idUni AgentID
	choixEleve AgentID
	c chan []AgentID
}

type AgentProposant struct {
	id    AgentID
	Name  string
	Prefs []AgentID
	cin   chan bool
	cout  chan RequestEleve
}

type AgentDisposant struct {
	id    AgentID
	Name  string
	Prefs []AgentID
	cin   chan []AgentID 
	cout  chan RequestUni
}

type AgentAppariteur struct {
	cinUniversités chan RequestUni
	cinElèves chan RequestEleve
}


func NewAgentDisposant(id AgentID, name string, prefs []AgentID, cin chan []AgentID, cout chan RequestUni) AgentDisposant {
	return AgentDisposant{id, name, prefs, cin, cout}
}

func NewAgentProposant(id AgentID, name string, prefs []AgentID, cin chan bool, cout chan RequestEleve) AgentProposant {
	return AgentProposant{id, name, prefs, cin, cout}
}

func NewAgentAppariteur(cinUniversités chan AgentID,  cinElèves chan AgentID) AgentAppariteur {
	return AgentAppariteur{ cinUniversités,cinElèves}
}

//le cout de Proposant est le même que le c du disposant

func (ag AgentProposant) Start() { // a = eleves = proposants
	go func() {
		for {
			ag.cout <- RequestEleve{ag.id,ag.Prefs[0],ag.cin} 
			accepted := <-ag.cin
			if (!accepted){
				fmt.Println("Je ne suis pas accepté :(")
				ag.Prefs = ag.Prefs[1:]
			}
		}
	}()
}

func (ag AgentDisposant) Start() { //b = uni = disposants
	go func() {
		for {
			liste_eleves := <- ag.cin
			ag.cout <- RequestUni{ag.id,retourner_préféré(liste_eleves, ag.Prefs),ag.cin} 
		}
	}()
}

func retourner_préféré(liste_eleves []AgentID, prefs []AgentID) AgentID{
	ret := liste_eleves[0]
	for _,eleve := range liste_eleves {
		if (Trouver_ID(eleve, prefs) < Trouver_ID(ret, prefs)){
			ret = eleve
		}
	}
	return ret
}

func (ag AgentAppariteur) Start(liste_universités []AgentProposant, liste_élèves []AgentDisposant) {
	choixEleves := make(map[AgentID][]AgentID)
	var wg sync.WaitGroup
	for i := 1; i <= len(liste_élèves); i++ {
		wg.Add(1)
		req := <- ag.cinElèves
		choixEleves[req.choixUni] = append(choixEleves[req.choixUni], req.idEleve)
	}
	wg.Wait()

}
