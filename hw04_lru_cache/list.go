package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый элемент списка
	Back() *listItem                   // последний элемент списка
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value interface{} // значение
	Next  *listItem   // следующий элемент
	Prev  *listItem   // предыдущий элемент
}

type list struct {
	len   int
	first *listItem
	last  *listItem
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *listItem {
	return l.first
}

func (l *list) Back() *listItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *listItem {
	firstItem := &listItem{Value: v}

	if l.Len() == 0 {
		l.last = firstItem
	} else {
		l.first.Prev = firstItem
		firstItem.Next = l.first
	}
	l.first = firstItem

	l.len++
	return firstItem
}

func (l *list) PushBack(v interface{}) *listItem {
	lastItem := &listItem{Value: v}

	if l.Len() == 0 {
		l.first = lastItem
	} else {
		l.last.Next = lastItem
		lastItem.Prev = l.last
	}
	l.last = lastItem

	l.len++
	return lastItem
}

func (l *list) Remove(i *listItem) {
	switch {
	case l.first == i && l.last == i:
		l.first = nil
		l.last = nil
	case l.first == i:
		l.first = l.first.Next
		l.first.Prev = nil
	case l.last == i:
		l.last = l.last.Prev
		l.last.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *listItem) {
	if l.first == i {
		return
	}

	if l.last == i {
		i.Prev.Next = nil
		l.last = i.Prev
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	i.Prev = nil
	i.Next = l.first
	l.first.Prev = i
	l.first = i
}
