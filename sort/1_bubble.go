package sort

//BubbleSort 冒泡排序
func BubbleSort(a []int) {
	n := len(a)
	if n < 2 {
		return
	}

	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}
}
