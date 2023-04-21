package agoda_contest2023

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Ex1Input() {
	reader := bufio.NewReader(os.Stdin)

	N := 0
	idx := 0
	nightlyRate := make([]uint64, 0)
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if idx == 0 {
			N, _ = strconv.Atoi(text)
		} else {
			num, _ := strconv.Atoi(text)
			nightlyRate = append(nightlyRate, uint64(num))
		}

		if idx == N {
			fmt.Println(CountRate(nightlyRate, 0))
			break
		}

		idx++
	}
}

func CountRate(nightRate []uint64, idx int) uint64 {
	if idx >= len(nightRate) {
		return 0
	}
	return uint64(math.Max(float64(nightRate[idx]+CountRate(nightRate, idx+2)), float64(CountRate(nightRate, idx+1))))
}
