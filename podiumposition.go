package piscine

func PodiumPosition(podium [][]string) [][]string {
	for range podium {
		for i, no := range podium {
			if i+1 < len(podium) && no[0][0] > podium[i+1][0][0] {
				podium[i], podium[i+1] = podium[i+1], podium[i]
			}
		}
	}
	return podium
}
