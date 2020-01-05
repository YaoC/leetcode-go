package _1295_FindNumbersWithEvenNumberOfDigits

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
func findNumbers(nums []int) int {
	ret := 0
	for _, num := range nums {
		if withEvenDigits(num) {
			ret++
		}
	}
	return ret
}

func withEvenDigits(num int) bool {
	if num == 0 {
		return false
	}
	ret := true
	for num != 0 {
		ret = !ret
		num /= 10
	}
	return ret
}
