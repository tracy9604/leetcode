package contest_14_04

import "math"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	maxLoop := int(math.Sqrt(float64(n)))
	for i := 2; i < maxLoop+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func DistancePrime(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		if !isPrime(nums[left]) {
			left++
			continue
		}
		if !isPrime(nums[right]) {
			right--
			continue
		}

		if isPrime(nums[left]) && isPrime(nums[right]) {
			return right - left
		}
	}
	return 0
}
