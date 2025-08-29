package problem0004

import "github.com/shilin83/leetcode-go/src/helpers"

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	// * 确保 nums1 的长度小于等于 nums2 的长度
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	// * low, high 分别表示 nums1 的左右边界
	low, high := 0, m

	// * nums1:  ……………… nums1[i-1] | nums1[i] ……………………
	// * nums2:  ……………… nums2[j-1] | nums2[j] ……………………
	for low <= high {
		// * i, j 分别表示 nums1 和 nums2 分割点
		i := (low + high) / 2
		j := (m+n+1)/2 - i

		// * 获取分界线左右两侧的最大值和最小值
		var maxLeft1, minRight1, maxLeft2, minRight2 int
		if 0 == i {
			maxLeft1 = helpers.MinInt32
		} else {
			maxLeft1 = nums1[i-1]
		}

		if m == i {
			minRight1 = helpers.MaxInt32
		} else {
			minRight1 = nums1[i]
		}

		if 0 == j {
			maxLeft2 = helpers.MinInt32
		} else {
			maxLeft2 = nums2[j-1]
		}

		if n == j {
			minRight2 = helpers.MaxInt32
		} else {
			minRight2 = nums2[j]
		}

		// * 检查分割是否正确
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			// * 偶数个元素，返回中间两个数的平均值
			if 0 == (m+n)%2 {
				return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2
			}

			// * 奇数个元素，返回左半部分的最大值
			return float64(max(maxLeft1, maxLeft2))
		} else if maxLeft1 > minRight2 {
			// * nums1 中的分界线划多了，要向左边移动
			high = i - 1
		} else {
			// * nums1 中的分界线划少了，要向右边移动
			low = i + 1
		}
	}

	return 0.0
}
