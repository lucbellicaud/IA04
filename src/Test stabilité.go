package main

import (
	"fmt"
	"time"
)

func DynamiqueLibre(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID) map[AgentID]AgentID {
	fmt.Println("Début du dynamique libre")
	reverseRes := InverserDic(res)
	out1:
	for _,élève := range Elèves {
		for _,université := range Universités {	
			if élève.PrefersWithID(université.ID, res[élève.ID]) && université.PrefersWithID(élève.ID,reverseRes[université.ID]) {
				time.Sleep(time.Microsecond)
				// fmt.Println("Couple critique :",élève.ID, res[élève.ID])
				// fmt.Println("Elève ",élève.ID, " préfère ", université.ID, " à", res[élève.ID])
				// fmt.Println("Université ",université.ID," préfère ",élève.ID, " à", reverseRes[université.ID])
				res[reverseRes[université.ID]]=res[élève.ID]
				reverseRes[res[élève.ID]]=reverseRes[université.ID]
				res[élève.ID]=université.ID
				reverseRes[université.ID]=élève.ID


				continue out1
			}
		}
	}
	fmt.Print("Fin du dynamique libre\n\n")
	Score(Elèves,Universités,res)
	return res
}

func Score(Elèves []Agent, Universités []Agent, res map[AgentID]AgentID){
	n:=float64(len(Elèves))
	reverseRes := InverserDic(res)
	var scoreEleves float64
	for _,élève := range Elèves{
		score,_ := élève.RankWithID(res[élève.ID])
		scoreEleves += n- float64(score)-1
	}
	fmt.Println("Score de l'algo pour les élèves :",scoreEleves/(n*(n-1)))
	var scoreUnis float64
	for _,uni := range Universités{
		score,_ := uni.RankWithID(reverseRes[uni.ID])
		scoreUnis += n- float64(score)-1
	}
	fmt.Println("Score de l'algo pour les unis :",scoreUnis/(n*(n-1)))
	fmt.Println("Score total :",(scoreEleves+scoreUnis)/(2*n*(n-1)))
	fmt.Print("Fin du calcul du score\n\n")
}