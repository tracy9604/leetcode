package sorting

func FindDifferenceString(s, t string) byte {
	map1 := make(map[string]int)
	for _, letter := range s {
		map1[string(letter)]++
	}

	for _, letter := range t {
		val, exist := map1[string(letter)]
		if !exist || val == 0 {
			return byte(letter)
		}

		map1[string(letter)]--
	}

	return 0
}
