package piscine
func ShoppingSummaryCounter(str string) map[string]int {
	sl := ""
	mp := make(map[string]int)

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			sl += string(str[i])
		} else if sl != "" {
			mp[sl]++
			sl = ""
		}
	}

	// Add the last word if any
	if sl != "" {
		mp[sl]++
	}

	return mp
}