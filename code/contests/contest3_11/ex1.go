package contest3_11

func FindPeaks(mountain []int) []int {
	rs := make([]int, 0)
	for i := 0; i < len(mountain); i++ {
		if i == 0 || i == len(mountain)-1 {
			continue
		}

		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			rs = append(rs, i)
		}
	}
	return rs
}
