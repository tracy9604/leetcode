package contest_21_11

func FindBeautifulString(s string, k int) int {
	count := 0
	i := 0
	j := 0
	vowels := 0
	consonants := 0
	for i < len(s) {
		if string(s[j]) == "a" || string(s[j]) == "e" || string(s[j]) == "i" ||
			string(s[j]) == "o" || string(s[j]) == "u" {
			vowels++
		} else {
			consonants++
		}
		if vowels == consonants && (vowels*consonants)%k == 0 {
			count++
		}
		j++
		if j == len(s) {
			i++
			j = i
			vowels = 0
			consonants = 0
		}
	}
	//for i := 0; i < len(s); i++ {
	//	vowels := 0
	//	consonants := 0
	//	for j := i; j < len(s); j++ {
	//		if string(s[j]) == "a" || string(s[j]) == "e" || string(s[j]) == "i" ||
	//			string(s[j]) == "o" || string(s[j]) == "u" {
	//			vowels++
	//		} else {
	//			consonants++
	//		}
	//		if vowels == consonants && (vowels*consonants)%k == 0 {
	//			count++
	//		}
	//	}
	//
	//}
	return count
}
