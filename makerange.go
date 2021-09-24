package piscine

func MakeRange(min, max int) []int {
	answer := make([]int, max-min)
	for i := min; i < max; i++ {
		answer[i] = i + 1
	}
	return answer
}
