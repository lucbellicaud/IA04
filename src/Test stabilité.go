package main

import (
	"fmt"
	"time"
)

func DynamiqueLibre(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) {
	out:
	for élèveRes, uniRes := range res {
		for _,élève := range Elèves {
			for _,université := range Universités {
				if élève.PrefersWithID(université.ID, uniRes) && université.PrefersWithID(élève.ID, élèveRes) {
					fmt.Println("Couple instable :",élèveRes, uniRes)
					time.Sleep(time.Second)
					res[élèveRes] = université.ID
					res[élève.ID] = uniRes
					break out
				}
			}
		}
	}
}