package problem0003

func lengthOfLongestSubstring(s string) int {
	hashTable, maxLen, left := make(map[byte]int, len(s)), 0, 0

	for right, c := range s {
		char := byte(c)

		// * 如果哈希表中存在当前元素，则移动滑动窗口左边界到当前元素的下一个位置
		if idx, ok := hashTable[char]; ok && idx >= left {
			left = idx + 1
		}

		// * 将当前元素及其索引存入哈希表
		hashTable[char] = right

		// * 移动滑动窗口右边界
		// * 计算滑动窗口最大长度
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
