package two_pointers_technique

func IsSubsequence(s string, t string) bool {
	sByte := []byte(s)
	tByte := []byte(t)

	i, j := 0, 0

	for i < len(sByte) && j < len(tByte) {
		if sByte[i] == tByte[j] {
			i++
			j++
			continue
		}
		j++
	}

	return i >= len(sByte)-1

}
