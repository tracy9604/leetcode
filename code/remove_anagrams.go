package code

func RemoveAnagrams(arr []string) []string {
	result := make([]string, 0)
	result = append(result, arr[0])
	for i := 1; i < len(arr); i++ {
		if !checkIsAnagrams(arr[i], arr[i-1]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func checkIsAnagrams(str1, str2 string) bool {
	str1Map := make(map[string]int)
	for _, letter := range str1 {
		str1Map[string(letter)]++
	}
	for _, letter := range str2 {
		if _, found := str1Map[string(letter)]; !found {
			return false
		}
		str1Map[string(letter)]--
	}

	for _, value := range str1Map {
		if value != 0 {
			return false
		}
	}
	return true
}
