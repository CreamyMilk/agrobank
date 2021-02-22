package sum_test

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/sum"
)

func ExampleInts() {
	s := sum.Ints(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1..5 is %v", s)
	// Output:
	// Sum of 1..5 is 15
}
