package algorithms

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			// dịch xuống 1 vị trí
			arr[j+1] = arr[j]
			j--
		}
		// khi tìm được vị trí mà arr[j] <= key thì đặt key ở vị trí j+1
		arr[j+1] = key
	}
}
