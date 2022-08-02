package main

import "fmt"

func main() {
	pln := fmt.Println
	playersIn := 15
	playersOut := 7
	rules(playersIn, playersOut)
	pln(rules)
}

func rules(totalPlayers int, oneTeam int) {
	pln := fmt.Println
	if totalPlayers == 22 {
		pln("The match can be played without any problem!")
	} else if totalPlayers <= 22 && (oneTeam > 6 && oneTeam < 11) {
		pln("If oneTeam lose any more player,they counted as a loser team!")
	} else {
		pln("The match is over!")
	}
}
