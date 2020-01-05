package _04_MedianOfTwoSortedArrays

// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if n < m {
		return findMedianSortedArrays(nums2, nums1)
	}

	var (
		l1, l2, r1, r2, c1, c2 int
		lo, hi                 = 0, 2 * m
	)
	for lo <= hi {
		c1 = lo + (hi-lo)/2
		l1 = getLeft(&nums1, c1)
		r1 = getRight(&nums1, c1)
		c2 = m + n - c1
		l2 = getLeft(&nums2, c2)
		r2 = getRight(&nums2, c2)

		if l1 > r2 {
			hi = c1 - 1
		} else if l2 > r1 {
			lo = c1 + 1
		} else {
			break
		}
	}
	return float64(max(l1, l2)+min(r1, r2)) / 2.0
}

func getLeft(nums *[]int, c int) int {
	if c == 0 {
		return -1 << 31
	}
	return (*nums)[(c-1)/2]
}

func getRight(nums *[]int, c int) int {
	if c == 2*len(*nums) {
		return 1<<31 - 1
	}
	return (*nums)[c/2]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
