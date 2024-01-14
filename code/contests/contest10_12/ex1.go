package contest10_12

import "math"

func FindTestedDevices(batteryPercentages []int) int {
	n := len(batteryPercentages)
	rs := 0
	for i := 0; i < n; i++ {
		if batteryPercentages[i] <= 0 {
			continue
		}
		for j := i + 1; j < n; j++ {
			batteryPercentages[j] = int(math.Max(0, float64(batteryPercentages[j]-1)))
		}
		rs++
	}
	return rs
}
