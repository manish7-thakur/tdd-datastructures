package bubble_sort

func BubbleSort(arr []int) []int {
	arrlen := len(arr)
	for i := arrlen - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
