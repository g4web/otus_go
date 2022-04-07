package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	frontItem    *ListItem
	backItem     *ListItem
	countOfItems int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.countOfItems
}

func (l *list) Front() *ListItem {
	return l.frontItem
}

func (l *list) Back() *ListItem {
	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	frontItem := l.frontItem
	l.frontItem = &ListItem{Value: v, Next: frontItem, Prev: nil}

	if frontItem != nil {
		frontItem.Prev = l.frontItem
	}

	if l.countOfItems == 0 {
		l.backItem = l.frontItem
	}

	if l.countOfItems == 1 {
		l.backItem.Prev = l.frontItem
	}

	l.countOfItems++

	return l.frontItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	backItem := l.backItem
	l.backItem = &ListItem{Value: v, Next: nil, Prev: backItem}

	if backItem != nil {
		backItem.Next = l.backItem
	}

	if l.countOfItems == 0 {
		l.frontItem = l.backItem
	}

	if l.countOfItems == 1 {
		l.frontItem.Next = l.backItem
	}

	l.countOfItems++

	return l.backItem
}

func (l *list) Remove(i *ListItem) {
	if l.countOfItems > 0 {
		l.countOfItems--
	}

	if l.countOfItems == 0 {
		l.frontItem = nil
		l.backItem = nil

		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.frontItem = l.frontItem.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.backItem = l.backItem.Prev
	}
}

func (l *list) MoveToFront(movableElement *ListItem) {
	l.Remove(movableElement)
	l.PushFront(movableElement.Value)
}
