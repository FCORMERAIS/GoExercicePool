package piscine

func AppendRange(min, max int) []int {
	var answer []int
	if min > max {
		return answer
	} else if min == max {
		answer = append(answer, min)
		return answer
	} else {
		for i := max - min; i <= max; i++ {
			answer = append(answer, i)
		}
		return answer
	}
}
