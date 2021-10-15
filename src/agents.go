package main

import (
	"errors"
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

func GetAgent(id AgentID, pool []Agent) (Agent, error) {
	for _, v := range pool {
		if v.ID == id {
			return v, nil
		}
	}
	return Agent{}, errors.New("Agent %s not found" + string(id))
}

func (ag Agent) Rank(ag1 Agent) (int, error) {
	for i, v := range ag.Prefs {
		if v == ag1.ID {
			return i, nil
		}
	}
	return -1, errors.New("Agent %s not found" + string(ag1.ID))
}

func (ag Agent) RankWithID(agID1 AgentID) (int, error) {
	for i, v := range ag.Prefs {
		if v == agID1 {
			return i, nil
		}
	}
	return -1, errors.New("Agent %s not found" + string(agID1))
}

func (ag Agent) Prefers(ag1 Agent, ag2 Agent) (bool, error) {
	rank1, err1 := ag.Rank(ag1)
	rank2, err2 := ag.Rank(ag2)

	if err1 != nil {
		return false, errors.New("Agent %s not found" + string(ag1.ID))
	}

	if err2 != nil {
		return false, errors.New("Agent %s not found" + string(ag2.ID))
	}
	return rank1 < rank2, nil

}

func (ag Agent) PrefersWithID (agID1 AgentID, agID2 AgentID) (bool, error){
	rank1, err1 := ag.RankWithID(agID1)
	rank2, err2 := ag.RankWithID(agID2)

	if err1 != nil {
		return false, errors.New("Agent %s not found" + string(agID1))
	}

	if err2 != nil {
		return false, errors.New("Agent %s not found" + string(agID2))
	}
	return rank1 < rank2, nil


}

func RandomPrefs(ids []AgentID) (res []AgentID) {
	res = make([]AgentID, len(ids))
	copy(res, ids)
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	return
}

func ReturnFavorite(listeEleves []AgentID,Université Agent) (AgentID,error){
	favorite := listeEleves[0]
	for _,eleve := range listeEleves {
		prefersNew,err1 := Université.PrefersWithID(eleve,favorite)
		if err1 != nil {
			return AgentID(""), errors.New("Agent %s not found" + string(eleve))
		}
		if  prefersNew {
			favorite = eleve
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

func Remove(s []Agent, i int) []Agent {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func OrderedRemove(slice []AgentID, s int) []AgentID {
	return append(slice[:s], slice[s+1:]...)
}

func MergeAgentIDMaps(a map[AgentID]AgentID, b map[AgentID]AgentID) map[AgentID]AgentID {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func RemoveTuples(to_remove map[AgentID]AgentID, Elèves []Agent, Universités []Agent) ([]Agent, []Agent) {
	for eleve, uni := range to_remove {
		Elèves = Remove(Elèves, Trouver(Elèves, eleve))
		for _,v := range Elèves{
			v.Prefs = OrderedRemove(v.Prefs, Trouver_ID(uni,v.Prefs))
		}
		Universités = Remove(Universités, Trouver(Universités, uni))
		for _,v := range Universités{
			v.Prefs = OrderedRemove(v.Prefs, Trouver_ID(eleve,v.Prefs))
		}

	}
	return Elèves, Universités
}