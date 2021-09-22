package piscine

import "github.com/01-edu/z01"

func EightQueens() {
erreur := 0
for dame1 := '1'; dame1 <= '8'; dame1++ {
	for dame2 := '1'; dame2 <= '8'; dame2++ {
		if dame2 == dame1 || dame2 == dame1 -'1' || dame2 == dame1 +'1' {
			erreur = 1
		}
		for dame3 := '1'; dame3 <= '8'; dame3++ {
			if dame3 == dame1 ||dame3 == dame2 || dame3 == dame1 -'2' || dame3 == dame1 +'2' {
				erreur = 1
			}
			for dame4 := '1'; dame4 <= '8'; dame4++ {
				for dame5 := '1'; dame5 <= '8'; dame5++ {
					for dame6 := '1'; dame6 <= '8'; dame6++ {
						for dame7 := '1'; dame7 <= '8'; dame7++ {
							for dame8 := '1'; dame8 <= '8'; dame8++ {
								if dame2 == 0{
									if erreur == 0 {
										z01.PrintRune(rune(dame1))
										z01.PrintRune(rune(dame2))
										z01.PrintRune(rune(dame3))
										z01.PrintRune(rune(dame4))
										z01.PrintRune(rune(dame5))
										z01.PrintRune(rune(dame6))
										z01.PrintRune(rune(dame7))
										z01.PrintRune(rune(dame8))
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}