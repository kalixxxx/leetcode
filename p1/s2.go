package p1

// 执行用时分布3ms， 消耗内存5.91MB
func solution2(nums []int, target int) []int {
	numsMap := make(map[int][]int)

	for i, num := range nums {
		if _, ok := numsMap[num]; ok {
			numsMap[num] = append(numsMap[num], i)
		} else {
			numsMap[num] = []int{i}
		}
	}

	for num, idxs := range numsMap {
		if indexs := numsMap[target-num]; len(indexs) > 0 {
			for _, index := range indexs {
				if idxs[0] != index {
					return []int{idxs[0], index}
				}
			}
		}
	}
	return nil
}
