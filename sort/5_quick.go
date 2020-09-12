package sort

//QuickSort 快速排序
func QuickSort(a []int) {
	n := len(a)
	if n < 2 {
		return
	}

	quickSort(a, 0, n)
}

func quickSort(a []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(a, start, end)
	quickSort(a, start, i-1)
	quickSort(a, i+1, end)
}

func partition(a []int, start, end int) int {
	pivot := a[end]

	i := start
	for j := start; j < end; j++ {
		if a[j] < pivot {
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	a[i], a[end] = a[end], a[i]

	return i
}
