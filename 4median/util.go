package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func earlyReturn(nums1 []int, nums2 []int, beginN1 int, endN1 int, beginN2 int, endN2 int) bool {
	middle := (endN1 + beginN1 + 1) / 2
	valueMiddle1 := float64(nums1[middle]) - 0.5
	deviceIndex1 := divide(nums1, valueMiddle1)
	deviceIndex2 := divide(nums2, valueMiddle1)

	sumLeft := deviceIndex1 + deviceIndex2
	sumRight := len(nums1) + len(nums2) - sumLeft

	valueMiddle2 := float64(nums1[middle]) + 0.5
	deviceIndex11 := divide(nums1, valueMiddle2)
	deviceIndex21 := divide(nums2, valueMiddle2)

	sumLeft2 := deviceIndex11 + deviceIndex21
	sumRight2 := len(nums1) + len(nums2) - sumLeft2

	if sumLeft+1 < sumRight && sumLeft2 > sumRight2+1 {
		return true
	}
	return false
}

// endN1 < len(nums1)
// endN2 < len(nums2)
func partion(nums1 []int, nums2 []int, beginN1 int, endN1 int, beginN2 int, endN2 int) float64 {
	middle := (endN1 + beginN1 + 1) / 2
	valueMiddle1 := float64(nums1[middle]) - 0.5
	deviceIndex1 := divide(nums1, valueMiddle1)
	deviceIndex2 := divide(nums2, valueMiddle1)

	sumLeft := deviceIndex1 + deviceIndex2
	sumRight := len(nums1) + len(nums2) - sumLeft

	if earlyReturn(nums1, nums2, beginN1, endN1, beginN2, endN2) {
		return float64(nums1[middle])
	}

	if sumLeft == sumRight+2 {
		var minN1 = 0
		var minN2 = 0
		if deviceIndex1 < 1 {
			minN1 = nums2[deviceIndex2-1]
		} else if deviceIndex2 < 1 {
			minN1 = nums1[deviceIndex1-1]
		} else {
			minN1 = max(nums1[deviceIndex1-1], nums2[deviceIndex2-1])
		}
		if deviceIndex1 > 0 && minN1 == nums1[deviceIndex1-1] {
			if deviceIndex1 < 2 {
				minN2 = nums2[deviceIndex2-1]
			} else if deviceIndex2 < 1 {
				minN2 = nums1[deviceIndex1-2]
			} else {
				minN2 = max(nums1[deviceIndex1-2], nums2[deviceIndex2-1])
			}
		} else {
			if deviceIndex1 < 1 {
				minN2 = nums2[deviceIndex2-2]
			} else if deviceIndex2 < 2 {
				minN2 = nums1[deviceIndex1-1]
			} else {
				minN2 = max(nums1[deviceIndex1-1], nums2[deviceIndex2-2])
			}
		}
		return float64(minN1+minN2) / 2.0
	} else if sumLeft+2 == sumRight {
		var maxN1, maxN2 = 0, 0
		maxN1 = min(nums1[deviceIndex1], nums2[deviceIndex2])
		if maxN1 == nums1[deviceIndex1] {
			if deviceIndex1+1 == len(nums1) {
				maxN2 = nums2[deviceIndex2]
			} else {
				maxN2 = min(nums1[deviceIndex1+1], nums2[deviceIndex2])
			}
		} else {
			if deviceIndex2+1 == len(nums2) {
				maxN2 = nums1[deviceIndex1]
			} else {
				maxN2 = min(nums1[deviceIndex1], nums2[deviceIndex2+1])
			}
		}
		return float64(maxN1+maxN2) / 2.0
	} else if sumLeft == sumRight+1 {
		if deviceIndex1 < 1 {
			return float64(nums2[deviceIndex2-1])
		} else if deviceIndex2 < 1 {
			return float64(nums1[deviceIndex1-1])
		}
		return float64(max(nums1[deviceIndex1-1], nums2[deviceIndex2-1]))
	} else if sumLeft+1 == sumRight {
		if deviceIndex1 >= len(nums1) {
			return float64(nums2[deviceIndex2])
		} else if deviceIndex2 >= len(nums2) {
			return float64(nums1[deviceIndex1])
		}
		return float64(min(nums1[deviceIndex1], nums2[deviceIndex2]))
	} else if sumLeft == sumRight {
		var maxN, minN = 0, 0
		if deviceIndex1 == len(nums1) {
			minN = nums2[deviceIndex2]
		} else if deviceIndex2 == len(nums2) {
			minN = nums1[deviceIndex1]
		} else {
			minN = min(nums1[deviceIndex1], nums2[deviceIndex2])
		}

		if deviceIndex1 < 1 {
			maxN = nums2[deviceIndex2-1]
		} else if deviceIndex2 < 1 {
			maxN = nums1[deviceIndex1-1]
		} else {
			maxN = max(nums1[deviceIndex1-1], nums2[deviceIndex2-1])
		}
		return float64(minN+maxN) / 2.0
	} else if sumLeft < sumRight {
		if deviceIndex2 <= endN2 {
			return partion(nums2, nums1, deviceIndex2, endN2, deviceIndex1, endN1)
		}
		return partion(nums1, nums2, deviceIndex1, endN1, deviceIndex2, endN2)
	} else {
		if beginN2 <= deviceIndex2-1 {
			return partion(nums2, nums1, beginN2,
				deviceIndex2-1, beginN1,
				deviceIndex1-1)
		}
		return partion(nums1, nums2, beginN1,
			deviceIndex1-1, beginN2,
			deviceIndex2-1)
	}
}

func divide(nums1 []int, middle float64) int {

	return findIndex(nums1, middle, 0, len(nums1)-1) + 1
}

// find the index_x
// if index_x = 0, all number > middle
// if index_x = -1, all number < middle
// left index_x - 1 < middle and right index_x > middle
// index_x begin 0
func findIndex(nums1 []int, middle float64, begin int, end int) int {
	numMiddle := (begin + end) / 2
	if float64(nums1[0]) > middle {
		return -1
	} else if begin == end {
		return begin
	} else if float64(nums1[numMiddle]) < middle && float64(nums1[numMiddle+1]) > middle {
		return numMiddle
	} else if float64(nums1[numMiddle+1]) < middle {
		return findIndex(nums1, middle, numMiddle+1, end)
	} else {
		return findIndex(nums1, middle, 0, numMiddle)
	}
}
