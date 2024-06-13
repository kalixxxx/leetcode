package main

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(nums, 2)
}

func removeElement(nums []int, val int) int {
	size := len(nums)
	count := 0
	for i := 0; i < size; i++ {
		if nums[i] == val {
			count++
			for j := i; j < size-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[size-1] = 0
		}
	}
	return count
}
