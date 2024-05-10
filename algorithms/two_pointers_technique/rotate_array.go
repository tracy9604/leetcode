package two_pointers_technique

func RotateArray(nums []int, k int) {
	k %= len(nums)

	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])

}
func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
