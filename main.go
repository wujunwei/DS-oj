package main

import (
	"ds/data_struct"
	"fmt"
)

func main() {
	sl := data_struct.NewSegment(0, 9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(sl.Get(0, 9))
	fmt.Println(sl.Get(1, 4))
	sl.Update(0, 4, 3)
	fmt.Println(sl.Get(1, 4))
}
