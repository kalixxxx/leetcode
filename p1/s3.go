package p1

// 执行用时分布3ms， 消耗内存4.63MB
func solution3(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, num := range nums {
		if idx, ok := numsMap[target-num]; ok {
			return []int{idx, i}
		} else {
			numsMap[num] = i
		}
	}
	return nil
}
