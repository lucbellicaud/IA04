package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type AgentID string

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []AgentID
}

func NewAgent(id AgentID, name string, prefs []AgentID) Agent {
	return Agent{id, name, prefs}
}

func Equal(ag1 Agent, ag2 Agent) bool {
	return ag1.ID == ag2.ID
}

func (ag Agent) RankWithID(agID1 AgentID) (int, error) {
	for i, v := range ag.Prefs {
		if v == agID1 {
			return i, nil
		}
	}
	return -1, errors.New("Agent %s not found" + string(agID1))
}

func (ag Agent) PrefersWithID(agID1 AgentID, agID2 AgentID) (bool){
	rank1, err1 := ag.RankWithID(agID1)
	rank2, err2 := ag.RankWithID(agID2)

	if err1 != nil {
		log.Fatal(errors.New("Agent %s not found" + string(agID1)))
	}

	if err2 != nil {
		log.Fatal(errors.New("Agent %s not found" + string(agID2)))
	}
	return rank1 < rank2


}

func RandomPrefs(ids []AgentID) (res []AgentID) {
	res = make([]AgentID, len(ids))
	copy(res, ids)
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	return
}

func ReturnFavorite(listeElèves []AgentID,Université Agent) (AgentID,error){
	favorite := listeElèves[0]
	for _,élève := range listeElèves {
		if  Université.PrefersWithID(élève,favorite) {
			favorite = élève
		}
	}
	return favorite, nil
}

func Trouver(a []Agent, x AgentID) int {
	for i, n := range a {
		if x == n.ID {
			return i
		}
	}
	log.Fatal(errors.New("l'agent n'a pas été trouvé"))
	return len(a)
}

func Trouver_ID(élève AgentID, liste_élèves []AgentID) int {
	for i, val := range liste_élèves {
		if val == élève {
			return i
		}
	}
	fmt.Println(élève,liste_élèves)
	return len(liste_élèves)
}

func Remove(s []Agent, i int) []Agent {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func OrderedRemove(slice []AgentID, i int) []AgentID {
	ret := make([]AgentID, len(slice)-1)
	j:=0
	for j<i {
		ret[j]=slice[j]
		j++
	}
	for j<len(slice)-1 {
		ret[j]=slice[j+1]
		j++
	}
	return ret
}

func MergeAgentIDMaps(a map[AgentID]AgentID, b map[AgentID]AgentID) map[AgentID]AgentID {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func RemoveTuples(to_remove map[AgentID]AgentID, Elèves []Agent, Universités []Agent) ([]Agent, []Agent) {
	for élève, uni := range to_remove {
		Elèves = Remove(Elèves, Trouver(Elèves, élève))
		for i,v := range Elèves{
			Elèves[i].Prefs = OrderedRemove(v.Prefs, Trouver_ID(uni,v.Prefs))
		}
		Universités = Remove(Universités, Trouver(Universités, uni))
		for i,v := range Universités{
			Universités[i].Prefs = OrderedRemove(v.Prefs, Trouver_ID(élève,v.Prefs))
		}

	}
	return Elèves, Universités
}

func DicoChoixUnis(choixElèves map[AgentID][]AgentID, Universités []Agent) map[AgentID]AgentID{
	choixUnis := make(map[AgentID]AgentID)
	for _,uni := range Universités {
		_, exists := choixElèves[uni.ID]
		if exists {
			élève,err1 := ReturnFavorite(choixElèves[uni.ID],uni)
			if err1!=nil{
				log.Fatal(err1)
			}
			choixUnis[élève] = uni.ID
		}
	}
	return choixUnis
}