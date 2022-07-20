package sorting

import "fmt"

type OSRMBatch struct {
	Origins      []string
	Destinations []string
}

func BatchMatrix() {
	origins := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1"}
	destinations := []string{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1", "o1", "p1"}

	lenDes := len(destinations)
	lenOrigin := len(origins)
	origin := make([]string, 0)
	des := make([]string, 0)
	maxBatch := 9
	batches := make([]OSRMBatch, 0)
	row := 0
	col := 0

	for col < lenDes && row < lenOrigin {
		origin = append(origin, origins[row])
		des = append(des, destinations[col])

		row++
		col++

		if len(origin)*len(des) >= maxBatch {
			batches = append(batches, OSRMBatch{
				Origins:      origin,
				Destinations: des,
			})
			origin = make([]string, 0)
			des = make([]string, 0)
			row = 0
			continue
		}
	}

	for _, batch := range batches {
		fmt.Println(batch)
	}
}
