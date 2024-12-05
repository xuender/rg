package str_test

import (
	"fmt"

	"github.com/xuender/ramdago/str"
)

func ExampleMatch() {
	fmt.Println(str.Match("[a-z]a")("bananas"))
	fmt.Println(str.Match("a")("b"))

	// Output:
	// [ba na na]
	// []
}
