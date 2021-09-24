package piscine

func MakeRange(min, max int) []int {
	pas_bon := []int{}
	answer := make([]int, max-min)
	if max-min < 0 {
		return pas_bon
	}
	for i := min; i < max; i++ {
		answer[i] = i + 1
	}
	return answer
}
