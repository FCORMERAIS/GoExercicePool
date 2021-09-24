package piscine

func MakeRange(min, max int) []int {
	var pas_bon []int
	if min>max {
		return pas_bon
	}
	answer := make([]int, max-min)
	for i := min; i < max; i++ {
		answer[i-min] = i
	}
	return answer
}
