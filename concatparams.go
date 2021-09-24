package piscine

func ConcatParams(args []string) string {
	list := ""
	for i := 0; i < len(args)-1; i++ {
		list = list + args[i]
		list = list + "\n"
	}
	list = list + args[len(args)-1]
	return list
}
