package main

import (
	"fmt"
)

func DynamiqueLibre(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) map[AgentID]AgentID {
	fmt.Println("Nouveau tour")
	reverseRes := reverseMap(res)
	out1:
	for _,élève := range Elèves {
		for _,université := range Universités {	
			if élève.PrefersWithID(université.ID, res[élève.ID]) && université.PrefersWithID(élève.ID,reverseRes[université.ID]) {
				// fmt.Println("Couple instable :",élève.ID, res[élève.ID])
				fmt.Println("Elève ",élève.ID, " préfère ", université.ID, " à", res[élève.ID])
				fmt.Println("Université ",université.ID," préfère ",élève.ID, " à", reverseRes[université.ID])
				res[reverseRes[université.ID]]=res[élève.ID]
				reverseRes[res[élève.ID]]=reverseRes[université.ID]
				res[élève.ID]=université.ID
				reverseRes[université.ID]=élève.ID
				// fmt.Println("res",res)
				// fmt.Println("reverseres",reverseRes)

				continue out1
			}
		}
	}
	fmt.Println("Fin du dynamique libre")
	Score(Elèves,Universités,res)
	return res
}

func Score(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID){
	reverseRes := reverseMap(res)
	scoreEleves:=0
	for _,élève := range Elèves{
		score,_ := élève.RankWithID(res[élève.ID])
		scoreEleves += len(Elèves)- score
	}
	fmt.Println("Score de l'algo pour les élèves :",scoreEleves)
	scoreUnis:=0
	for _,uni := range Universités{
		score,_ := uni.RankWithID(reverseRes[uni.ID])
		scoreUnis += len(Elèves)- score
	}
	fmt.Println("Score de l'algo pour les unis :",scoreUnis)
	fmt.Println("Score total :",scoreEleves+scoreUnis)
}