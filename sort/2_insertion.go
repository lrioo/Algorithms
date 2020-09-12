package sort

//InsertionSort 插入排序
func InsertionSort(a []int) {
	n := len(a)
	if n < 2 {
		return
	}

	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] <= v {
				break
			}

			a[j+1] = a[j]
		}
		a[j+1] = v
	}
}
