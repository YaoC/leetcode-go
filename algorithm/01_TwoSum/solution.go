package _01_TwoSum

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return []int{j, i}
		}
		m[target-v] = i
	}
	return []int{}
}
