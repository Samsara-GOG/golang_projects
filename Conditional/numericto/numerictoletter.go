package main

import "fmt"

// Create a numeric grade to letter grade conditional (or vice versa).

func main() {
	noteCollege := "F"
	fmt.Printf("Votre note du college (6e à la 3e) est : %v. ", noteCollege)

	switch noteCollege {
	case "A+":
		fmt.Printf("Vous avez une note entre 16 et 20 en France.")
	case "A":
		fmt.Printf("Vous avez la note de 15/20 en France.")
	case "A-":
		fmt.Printf("Vous avez la note de 14/20 en France.")
	case "B+":
		fmt.Printf("Vous avez la note de 13/20 en France.")
	case "B":
		fmt.Printf("Vous avez la note de 12/20 en France.")
	case "B-":
		fmt.Printf("Vous avez la note de 11/20 en France.")
	case "C+":
		fmt.Printf("Vous avez la note de 10/20 en France.")
	case "C":
		fmt.Printf("Vous avez la note de 9/20 en France.")
	case "C-":
		fmt.Printf("Vous avez la note de 8/20 en France.")
	case "D+":
		fmt.Printf("Vous avez la note de 7/20 en France.")
	case "D":
		fmt.Printf("Vous avez la note de 6/20 en France.")
	case "D-":
		fmt.Printf("Vous avez la note de 5/20 en France.")
	//	case "F":
	default:
		fmt.Printf("Vous avez la note de 4/20 en France.")

		//   Source :  https://webcache.googleusercontent.com/search?q=cache:ZnPxlPLDUzcJ:https://miami.consulfrance.org/IMG/pdf/Equivalence_de_notation_entre_le_systeme_Francais_et_le_systeme_Americain.pdf+&cd=3&hl=fr&ct=clnk&gl=fr&client=firefox-b-d
		/*
					   6e à la 3e

					   A+ 16-20
					   A 15
					   A- 14
					   B+ 13
					   B 12
					   B- 11
					   C+ 10
					   C 09
					   C- 08
					   D+ 07
					   D 06
					   D- 05
					   F 04
			             --------------------------
					   2nd à la terminale

					   A+ 15-20
					   A 14
					   A- 13
					   B+ 12
					   B 11
					   B- 10
					   C+ 09
					   C 08
					   C- 07
					   D+ 06
					   D 05
					   D- 04
					   F 03
		*/

	}
}
