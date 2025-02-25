package problem0005

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// * start 和 end 分别表示最长回文子串的左右边界
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)

		// * 更新最长回文子串的左右边界
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - ((maxLen - 1) >> 1)
			end = i + (maxLen >> 1)
		}
	}

	return s[start : end+1]
}

// * 中心扩散法
func expandAroundCenter(s string, left int, right int) int {
	// * 当左右指针在有效范围内且对应元素相等时，继续向两边扩散
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
