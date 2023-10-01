package leetcodearraystring

import "fmt"

const (
	test_nums1, test_nums2 = 1, 2
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	j, k := 1, 2
	x, y, z := 9, 65, 0

	fmt.Println("go output test: ", test_nums1, test_nums2, j, k)
	fmt.Println(x>>1, x/2, x%2)
	fmt.Println(y>>1, y>>2)
	fmt.Println(z >> 1)

	fmt.Println("test the meaning!")
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for idx, value := range strDict {
		fmt.Println(idx, value)
	}

	strArr := []string{"Jap", "Cha", "Aus"}
	for idx, val := range strArr {
		fmt.Println(idx, val)
	}

	i := 5
	for {
		fmt.Println("still in the loop", i)
		if i == 10 {
			break
		}
		i++
	}
}
