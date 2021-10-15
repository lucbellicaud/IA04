package main

import (
	"fmt"
)

func PrintAg(ag Agent) {
	fmt.Println(ag)
}

func PrintInt(i int) {
	fmt.Println(i)
}

func main() {
	Anames := [...]string{
		"Khaled",
		"Sylvain",
		"Emmanuel",
		"Bob",
		"Michel",
		"Bernard",
		"Jean-Claude",
		"Jean-René",
		"L'abbé Pierre",
		"Luc",
		"Nathan",
		"Pierre",
		"Paul",
		"Jacques",
		"René",
		"Albert",
	}
	Bnames := [...]string{
		"Nathalie",
		"Annaïck",
		"Brigitte",
		"Le 4ème",
		"le 5",
		"le 6",
		"Uni 7",
		"UTC",
		"X",
		"Les mines",
		"INSA LYON",
		"INSA TOULOUSE",
		"INSA STR",
		"INSA CVL",
		"UTT",
		"UTBM",
	}

	Elèves := make([]Agent, 0, len(Anames))
	Universités := make([]Agent, 0, len(Bnames))

	groupA_prefix := "a"
	groupB_prefix := "b"

	prefsA := make([]AgentID, len(Anames))
	prefsB := make([]AgentID, len(Bnames))
	for i := 0; i < len(Anames); i++ {
		prefsA[i] = AgentID(groupA_prefix + fmt.Sprintf("%d", i))
	}

	for i := 0; i < len(Bnames); i++ {
		prefsB[i] = AgentID(groupB_prefix + fmt.Sprintf("%d", i))
	}

	for i := 0; i < len(Anames); i++ {
		prefs := RandomPrefs(prefsB)
		a := Agent{prefsA[i], Anames[i], prefs}
		Elèves = append(Elèves, a)
	}

	for i := 0; i < len(Bnames); i++ {
		prefs := RandomPrefs(prefsA)
		b := Agent{prefsB[i], Bnames[i], prefs}
		Universités = append(Universités, b)
	}

	// c := make(chan Request)

	// for i := 0; i < len(Bnames); i++ {
	// 	prefs := RandomPrefs(prefsA)
	// 	b := NewAgentDisposant(prefsB[i], Bnames[i], prefs,c)
	// 	b.Start()
	// 	Universités = append(Universités, b)
	// }

	// for i := 0; i < len(Anames); i++ {
	// 	prefs := RandomPrefs(prefsB)
	// 	a := NewAgentProposant(prefsA[i], Anames[i], prefs, c)
	// 	a.Start()
	// 	Elèves = append(Elèves, a)
	// }

	// time.Sleep(time.Minute)

	

	for _, a := range Elèves {
		fmt.Println(a)
	}

	for _, b := range Universités {
		fmt.Println(b)
	}

	fmt.Println(TTC(Elèves, Universités))

}
