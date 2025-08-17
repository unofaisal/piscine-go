package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	count := make(map[string]int)
	if str == "" {
		count[""] = 1
		return count
	}

	item := ""
	for i, ch := range str {
		if ch != ' ' {
			item += string(ch)
			if i == len(str)-1 {
				count[item]++
			}
		} else {
			count[item]++
			item = ""
			if i == len(str)-1 {
				count[""]++
			}
		}
	}
	return count
}
