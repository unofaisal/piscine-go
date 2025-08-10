package piscine

func ConcatParams(args []string) string {
	nString := ""
	for i, word := range args {
		nString += word
		if i < len(args)-1 {
			nString += "\n"
		}

	}
	return nString
}
