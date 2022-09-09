package h3_demo

import (
	"fmt"
	"github.com/uber/h3-go/v4"
)

func GetIndex() {
	latLng := h3.NewLatLng(21.016492249180757, 105.79668023976954)
	resolution := 9 // between 0 (biggest cell) and 15 (smallest cell)

	cell := h3.LatLngToCell(latLng, resolution)

	neighbors := h3.GridDisk(cell, 2)

	fmt.Println("len", len(neighbors))

	for _, item := range neighbors {
		fmt.Printf("%s", item)
		fmt.Println(" ")
	}
}
