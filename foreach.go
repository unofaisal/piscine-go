package piscine

func ForEach(f func(int), a []int) {
	for ind, in := range a {
		a[ind] = f(in)
	}
}
