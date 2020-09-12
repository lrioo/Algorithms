package sort

//MergeSort 归并排序
func MergeSort(a []int) {
	n := len(a)
	if n < 2 {
		return
	}

	mergeSort(a, 0, n)
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}

	mid := (end - start) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)
	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	if a[mid] <= a[mid+1] {
		return
	}

	tmp := make([]int, end-start+1)
	i, j, k := start, mid+1, 0
	for ; i <= mid && j <= end; k++ {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
	}

	if i <= mid {
		for ; i <= mid; i++ {
			tmp[k] = a[i]
			k++
		}
	}
	if j <= end {
		for ; j <= end; j++ {
			tmp[k] = a[j]
			k++
		}
	}

	copy(a[start:end+1], tmp)
}
