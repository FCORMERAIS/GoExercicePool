package piscine

func MakeRange(min, max int) []int {
	pas_bon := []int{}
	if max-min < 0 {
		return pas_bon
	}
	answer := make([]int, max-min)
	for i := min; i < max; i++ {
		answer[i-min] = i
	}
	return answer
}
