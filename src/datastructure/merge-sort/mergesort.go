package merge_sort

func MergeSort(arr []int) []int {
	arrlen := len(arr)
	if arrlen < 2 {
		return arr
	}
	mid := arrlen / 2
	arr1 := MergeSort(arr[:mid])
	arr2 := MergeSort(arr[mid:])
	return merge(arr1, arr2)
}

func merge(first []int, second []int) []int {
	flen := len(first)
	if flen == 0 {
		return second
	}
	slen := len(second)
	if slen == 0 {
		return first
	}
	j := 0
	k := 0
	reslen := flen + slen
	res := make([]int, reslen)
	for i := 0; i < reslen; i++ {
		if j < flen && k < slen {
			if first[j] < second[k] {
				res[i] = first[j]
				j++
			} else {
				res[i] = second[k]
				k++
			}
		} else if j < flen {
			res[i] = first[j]
			j++
		} else {
			res[i] = second[k]
			k++
		}
	}
	return res
}
