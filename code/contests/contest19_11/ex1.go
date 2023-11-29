package contest19_11

import "math"

func MinimizeString(s1 string, s2 string, s3 string) int {
	if s1 == s2 && s2 == s3 {
		return 0
	}

	minLen := int(math.Min(float64(len(s1)), float64(len(s2))))
	minLen = int(math.Min(float64(minLen), float64(len(s3))))

	count := 0
	for i := 0; i < minLen; i++ {
		if s1[i] == s2[i] && s1[i] == s3[i] {
			count++
			continue
		}
		break

	}

	if count == 0 {
		return -1
	}

	return len(s1) + len(s2) + len(s3) - 3*count
}
