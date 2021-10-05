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
	}
	Bnames := [...]string{
		"Nathalie",
		"Annaïck",
		"Brigitte",
	}

	poolA := make([]Agent, 0, len(Anames))
	poolB := make([]Agent, 0, len(Bnames))

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
		poolA = append(poolA, a)
	}

	for i := 0; i < len(Bnames); i++ {
		prefs := RandomPrefs(prefsA)
		b := Agent{prefsB[i], Bnames[i], prefs}
		poolB = append(poolB, b)
	}

	for _, a := range poolA {
		fmt.Println(a)
	}

	for _, b := range poolB {
		fmt.Println(b)
	}

}
