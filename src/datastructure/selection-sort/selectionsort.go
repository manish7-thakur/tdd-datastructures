package selection_sort

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minidx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minidx] {
				minidx = j
			}
		}
		arr[i], arr[minidx] = arr[minidx], arr[i]
	}
	return arr
}
