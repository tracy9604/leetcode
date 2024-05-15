package two_pointers_technique

import "strconv"

func StringCompression(chars []byte) int {
	i, j, k := 0, 0, 0

	for i < len(chars) && j < len(chars) {
		chars[k] = chars[i]
		k++
		for j < len(chars) && chars[i] == chars[j] {
			j++
		}

		if j-i > 1 {
			runLength := strconv.Itoa(j - i)
			for _, item := range runLength {
				itemByte := byte(item)
				chars[k] = itemByte
				k++
			}
		}
		i = j
	}

	return k

}
