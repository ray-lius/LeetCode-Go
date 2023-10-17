package Test01

import "fmt"

type fn func(int) int

func square(num int) int {
	return num*num
}

func mapper(f fn, alist []int) []int {
	// make type, kength, and  capacity
	a :=  make([]int, len(alist), len(alist))
	for index, val := range alist {
		a[index] = f(val)
	}
	return a
}
func TestCase() {
	alist := []int{4,5,6,7}
	result := mapper(square, alist)
	fmt.Println(result)
}
func main() {
	alist := []int{4,5,6,7}
	result := mapper(square, alist)
	fmt.Println(result)
	
}
