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
