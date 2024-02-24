package arrays

func FindDisappearedNumbers(nums []int) []int {
	cnt := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}
	ans := make([]int, 0)
	for i := 1; i < len(cnt); i++ {
		if cnt[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
