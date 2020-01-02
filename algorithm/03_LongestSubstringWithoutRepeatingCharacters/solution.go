package _03_LongestSubstringWithoutRepeatingCharacters

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	// T[i] -> 以 s[i-1] 结尾的最长子串的长度
	T := []int{0}
	// 用于判断当前字符是否在子串中出现
	M := map[int32]int{}
	max := 0
	for i, c := range s {
		idx, ok := M[c]
		v := 0
		if !ok || idx < i-T[i] {
			// 当前字符没有在以s[i-1]结尾的最长子串中出现
			v = T[i] + 1
		} else {
			// 当前字符在以s[i-1]结尾的最长子串中出现
			v = i - idx
		}
		T = append(T, v)
		M[c] = i
		if v > max {
			max = v
		}
	}
	return max
}
