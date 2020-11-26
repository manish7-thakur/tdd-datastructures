package selection_procedure

func Partition(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	pos := 0
	pivot := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			pos++
		}
	}
	return pos
}
