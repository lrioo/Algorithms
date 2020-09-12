package list

//IsPalindrome1 判断是否为回文序列，栈
func IsPalindrome1(l *LinkedList) bool {
	if l == nil || &l.Head == nil || l.Head.Next == nil || l.len == 0 {
		return false
	}
	if l.len == 1 {
		return true
	}

	cur := &l.Head
	v := make([]interface{}, 0, l.len/2)
	for i := 0; i < l.len; i++ {
		cur = cur.Next
		if l.len%2 != 0 && i == l.len/2 {
			continue
		}

		if i < l.len/2 {
			v = append(v, cur.Val)
		} else {
			if v[l.len-1-i] != cur.Val {
				return false
			}
		}
	}

	return true
}

func IsPalindrome2(l *LinkedList) bool {
	if l == nil || &l.Head == nil || l.Head.Next == nil || l.len == 0 {
		return false
	}
	if l.len == 1 {
		return true
	}

	// 假设无长度，使用快慢指针寻找中间节点，同时转置前半部分链表
	var pre *Node
	slow, fast := l.Head.Next, l.Head.Next
	for fast != nil && fast.Next != nil {
		slow.Next, pre, slow, fast = pre, slow, slow.Next, fast.Next.Next
	}

	// 声明新的变量表示前后两部分的链表，以便于后续将链表恢复
	l1 := pre
	l2 := slow
	if fast != nil {
		// 若节点数为奇数，直接跳过中间节点
		l2 = l2.Next
	}

	flag := true
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			flag = false
			break
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	// 将链表恢复
	for pre != nil {
		pre.Next, slow, pre = slow, pre, pre.Next
	}
	l.Head.Next = slow

	return flag
}
