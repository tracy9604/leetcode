package code

func HammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num & 1)
		num = num >> 1
	}
	return count
}
