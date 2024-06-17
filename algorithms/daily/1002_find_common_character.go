package daily

import (
	"fmt"
	"math"
)

func commonChars(words []string) []string {
	freqMap := make(map[string]int)
	for i := 0; i < len(words[0]); i++ {
		freqMap[string(words[0][i])]++
	}
	for i := 1; i < len(words); i++ {
		subFreqMap := make(map[string]int)
		item := words[i]
		for j := 0; j < len(item); j++ {
			subFreqMap[string(item[j])]++
		}

		for key, val := range freqMap {
			freq, found := subFreqMap[key]
			if !found {
				delete(freqMap, key)
				continue
			}

			freqMap[key] = int(math.Min(float64(freq), float64(val)))
		}
	}

	ans := make([]string, 0)
	for key, val := range freqMap {
		for val > 0 {
			ans = append(ans, key)
			val--
		}
	}
	return ans
}

func TestCommonChars() {
	words := []string{"cool", "lock", "cook"}
	fmt.Println(commonChars(words))
}
