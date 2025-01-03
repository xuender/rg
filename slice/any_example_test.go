package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleAny() {
	greaterThan0 := func(num int) bool { return 0 > num }
	greaterThan2 := func(num int) bool { return 2 > num }

	fmt.Println(slice.Any([]int{1, 2}, greaterThan0))
	fmt.Println(slice.Any([]int{1, 2}, greaterThan2))

	// Output:
	// false
	// true
}
