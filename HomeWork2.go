package main

import "fmt"

func main() {
	pln := fmt.Println
	playersIn := 15
	playersOut := 8
	totalPlayers := playersIn + playersOut
	rules(totalPlayers, playersOut)
	pln(rules)
}

func rules(totalPlayers int, oneTeam int) {
	pln := fmt.Println
	if totalPlayers == 22 {
		pln("The match can be played without any problem!")
	} else if totalPlayers <= 22 && oneTeam == 7 {
		pln("If oneTeam lose any more player,they counted as a loser team!")
	} else if oneTeam <= 6 {
		pln("The match is over!")
	} else {
		pln("The match is can not be played!")
	}
}
