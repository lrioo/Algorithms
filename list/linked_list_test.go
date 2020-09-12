package list

import (
	"testing"
)

func TestLinkedList_Reverse(t *testing.T) {
	l := Build([]interface{}{1, 2, 3})
	t.Log(l)
	l.Reverse()
	t.Log(l)
}

func TestLinkedList_HasCycle(t *testing.T) {
	l := BuildCircular([]interface{}{0, 1, 2, 3, 4}, 2)
	t.Log(l)
	t.Log(l.HasCycle())

	l = BuildCircular([]interface{}{0, 1, 2, 3, 4}, 0)
	t.Log(l)
	t.Log(l.HasCycle())

	l = BuildCircular([]interface{}{0, 1, 2, 3, 4}, -1)
	t.Log(l)
	t.Log(l.HasCycle())
}

func TestLinkedList_RemoveN(t *testing.T) {
	l := Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveN(0), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveN(1), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveN(3), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveN(5), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveN(6), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveN(-1), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveN(-2), l)
}

func TestLinkedList_RemoveBottomN(t *testing.T) {
	l := Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveBottomN(0), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveBottomN(1), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveBottomN(3), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveBottomN(5), l)

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.RemoveBottomN(6), l)
}

func TestLinkedList_GetMid(t *testing.T) {
	l := Build([]interface{}{0})
	t.Log(l)
	t.Log(l.GetMid())

	l = Build([]interface{}{0, 1})
	t.Log(l)
	t.Log(l.GetMid())

	l = Build([]interface{}{0, 1, 2})
	t.Log(l)
	t.Log(l.GetMid())

	l = Build([]interface{}{0, 1, 2, 3})
	t.Log(l)
	t.Log(l.GetMid())

	l = Build([]interface{}{0, 1, 2, 3, 4})
	t.Log(l)
	t.Log(l.GetMid())
}

func TestMergeList(t *testing.T) {
	l1 := Build([]interface{}{0, 2, 4})
	l2 := Build([]interface{}{1, 3, 5})
	t.Log(l1, l2)
	l := MergeList(l1, l2)
	t.Log(l)
}
