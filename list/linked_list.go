package list

import (
	"fmt"
)

//LinkedList 单链表
type LinkedList struct {
	Head Node
	len  int
}

// String 打印链表
func (l *LinkedList) String() string {
	if l == nil || l.Head.Next == nil {
		return "<nil>"
	}

	s := "head"
	cur, length := l.Head.Next, l.len
	for cur != nil && length >= 0 {
		s = fmt.Sprintf("%s->%+v", s, cur.Val)
		cur = cur.Next
		length--
	}
	return s
}

// Reverse 单链表反转，时间复杂度：O(N)
func (l *LinkedList) Reverse() {
	if l == nil || l.Head.Next == nil {
		return
	}

	var pre *Node
	for cur := l.Head.Next; cur != nil; {
		//tmp := cur.next
		//cur.next = pre
		//pre = cur
		//cur = tmp
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	l.Head.Next = pre
}

//HasCycle 判断单链表是否有环
func (l *LinkedList) HasCycle() bool {
	if l == nil || l.Head.Next == nil {
		return false
	}

	slow, fast := l.Head.Next, l.Head.Next.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow, fast = slow.Next, fast.Next.Next
	}

	return false
}

//RemoveN 删除第 N 个节点
func (l *LinkedList) RemoveN(n int) *Node {
	if n == 0 || l == nil || l.Head.Next == nil {
		return nil
	}

	if n < 0 {
		n += l.len + 1
	}
	if n > l.len || n <= 0 {
		return nil
	}

	pre := &l.Head
	for i := 1; i < n; i++ {
		pre = pre.Next
	}

	tmp := pre.Next
	pre.Next = pre.Next.Next
	return tmp
}

//RemoveBottomN 删除倒数第 N 个节点, 长度未知
func (l *LinkedList) RemoveBottomN(n int) *Node {
	if n <= 0 || l == nil || l.Head.Next == nil {
		return nil
	}

	fast := &l.Head
	for i := 0; i < n; i++ {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next
	}

	slow := &l.Head
	for ; fast.Next != nil; fast = fast.Next {
		slow = slow.Next
	}

	tmp := slow.Next
	slow.Next = slow.Next.Next
	return tmp
}

//GetMid 获取中间节点
func (l *LinkedList) GetMid() *Node {
	if l == nil || l.Head.Next == nil {
		return nil
	}

	slow, fast := &l.Head, &l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// New 初始化一个带头结点的空链表
func New() *LinkedList { return new(LinkedList) }

// Build 构建链表
func Build(nodes []interface{}) *LinkedList {
	if len(nodes) == 0 {
		return New()
	}

	list := &LinkedList{len: len(nodes)}
	pre := &list.Head
	for i := 0; i < len(nodes); i++ {
		next := &Node{
			Val:  nodes[i],
			Next: nil,
		}
		pre.Next = next
		pre = next
	}

	return list
}

//BuildCircular 构建成环的链表
func BuildCircular(nodes []interface{}, pos interface{}) *LinkedList {
	if len(nodes) == 0 {
		return nil
	}

	var cNode *Node
	list := &LinkedList{len: len(nodes)}
	pre := &list.Head
	for i := 0; i < len(nodes); i++ {
		next := &Node{
			Val:  nodes[i],
			Next: nil,
		}
		pre.Next = next
		pre = next

		if i == pos {
			cNode = next
		}
		if i == len(nodes)-1 {
			pre.Next = cNode
		}
	}

	return list
}

//MergeList 合并有序链表
func MergeList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == &l1.Head || nil == l1.Head.Next {
		return l2
	}
	if nil == l2 || nil == &l2.Head || nil == l2.Head.Next {
		return l1
	}

	l := New()
	cur, cur1, cur2 := &l.Head, l1.Head.Next, l2.Head.Next
	for cur1 != nil && cur2 != nil {
		if cur1.Val.(int) < cur2.Val.(int) {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		l.len++
		cur = cur.Next
	}

	if cur1 != nil {
		cur.Next = cur1
	} else if cur2 != nil {
		cur.Next = cur2
	}
	return l
}
