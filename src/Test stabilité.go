package main

import (
	"fmt"
	"time"
)

func DynamiqueLibre(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) map[AgentID]AgentID {
	fmt.Println("Nouveau tour")
	out1:
		for élèveRes, uniRes := range res {
			for _,élève := range Elèves {
				for _,université := range Universités {
					if élève.PrefersWithID(université.ID, uniRes) && université.PrefersWithID(élèveRes,élève.ID) {
						fmt.Println("Couple instable :",élèveRes, uniRes)
						fmt.Println("Elève ",élèveRes, " préfère ", université.ID, " à", uniRes)
						fmt.Println("Université ",université.ID," préfère ",élèveRes, " à", élève.ID)
						time.Sleep(time.Second)
						res[élèveRes] = université.ID
						res[élève.ID] = uniRes
						continue out1
					}
				}
			}
		}
	return res
}