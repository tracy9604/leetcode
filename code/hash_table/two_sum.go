package hash_table

func FindTwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for idx, num := range nums {
		if pos, found := hashMap[target-num]; found {
			return []int{pos, idx}
		}
		hashMap[num] = idx

	}
	return []int{}
}
