package h3_demo

import (
	"fmt"
	"github.com/uber/h3-go/v4"
	"math"
)

func GetIndex() {
	latLng := h3.NewLatLng(21.016492249180757, 105.79668023976954)
	resolution := 9 // between 0 (biggest cell) and 15 (smallest cell)

	cell := h3.LatLngToCell(latLng, resolution)

	neighbors2 := h3.GridDisk(cell, 2)
	neighbors3 := h3.GridDisk(cell, 3)
	result := make([]h3.Cell, 0)

	fmt.Println("neighbors2", neighbors2)
	fmt.Println("neighbors3", neighbors3)

	neighbors2Map := make(map[h3.Cell]h3.Cell)
	for i := 1; i < len(neighbors2); i++ {
		if i == len(neighbors2)-1 {
			neighbors2Map[neighbors2[i]] = neighbors2[0]
			break
		}
		neighbors2Map[neighbors2[i]] = neighbors2Map[neighbors2[i+1]]
	}

	// k(0) = 1
	// k(1) = 6 + 1 = 7
	// k(2) = 6 * 2 + k(1) = 19
	// k(3) = 6 * 3 + k(2) = 37

	firstIndex := math.MinInt
	for i := len(neighbors2); i < len(neighbors3); i++ {
		if firstIndex >= 0 {
			break
		}
		k1Neighbors := h3.GridDisk(neighbors3[i], 1)
		k1NeighborsMap := make(map[h3.Cell]h3.Cell)
		for j := 1; j < len(k1Neighbors); j++ {
			if j == len(k1Neighbors)-1 {
				k1NeighborsMap[k1Neighbors[j]] = k1NeighborsMap[k1Neighbors[0]]
				break
			}
			k1NeighborsMap[k1Neighbors[j]] = k1NeighborsMap[k1Neighbors[j+1]]
		}

		for key, value := range k1NeighborsMap {
			if val, found := neighbors2Map[key]; found && value == val {
				result = append(result, neighbors3[i])
				firstIndex = i
				break
			}
		}
	}

	for i := 2; i <= 6; i++ {
		result = append(result, neighbors3[firstIndex+3])
	}

	fmt.Println("result", result)
}
