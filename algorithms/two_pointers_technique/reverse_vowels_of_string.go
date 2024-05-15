package two_pointers_technique

func ReversePowerOfAString(s string) string {
	sByte := []byte(s)
	i, j := 0, len(sByte)-1
	vowelsMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'O': true,
		'I': true,
		'U': true,
	}

	for i < j {
		_, foundI := vowelsMap[sByte[i]]
		_, foundJ := vowelsMap[sByte[j]]
		if foundI && foundJ {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			i++
			j--
			continue
		}
		if !foundI {
			i++
		}
		if !foundJ {
			j--
		}
	}
	return string(sByte)
}
