package quick_sort

func partition(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	j := len(arr) - 1
	i := 0
	pivot := arr[0]
	for ; i < j; {
		for ; i < len(arr) && arr[i] <= pivot; i++ {
		}
		for ; j > 0 && arr[j] >= pivot; j-- {
		}
		if i < j {
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[0], &arr[j])

	return arr
}

func swap(i *int, j *int) {
	temp := *i
	*i = *j
	*j = temp
}
