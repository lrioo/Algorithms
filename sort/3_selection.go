package sort

//SelectionSort 选择排序
func SelectionSort(a []int) {
	n := len(a)
	if n < 2 {
		return
	}

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
