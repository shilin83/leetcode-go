package problem0001

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int, len(nums))

	for i, num := range nums {
		// * 如果哈希表中存在差值，则返回差值与当前元素的索引
		if j, ok := hashTable[target-num]; ok {
			return []int{j, i}
		}

		// * 将当前元素及其索引存入哈希表
		hashTable[num] = i
	}

	return []int{}
}
