package piscine

func BasicJoin(elems []string) string {
	sentence := ""
	for i := 0; i < len(elems); i++ {
		sentence = sentence + elems[i]
	}
	return sentence
}
