package circular

// SinglyElement is an element of a single linked list.
type SinglyElement struct {
	// The next Element of list
	next *SinglyElement
	// The list to which this element belongs.
	list *SinglyCircularList
	// The Value of this Element
	Value interface{}
}

// Next returns the next list element or nil.
func (e *SinglyElement) Next() *SinglyElement {
	if e.list != nil && e != e.list.last {
		return e.next
	}
	return nil
}

// SinglyCircularList represents a single linked list.
// The zero value for List is an empty list ready to use.
type SinglyCircularList struct {
	root SinglyElement  // sentinel list element, only &root and root.next are used
	last *SinglyElement // the last element of list
	len  int            // current list len excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *SinglyCircularList) Init() *SinglyCircularList {
	l.root.next = &l.root
	l.last = &l.root
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *SinglyCircularList { return new(SinglyCircularList).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *SinglyCircularList) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *SinglyCircularList) Front() *SinglyElement {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *SinglyCircularList) Back() *SinglyElement {
	if l.len == 0 {
		return nil
	}
	return l.last
}

// lazyInit lazily initializes a zero List value.
func (l *SinglyCircularList) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *SinglyCircularList) insert(e, at *SinglyElement) *SinglyElement {
	n := at.next
	at.next = e
	e.next = n
	e.list = l
	l.len++
	if n == &l.root {
		l.last = e
	}
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *SinglyCircularList) insertValue(v interface{}, at *SinglyElement) *SinglyElement {
	return l.insert(&SinglyElement{Value: v}, at)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *SinglyCircularList) PushFront(v interface{}) *SinglyElement {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *SinglyCircularList) PushBack(v interface{}) *SinglyElement {
	l.lazyInit()
	return l.insertValue(v, l.last)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *SinglyCircularList) InsertBefore(v interface{}, mark *SinglyElement) *SinglyElement {
	if mark.list != l {
		return nil
	}

	pre, cur := &l.root, l.root.next
	for cur != &l.root {
		if cur == mark {
			break
		}

		pre = cur
		cur = cur.next
	}
	if cur == &l.root {
		return nil
	}

	return l.insertValue(v, pre)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *SinglyCircularList) InsertAfter(v interface{}, mark *SinglyElement) *SinglyElement {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// FindByIndex the index of list from 0 begin.
// if the param index < 0, means from the tail to count.
func (l *SinglyCircularList) FindByIndex(index int) *SinglyElement {
	if index >= l.len {
		return nil
	}

	if index < 0 {
		index += l.len
		if index < 0 {
			return nil
		}
	}

	cur := l.root.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// remove removes e from its list, decrements l.len, and returns e.
func (l *SinglyCircularList) remove(e, after *SinglyElement) *SinglyElement {
	after.next = e.next
	e.next = nil // avoid memory leaks
	e.list = nil
	l.len--
	if after.next == &l.root {
		l.last = after
	}
	return e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *SinglyCircularList) Remove(e *SinglyElement) interface{} {
	if e.list != l {
		return nil
	}

	pre, cur := &l.root, l.root.next
	for cur != &l.root {
		if cur == e {
			break
		}

		pre = cur
		cur = cur.next
	}
	if cur == &l.root {
		return nil
	}
	l.remove(e, pre)

	return e.Value
}
