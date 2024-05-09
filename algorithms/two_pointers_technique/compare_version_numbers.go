package two_pointers_technique

func CompareVersionNumbers(version1 string, version2 string) int {
	v1Length, v2Length := len(version1), len(version2)
	i, j := 0, 0
	for i < v1Length || j < v2Length {
		num1, num2 := 0, 0
		for i < v1Length && version1[i] != '.' {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}
		for j < v2Length && version2[j] != '.' {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}

		if num1 == num2 {
			i++
			j++
			continue
		}

		if num1 < num2 {
			return -1
		}

		return 1
	}
	return 0
}
