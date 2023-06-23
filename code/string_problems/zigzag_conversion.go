package string_problems

func FindZigZagConversion(s string, numRows int) string {
	direction := "down"
	row := 0
	arr := make([][]string, numRows)
	for idx, _ := range arr {
		arr[idx] = make([]string, 0)
	}
	for i := 0; i < len(s); i++ {
		arr[row] = append(arr[row], string(s[i]))
		if row == numRows-1 {
			direction = "up"
			row = 0
			continue
		}
		row++
	}
	return direction
}
