package contest19_11

func SwapBalls(s string) int64 {
	arr := make([]string, 0)
	for i := 0; i < len(s); i++ {
		arr = append(arr, string(s[i]))
	}
	count := 0
	lastOne := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == "1" && arr[i+1] == "0" {
			j := i
			for j > lastOne {
				count += 1
				j--
			}
			arr[j], arr[i+1] = arr[i+1], arr[j]
			lastOne = j
		}
	}
	return int64(count)
}
