package piscine

func Index(s string, toFind string) int {
	sentence := []rune(s)
	motif := []rune(toFind)
	counter_motif := 0
	counter := 0
	for i := 0; i < len(sentence)-len(motif); i++ {
		counter = 0
		for k := 0; k < len(motif); k++ {
			if motif[k] == sentence[i+k] {
				counter++
			}
		}
		if counter == len(motif) {
			return i
		}
	}
	if counter_motif == 0 {
		return -1
	} else {
		return counter_motif
	}
}
