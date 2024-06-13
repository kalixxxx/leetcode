package p3

func main() {
	s2("acd")
}

func lengthOfLongestSubstring(s string) int {
	cMap := make(map[uint8]int)
	maxLen := 0
	length := 1

	if len(s) <= 1 {
		return len(s)
	}

	for i, _ := range s {
		cMap[s[i]] = i
		for j := i + 1; j < len(s); j++ {
			if _, ok := cMap[s[j]]; ok {
				maxLen = maxInt(length, maxLen)
				cMap = map[uint8]int{}
				length = 1
				break
			}
			cMap[s[j]] = j
			length += 1
		}
	}

	return maxInt(maxLen, length)
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
