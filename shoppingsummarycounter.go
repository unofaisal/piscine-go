package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	count := make(map[string]int)
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
		}
	}
	return count
}
