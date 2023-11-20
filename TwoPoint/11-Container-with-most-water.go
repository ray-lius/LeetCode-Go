/*

11. Container With Most Water

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:

Input: height = [1,1]
Output: 1

*/

package TwoPoint

import "fmt"

// second practice
func Sum3MaxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	for left < right {
		area := min(height[left], height[right]) * (right-left)
		if area > maxArea {
			maxArea = area
			left ++
		}else{
			right --
		}
	}
	return maxArea
}



func maxArea(height []int) int {
	maxArea, left, right := 0, 0, len(height)-1

	for left<right {
		area := min(height[left],height[right]) * (right-left)
		if area>maxArea {
			maxArea = area
			left++
		}else{
			right--
		}

	}
	return maxArea
}

func TestCase2(){
	test1 := []int{1,8,6,2,5,4,8,3,7} 
	test2 := []int{1,1}

	testCases := [][]int{test1, test2}

	for _, val := range testCases {
		fmt.Println(maxArea(val))
	}
	
}