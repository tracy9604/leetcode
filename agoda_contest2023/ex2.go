package agoda_contest2023

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Ex2Input() {
	reader := bufio.NewReader(os.Stdin)
	idx := 0
	Y := 0
	dictinctiveness := make([]int, 0)
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		textSplits := strings.Split(text, " ")

		if idx == 0 {
			Y, _ = strconv.Atoi(textSplits[0])
		} else {
			for _, item := range textSplits {
				numItem, _ := strconv.Atoi(item)
				dictinctiveness = append(dictinctiveness, numItem)
			}
			fmt.Println(GetIslands(Y, dictinctiveness))
			break
		}
		idx++
	}
}

func IsSumPrimes(num int, primes []int) bool {
	for i := len(primes) - 1; i >= 0; i-- {
		if num-primes[i] > 0 {
			num = num - primes[i]
		} else if num-primes[i] == 0 {
			return true
		}
	}
	return false
}

func GetIslands(Y int, dictinctiveness []int) []int {
	primes := SeparateIntoPrimes(Y)
	sort.Ints(primes)
	primesMap := make(map[int]int)
	for _, item := range primes {
		primesMap[item]++
	}
	rs := make([]int, 0)
	for idx, item := range dictinctiveness {
		if IsSumPrimes(item, primes) {
			rs = append(rs, idx+1)
		}
	}
	return rs
}

func SeparateIntoPrimes(n int) []int {
	pfs := make([]int, 0)
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}
	return pfs
}
