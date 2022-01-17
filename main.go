// Author: Mitch Allen

package main

import (
	"fmt"

	"github.com/mitchallen/pick"
)

func main() {
	list1 := []int{20, 10, 30}
	fmt.Println(pick.OneInt(list1))

	i1 := pick.RandomIndex(len(list1))
	fmt.Println(i1)

	list2 := []string{"beta", "alpha", "gamma"}
	fmt.Println(pick.OneString(list2))

	weights := []float32{0.50, 0.30, 0.20}
	for i := 0; i < 5; i++ {
		m := make(map[int]int)
		for j := 0; j < 1000; j++ {
			m[pick.WeightedIndex(weights)]++
		}
		fmt.Println(m)
	}

}
