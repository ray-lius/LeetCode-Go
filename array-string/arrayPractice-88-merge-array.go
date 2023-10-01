/***

88. Merge Sorted Array
Solved
Easy
Topics
Companies
Hint
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.



Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

*****/

package leetcodearraystring

// import "fmt"

// const (
// 	test_nums1, test_nums2 = 1, 2
// )

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] >= nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

// func main() {
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	m := 3
// 	nums2 := []int{2, 5, 6}
// 	n := 3
// 	merge(nums1, m, nums2, n)
// 	fmt.Println(nums1)

// 	j, k := 1, 2
// 	x, y, z := 9, 65, 0

// 	fmt.Println("go output test: ", test_nums1, test_nums2, j, k)
// 	fmt.Println(x>>1, x/2, x%2)
// 	fmt.Println(y>>1, y>>2)
// 	fmt.Println(z >> 1)

// 	fmt.Println("test the meaning!")
// 	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
// 	for idx, value := range strDict {
// 		fmt.Println(idx, value)
// 	}

// 	strArr := []string{"Jap", "Cha", "Aus"}
// 	for idx, val := range strArr {
// 		fmt.Println(idx, val)
// 	}

// 	i := 5
// 	for {
// 		fmt.Println("still in the loop", i)
// 		if i == 10 {
// 			break
// 		}
// 		i++
// 	}
// }
