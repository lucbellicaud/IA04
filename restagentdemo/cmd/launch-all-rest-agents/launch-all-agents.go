package main

import (
	"math/rand"
	"fmt"
	"log"

	"gitlab.utc.fr/lagruesy/restagentdemo/restclientagent"
	"gitlab.utc.fr/lagruesy/restagentdemo/restserveragent"

	rad "gitlab.utc.fr/lagruesy/restagentdemo"
)

func shuffle(src []rad.AgentID) []rad.AgentID {
	final := make([]rad.AgentID, len(src))
	perm := rand.Perm(len(src))

	for i, v := range perm {
			final[v] = src[i]
	}
	return final
}

func main() {
	const n = 100
	const url1 = ":8000"
	const url2 = "http://localhost:8000"

	clAgts := make([]restclientagent.RestClientAgent, 0, n)
	
	var listeVotants []rad.AgentID


	log.Println("démarrage des clients...")
	preflist := []rad.AgentID{"Test","Test2","Test3"}
	for i := 0; i < n; i++ {
		id := fmt.Sprintf("id%02d", i)
		randomlist := shuffle(preflist)
		agt := restclientagent.NewRestClientAgent(id, url2, randomlist)
		clAgts = append(clAgts, *agt)
		listeVotants = append(listeVotants, rad.AgentID(id))
	}

	servAgt := restserveragent.NewRestServerAgent(url1,listeVotants)
	log.Println("démarrage du serveur...")
	go servAgt.Start()

	for _, agt := range clAgts {
		// attention, obligation de passer par cette lambda pour faire capturer la valeur de l'itération par la goroutine
		func(agt restclientagent.RestClientAgent) {
			go agt.Start()
		}(agt)
	}

	fmt.Scanln()
}
