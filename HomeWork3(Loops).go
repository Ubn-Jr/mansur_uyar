package main

import "fmt"

func main() {
	var results [3]string = [3]string{"G", "B", "M"}
	var fenerbahce [6]string = [6]string{"G", "G", "B", "G", "B", "G"}
	var galatasaray [6]string = [6]string{"M", "G", "B", "M", "M", "B"}
	var besiktas [6]string = [6]string{"G", "M", "B", "M", "G", "M"}
	var trabzonBelediye [6]string = [6]string{"M", "M", "B", "G", "B", "B"}
	winning := "G"
	tie := "B"
	winningsOfFenerbahce := 0
	drawOfFenerbahce := 0
	winningsOfGalatasaray := 0
	drawOfGalatasaray := 0
	winningsOfBesiktas := 0
	drawOfBesiktas := 0
	winningsOfTrabzonBelediye := 0
	drawOfTrabzonBelediye := 0
	for i := 0; i < len(fenerbahce); i++ {
		var resultsOfFenerbahce string = string(fenerbahce[i])

		for j := 0; j < len(results); j++ {
			var resultsDeclaration string = string(results[j])
			if resultsDeclaration == resultsOfFenerbahce && resultsOfFenerbahce == winning {
				winningsOfFenerbahce++
			} else if resultsDeclaration == resultsOfFenerbahce && resultsOfFenerbahce == tie {
				drawOfFenerbahce++
			}
		}
	}
	for i := 0; i < len(galatasaray); i++ {
		var resultsOfGalatasaray string = string(galatasaray[i])

		for j := 0; j < len(results); j++ {
			var resultsDeclaration string = string(results[j])
			if resultsDeclaration == resultsOfGalatasaray && resultsOfGalatasaray == winning {
				winningsOfGalatasaray++
			} else if resultsDeclaration == resultsOfGalatasaray && resultsOfGalatasaray == tie {
				drawOfGalatasaray++
			}
		}
	}
	for i := 0; i < len(besiktas); i++ {
		var resultsOfBesiktas string = string(besiktas[i])

		for j := 0; j < len(results); j++ {
			var resultsDeclaration string = string(results[j])
			if resultsDeclaration == resultsOfBesiktas && resultsOfBesiktas == winning {
				winningsOfBesiktas++
			} else if resultsDeclaration == resultsOfBesiktas && resultsOfBesiktas == tie {
				drawOfBesiktas++
			}
		}
	}
	for i := 0; i < len(trabzonBelediye); i++ {
		var resultsOfTrabzonBelediye string = string(trabzonBelediye[i])

		for j := 0; j < len(results); j++ {
			var resultsDeclaration string = string(results[j])
			if resultsDeclaration == resultsOfTrabzonBelediye && resultsOfTrabzonBelediye == winning {
				winningsOfTrabzonBelediye++
			} else if resultsDeclaration == resultsOfTrabzonBelediye && resultsOfTrabzonBelediye == tie {
				drawOfTrabzonBelediye++
			}
		}
	}

	finalPointOfFenerbahce := winningsOfFenerbahce*3 + drawOfFenerbahce
	finalPointOfGalatasaray := winningsOfGalatasaray*3 + drawOfGalatasaray
	finalPointOfBesiktas := winningsOfBesiktas*3 + drawOfBesiktas
	finalPointOfTrabzonBelediye := winningsOfTrabzonBelediye*3 + drawOfTrabzonBelediye
	if finalPointOfFenerbahce > finalPointOfGalatasaray && finalPointOfFenerbahce > finalPointOfBesiktas && finalPointOfFenerbahce > finalPointOfTrabzonBelediye {
		fmt.Println("Our new champion is Fenerbahce.Congratulations to them!")
	} else if finalPointOfGalatasaray > finalPointOfFenerbahce && finalPointOfGalatasaray > finalPointOfBesiktas && finalPointOfGalatasaray > finalPointOfTrabzonBelediye {
		fmt.Println("Our new champion is Galatasaray.Congratulations to them!")
	} else if finalPointOfBesiktas > finalPointOfFenerbahce && finalPointOfBesiktas > finalPointOfGalatasaray && finalPointOfBesiktas > finalPointOfTrabzonBelediye {
		fmt.Println("Our new champion is Besiktas.Congratulations to them!")
	} else if finalPointOfTrabzonBelediye > finalPointOfFenerbahce && finalPointOfTrabzonBelediye > finalPointOfBesiktas && finalPointOfTrabzonBelediye > finalPointOfGalatasaray {
		fmt.Println("Our new champion is TrabzonBelediye.Congratulations to them!")
  } else {
    fmt.Println("Something must be went wrong!The results do not match!")
  }

}
