package main

import (
	"fmt"
	"math"
)

// 167. Two Sum II - Input array is sorted
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {

		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			break
		}
	}

	return []int{i + 1, j + 1}
}

// 633. 平方数之和
func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))

	for i <= j {
		x := i*i + j*j
		if x < c {
			i++
		} else if x > c {
			j--
		} else {
			return true
		}
	}

	return false
}

//  88. 合并两个有序数组
//
// 归并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j := m-1, n-1
	pos := m + n - 1

	for i >= 0 && j >= 0 {

		if nums1[i] >= nums2[j] {
			nums1[pos] = nums1[i]
			i--
		} else {
			nums1[pos] = nums2[j]
			j--
		}
		pos--
	}

	for j >= 0 {
		nums1[pos] = nums2[j]
		pos--
		j--
	}

}

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	n := len(s)
	cntT := make(map[byte]int)
	freq := make(map[byte]int)
	for i := range t {
		cntT[t[i]]++
	}

	isContainT := func() bool {
		for k, v := range cntT {
			if freq[k] < v {
				return false
			}
		}

		return true
	}

	ansL, ansR := -1, math.MaxInt32
	for i, j := 0, 0; j < n; j++ {
		if cntT[s[j]] > 0 {
			freq[s[j]] += 1
		}

		fmt.Println(j, isContainT())

		// 如果当前窗口包含了t
		for isContainT() && i <= j {
			if j-i+1 < ansR-ansL {
				// 更新结果
				ansL, ansR = i, j+1
			}

			if cntT[s[i]] > 0 {
				freq[s[i]]--
			}
			i++
		}
	}

	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR]
}

//  142. 环形链表 II
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表解法
// func detectCycle(head *ListNode) *ListNode {
// 	visit := make(map[*ListNode]bool)
//     p := head

//     for p != nil {
//         if visit[p] {
//             return p
//         }

//         visit[p] = true
//         p = p.Next
//     }

//	    return nil
//	}
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	flag := true

	for fast != slow || flag {
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next
		flag = false
	}

	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

// 680. 验证回文串 II
func validPalindrome(s string) bool {
	check := func(str string) bool {
		i, j := 0, len(str)-1
		for i < j {
			if str[i] != str[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return check(s[i:j]) || check(s[i+1:j+1])
		}
		i++
		j--
	}

	return true
}

// 524. 通过删除字母匹配到字典里最长单词
func findLongestWord(s string, dictionary []string) string {

	match := func(str1, str2 string) bool {
		i, j := 0, 0
		for i < len(str1) && j < len(str2) {
			if str1[i] == str2[j] {
				j++
			}
			i++
		}
		return j == len(str2)
	}

	ans := ""
	for _, word := range dictionary {
		if !match(s, word) {
			continue
		}

		if len(word) > len(ans) {
			ans = word
		}
		if len(word) == len(ans) && word < ans {
			ans = word
		}
	}

	return ans
}

// 340. 至多包含 K 个不同字符的最长子串
// 给定一个字符串 s 和一个整数 k，请你找出 至多包含 k 个不同字符的最长子串 的长度。

// 示例：
// 输入: s = "eceba", k = 2
// 输出: 3
// 解释: 最长子串是 "ece"，包含 'e' 和 'c' 两种字符。

// 输入: s = "aa", k = 1
// 输出: 2

// ✅ 解题思路：滑动窗口（Sliding Window）
// 这是典型的 “最多 K 个不同元素” 的滑动窗口问题，可用 双指针 + 哈希表 解决。
// 核心思想：
// 用两个指针 left 和 right 维护一个窗口 [left, right]
// 用哈希表 charCount 记录窗口中每个字符的出现次数
// 扩展 right：不断加入新字符
// 当窗口中不同字符数 > k 时，收缩 left 直到满足条件
// 每次满足条件时，更新最大长度
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	charCount := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		// 扩展右边界
		charCount[s[right]]++

		// 收缩左边界，直到不同字符数 <= k
		for len(charCount) > k {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		// 更新最大长度（此时窗口合法）
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

// fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2)) // 3
// fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1))    // 2
// fmt.Println(lengthOfLongestSubstringKDistinct("a", 2))     // 1
// fmt.Println(lengthOfLongestSubstringKDistinct("", 1))      // 0
// fmt.Println(lengthOfLongestSubstringKDistinct("abc", 0))   // 0
