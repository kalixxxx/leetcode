package p3

func s2(s string) int {
	cMap := make(map[rune]int, len(s))
	maxLen := 0
	length := 0
	maxStep := 0
	for i, c := range s {
		if j, ok := cMap[c]; ok {
			maxLen = maxInt(length, maxLen)
			maxStep = maxInt(maxStep, j)
			length = i - maxStep
		} else {
			length += 1
		}
		cMap[c] = i
	}

	return maxInt(length, maxLen)
}
