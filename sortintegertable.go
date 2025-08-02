package piscine

func SortIntegerTable(table []int) {
	for range table {
		for i, c := range table {
			if i < len(table)-1 && c > table[i+1] {
				table[i], table[i+1] = table[i+1], table[i]
			}
		}
	}
}
