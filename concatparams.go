package piscine

func ConcatParams(args []string) string {
	list := ""
	for i := 0; i < len(args); i++ {
		list = list + args[i]
		list = list + "\n"
	}
	list = list + "\n"
	return list
}
