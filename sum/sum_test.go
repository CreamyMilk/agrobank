package sum_test

import (
	"testing"

	"github.com/CreamyMilk/agrobank/sum"
)

func TestInts(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"1 to 3", []int{1, 2, 3}, 6},
		{"Nil Array", nil, 0},
		{"One and negative ", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := sum.Ints(tc.numbers...)
			if s != tc.sum {
				t.Errorf("Sum of %v should be %v ; but got %v", tc.name, tc.sum, s)
			}
		})
	}

}
