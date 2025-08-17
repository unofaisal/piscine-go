package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	// isAppend := true
	nS := ""
	counter := 0
	for i := 0; i < len(str); i++ {
		if str[i] > ' ' && str[i] <= '~' {
			nS += string(str[i])
			counter++
			if counter%5 == 0 {
				i++
				nS += " "
			}

			// fmt.Println(i+1, i+1%5)
		}
	}
	if len(nS) > 0 && nS[len(nS)-1] == ' ' {
		return nS[:len(nS)-1] + "\n"
	}
	return nS + "\n"
}
