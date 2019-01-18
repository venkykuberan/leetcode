package reverseint

import (
	"fmt"
	"math"
)

func reverseint(x int) int {
	var pop int
	var rev int
	rev = 0

	for x != 0 {
		pop = x % 10
		x = x / 10
		if rev > math.MaxInt32/10 {
			return 0
		}
		if rev < math.MinInt32/10 {
			return 0
		}
		rev = (rev * 10) + pop
	}
	// fmt.Println("Integer Max :", math.MaxInt32)
	fmt.Println("Reverse Int :", rev)
	return rev

}
