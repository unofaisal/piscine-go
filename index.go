package piscine

func Index(s string, toFind string) int {
	sL := len(s)
	toFindL := len(toFind)
	index := -1
	for x := 0; x <= sL-toFindL; x++ {
		if s[x:x+toFindL] == toFind {
			return x
		}
	}

	return index
}
