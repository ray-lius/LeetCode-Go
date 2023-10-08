/*
167. Two Sum II - Input Array Is Sorted
Solved
Medium
Topics
Companies
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.



Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

*/

package TwoPoint

import (
	"fmt"
	"sort"
)


func twoSum(nums []int, target int) []int {
	nums_len := len(nums)
	left, right := 0, nums_len-1
	for left<right {
		sum := nums[left] + nums[right]
		
		if sum == target {
			return []int{left+1, right+1} //+1 becasue they require output indice of position
		} else if sum < target {
			left ++
		} else {
			right --
		}
	}
	return nil
}

/*

15. 3Sum
Solved
Medium
Topics
Companies
Hint
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

 

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Solution:

we csn use the two sum solution toget thto get the solution
*/


func threeSum1(nums []int, target int) [][]int {

	sort.Ints(nums)
	nums_len := len(nums)
	result := [][]int{}
	// as we at 3least need 3 nums
	for i:=0; i< nums_len-2; i++ {
		
		two_sum_result := twoSum(nums[i+1:], target-nums[i])
		fmt.Println("in the loop",nums[i+1:], target-nums[i], result, two_sum_result)
		if two_sum_result != nil {
			result = append(result, []int{nums[i], nums[two_sum_result[0]+i], nums[two_sum_result[1]+i]})
		}
	}

	fmt.Println("result without deduplicate",result)
	// remove the same array
	result = removeDuplicateArray(result)

	return result

}

func removeDuplicateArray(nums [][]int) [][]int {
	res := [][]int{}
	for i:=0; i<len(nums); i++ {

		
		if len(res) == 0 {
			res = append(res, nums[i])
		}else{
			isHaveSameArray := false
			for j:=0; j<len(res); j++ {
				if isSameArray(nums[i], res[j]) {
					isHaveSameArray = true
					break
				}
			}
			if !isHaveSameArray {
				res = append(res, nums[i])
			}
		}
		
	}
	return res
}

func isSameArray(a []int, b[]int) bool {
	a_len, b_len := len(a), len(b)
	if a_len != b_len {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for idx, val := range a {
		if val != b[idx] {
			return false
		}
	}
	return true
}


// best solution
func threeSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	nums_len, start, index, end, sum := len(nums), 0, 0, 0, 0

	for index=1; index<nums_len-1; index++ {
		start, end = 0, nums_len-1
		if index>1 && nums[index] == nums[index-1] {
			start = index-1
		}

		for start <index && end>index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < nums_len -1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum = nums[start] + nums[index] + nums[end]
			if sum == target {
				res = append(res, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			}else if (sum > target) {
				end--
			}else {
				start++
			}
		}
		
	}
	return res
}

func TestCase(){
	nums := []int{-1,0,1,2,-1,-4}

	fmt.Println(threeSum1(nums, 0))
	fmt.Println(threeSum2(nums, 0))
}