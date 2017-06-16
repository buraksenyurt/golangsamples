package operations

import (
	"math"
)

func CircleSpace(r float32) float32 {
	return math.Pi * (r * r)
}

func Sum(numbers ...int) int {
	var total int = 0
	for _, n := range numbers {
		total += n
	}
	return total
}
