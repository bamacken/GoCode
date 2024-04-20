package main

import (
	easy "GoCode/algorithm/easy"
	hard "GoCode/algorithm/hard"
	medium "GoCode/algorithm/medium"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(medium.MyPow(2, 10))
	fmt.Print(medium.LongestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
	fmt.Println(medium.Reverse(1534236469))
	fmt.Println(easy.IsValid("{[]}"))
	fmt.Println(easy.FindBuildings([]int{1, 3, 2, 4}))

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	hard.FindMedianSortedArrays(nums1, nums2)
}

// mergeArray(nums1, 3, nums2, 3) is the solution to the following
// question: https://leetcode.com/problems/merge-sorted-array/
func mergeArray(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
