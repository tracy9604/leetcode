package main

import (
	"fmt"
	"leetcode/h3_demo"
)

func main() {
	lat := 21.016492249180757
	long := 105.79668023976954

	indexes := h3_demo.GetIndex(lat, long)

	rs := h3_demo.GetStrGeoFromH3Indexes(indexes)
	fmt.Println(rs)
}
