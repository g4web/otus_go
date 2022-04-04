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
	if l.countOfItems == 0 {
		return nil
	}

	return l.frontItem
}

func (l *list) Back() *ListItem {

	if l.countOfItems == 0 {
		return nil
	}

	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	frontItem := l.frontItem
	l.frontItem = &ListItem{v, frontItem, nil}

	if frontItem != nil {
		frontItem.Prev = l.frontItem
	}

	if l.countOfItems == 0 {
		l.backItem = l.frontItem
	}

	if l.countOfItems == 1 {
		if l.backItem != nil {
			l.backItem.Prev = l.frontItem
		}
	}

	l.countOfItems++

	return l.frontItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	backItem := l.backItem
	l.backItem = &ListItem{v, nil, backItem}

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

	if l.frontItem == i {
		l.frontItem = i.Next
		if l.frontItem != nil {
			l.frontItem.Next = nil
		}
		return
	}

	if l.backItem == i {
		l.backItem = i.Prev
		if l.backItem != nil {
			l.backItem.Next = nil
		}
		return
	}

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
}

func (l *list) MoveToFront(movableElement *ListItem) {

	leftElement := movableElement.Prev
	rightElement := movableElement.Next

	if leftElement == nil {
		return
	}

	movableElement.Prev = nil
	movableElement.Next = l.frontItem

	l.frontItem = movableElement
	if rightElement == nil {
		l.backItem = leftElement
	}

	leftElement.Next = rightElement
	if rightElement != nil {
		rightElement.Prev = leftElement
	}
}
