package array

import "sort"

func SortByBits(arr []int) []int {
	arrMap := make(map[int][]int)
	bitsMap := make(map[int]int)
	bits := make([]int, 0)
	result := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		count := Count1InBinary(arr[i])
		arrMap[count] = append(arrMap[count], arr[i])
		if _, found := bitsMap[count]; !found {
			bitsMap[count]++
			bits = append(bits, count)
		}
	}
	sort.Ints(bits)
	for i := 0; i < len(bits); i++ {
		if len(arrMap[bits[i]]) > 1 {
			sort.Ints(arrMap[bits[i]])
		}
		result = append(result, arrMap[bits[i]]...)
	}

	return result
}

func Count1InBinary(n int) int {
	count := 0
	for n > 0 {
		mid := n / 2
		if n-mid*2 == 1 {
			count++
		}
		n = mid
	}
	return count
}

func Count1InBinary2(n int) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += (n >> i) & 1
	}
	return count
}
