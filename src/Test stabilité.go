package main

import (
	"fmt"
	"time"
)

func DynamiqueLibre(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) map[AgentID]AgentID {
	fmt.Println("Nouveau tour")
	reverseRes := reverseMap(res)
	out1:
	for _,élève := range Elèves {
		for _,université := range Universités {	
			if élève.PrefersWithID(université.ID, res[élève.ID]) && université.PrefersWithID(élève.ID,reverseRes[université.ID]) {
				fmt.Println("Couple instable :",élève.ID, res[élève.ID])
				fmt.Println("Elève ",élève.ID, " préfère ", université.ID, " à", res[élève.ID])
				fmt.Println("Université ",université.ID," préfère ",élève.ID, " à", reverseRes[université.ID])
				time.Sleep(time.Second)
				res[reverseRes[université.ID]]=res[élève.ID]
				reverseRes[res[élève.ID]]=reverseRes[université.ID]
				res[élève.ID]=université.ID
				reverseRes[université.ID]=élève.ID
				fmt.Println("res",res)
				fmt.Println("reverseres",reverseRes)

				continue out1
			}
		}
	}
	fmt.Println("Fin du dynamique libre")
	return res
}

func reverseMap(m map[AgentID]AgentID) map[AgentID]AgentID {
    n := make(map[AgentID]AgentID, len(m))
    for k, v := range m {
        n[v] = k
    }
    return n
}