/*
189. Rotate Array
Solved
Medium
Topics
Companies
Hint
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.



Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105

*/

package arraystring

import "fmt"

func rotateArray(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v :=  range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func rotateArray2(nums []int, k int) {
	reverseArray(nums)
	reverseArray(nums[:k])
	reverseArray(nums[k:])
}

func reverseArray(nums []int){
	for i, n := 0, len(nums); i<n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func TestRotateArray() {
	nums1 := []int{1,2,3,4,5,6,7}
	k1 := 3
	// rotateArray(nums1, k1)

	rotateArray2(nums1, k1)
	fmt.Println(nums1)
}