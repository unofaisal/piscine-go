package piscine

func LoafOfBread(str string) string {
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
			if i != len(str)-1 && counter%5 == 0 {
				i++
				nS += " "
			}

			// fmt.Println(i+1, i+1%5)
		}
	}
	return nS + "\n"
}
