package sorting

func CheckValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tmpMap := make(map[string]int)

	for _, char := range s {
		tmpMap[string(char)]++
	}

	for _, char := range t {
		_, check := tmpMap[string(char)]
		if !check {
			return false
		}
		tmpMap[string(char)]--
	}

	for _, value := range tmpMap {
		if value != 0 {
			return false
		}
	}

	return true
}

func CheckValidAnagramV2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m [26]int
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
