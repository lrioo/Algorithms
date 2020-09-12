package list

import (
	"testing"
)

func TestIsPalindrome1(t *testing.T) {
	l := Build([]interface{}{1, 2, 3, 2, 1})
	t.Log(l)
	t.Log(IsPalindrome1(l))

	l = Build([]interface{}{1, 2, 2, 1})
	t.Log(l)
	t.Log(IsPalindrome1(l))

	l = Build([]interface{}{1, 2, 3, 1})
	t.Log(l)
	t.Log(IsPalindrome1(l))
}

func TestIsPalindrome2(t *testing.T) {
	l := Build([]interface{}{1, 2, 3, 2, 1})
	t.Log(l)
	t.Log(IsPalindrome2(l), l)

	l = Build([]interface{}{1, 2, 2, 1})
	t.Log(l)
	t.Log(IsPalindrome2(l), l)

	l = Build([]interface{}{1, 2, 3, 1})
	t.Log(l)
	t.Log(IsPalindrome2(l), l)
}
