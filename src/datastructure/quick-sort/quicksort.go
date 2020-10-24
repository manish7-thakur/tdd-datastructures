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

func partitionIndex(arr []int) (int, []int) {
	arrlen := len(arr)
	if arrlen < 2 {
		return 0, arr
	}
	pivot := arr[0]
	i := 0
	for j := 0; j < arrlen; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[0] = arr[0], arr[i]
	return i, arr
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i, _ := partitionIndex(arr)
	QuickSort(arr[:i])
	QuickSort(arr[i+1:])
	return arr
}
