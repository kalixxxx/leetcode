package p4

// 7ms, 4.57MB
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	total := m + n
	needAver := false
	if total%2 == 0 {
		needAver = true
	}
	mid := total / 2

	return findMedian(nums1, nums2, needAver, mid)
}

func findMedian(nums1 []int, nums2 []int, needAver bool, mid int) float64 {
	var (
		mid1, mid2, idx, i, j int
	)

	val := 0
	for i < len(nums1) && j < len(nums2) {
		if i < len(nums1) && nums1[i] < nums2[j] {
			val = nums1[i]
			i++
		} else {
			val = nums2[j]
			j++
		}
		if needAver {
			if idx == mid-1 {
				mid1 = val
			}
		}
		if idx == mid {
			mid2 = val
		}
		idx++
	}

	for i >= len(nums1) && j < len(nums2) {
		val = nums2[j]
		j++
		if needAver {
			if idx == mid-1 {
				mid1 = val
			}
		}
		if idx == mid {
			mid2 = val
		}
		idx++
	}

	for j >= len(nums2) && i < len(nums1) {
		val = nums1[i]
		i++
		if needAver {
			if idx == mid-1 {
				mid1 = val
			}
		}
		if idx == mid {
			mid2 = val
		}
		idx++
	}
	if needAver {
		return float64(mid1+mid2) / 2
	} else {
		return float64(mid2)
	}
}
