package list

import "fmt"

//Node 链表节点
type Node struct {
	Val  interface{}
	Next *Node
}

// String 打印当前节点
func (n *Node) String() string {
	if n == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%d", n.Val)
}

// WithExtend 从当前节点开始打印链表
func (n *Node) WithExtend() string {
	if n == nil {
		return "<nil>"
	}

	s := fmt.Sprintf("%d", n.Val)
	cur := n.Next
	for cur != nil {
		s = fmt.Sprintf("%s->%d", s, cur.Val)
		cur = cur.Next
	}
	return s
}

// BuildList 构建单链表
func BuildList(input []interface{}) *Node {
	if len(input) == 0 {
		return nil
	}

	head := &Node{
		Val:  input[0],
		Next: nil,
	}

	pre := head
	for i := 1; i < len(input); i++ {
		next := &Node{
			Val:  input[i],
			Next: nil,
		}
		pre.Next = next
		pre = next
	}

	return head
}

//BuildCircularList 构建成环的链表
func BuildCircularList(input []interface{}, pos interface{}) *Node {
	if len(input) == 0 {
		return nil
	}

	head := &Node{
		Val:  input[0],
		Next: nil,
	}

	var cNode *Node
	if pos == 0 {
		cNode = head
	}

	pre := head
	for i := 1; i < len(input); i++ {
		next := &Node{
			Val:  input[i],
			Next: nil,
		}
		pre.Next = next
		pre = next

		if i == pos {
			cNode = next
		}
		if i == len(input)-1 {
			pre.Next = cNode
		}
	}

	return head
}

//BuildIntersectList 构建两条相交链表
func BuildIntersectList(intersectVal interface{}, listA []interface{}, listB []interface{}, skipA, skipB int) (*Node, *Node) {
	if intersectVal == 0 {
		return BuildList(listA), BuildList(listB)
	}

	var intersection *Node
	headA := &Node{
		Val:  listA[0],
		Next: nil,
	}

	pre := headA
	for i := 1; i < len(listA); i++ {
		next := &Node{
			Val:  listA[i],
			Next: nil,
		}
		pre.Next = next
		pre = next

		if i == skipA {
			intersection = next
		}
	}

	headB := &Node{
		Val:  listB[0],
		Next: nil,
	}

	pre = headB
	for i := 1; i < skipB; i++ {
		next := &Node{
			Val:  listB[i],
			Next: nil,
		}
		pre.Next = next
		pre = next
	}
	pre.Next = intersection

	return headA, headB
}
