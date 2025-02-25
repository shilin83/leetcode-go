package problem0004

import "math"

const (
	MaxInt = math.MaxInt32
	MinInt = math.MinInt32
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	// * low,high 表示 nums1 的左右边界
	low, high := 0, m
	for low <= high {
		// * i,j 表示 nums1,nums2 的分界线
		i := (low + high) >> 1
		j := ((m + n + 1) >> 1) - i

		// * 计算分界线两边的最大值和最小值
		var maxLeft1, minRight1, maxLeft2, minRight2 int
		if i == 0 {
			maxLeft1 = MinInt
		} else {
			maxLeft1 = nums1[i-1]
		}
		if i == m {
			minRight1 = MaxInt
		} else {
			minRight1 = nums1[i]
		}
		if j == 0 {
			maxLeft2 = MinInt
		} else {
			maxLeft2 = nums2[j-1]
		}
		if j == n {
			minRight2 = MaxInt
		} else {
			minRight2 = nums2[j]
		}

		// * nums1:  ……………… nums1[i-1] | nums1[i] ……………………
		// * nums2:  ……………… nums2[j-1] | nums2[j] ……………………
		if maxLeft1 < minRight2 && maxLeft2 < minRight1 {
			if ((m + n) & 1) == 1 {
				// * 数组长度为奇数，返回左半部分的最大值
				return float64(max(maxLeft1, maxLeft2))
			}
			// * 数组长度为偶数，返回中间两个数的平均值
			return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
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
