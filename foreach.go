package piscine

func ForEach(f func(int), a []int) {
	for _, in := range a {
		f(in)
	}
}
