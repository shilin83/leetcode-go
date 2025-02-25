package problem0003

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	left, right, maxLen, hashTable := 0, 0, 0, make(map[byte]int, len(s))

	for right < length {
		// * 如果哈希表中存在当前元素，则移动滑动窗口左边界到当前元素的下一个位置
		char := s[right]
		if idx, ok := hashTable[char]; ok {
			left = max(left, idx+1)
		}

		// * 将当前元素及其索引存入哈希表
		// * 移动滑动窗口右边界
		// * 计算滑动窗口最大长度
		hashTable[char] = right
		right++
		maxLen = max(maxLen, right-left)
	}

	return maxLen
}
