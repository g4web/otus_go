package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    newItems(capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {

	item, wasInCache := l.items[key]
	if wasInCache {
		updateValue(key, value, item, l)
	} else {
		pushValue(key, value, l)
	}

	return wasInCache
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	item, wasInCache := l.items[key]
	if wasInCache {
		l.queue.MoveToFront(item)
		switch cacheItemVal := item.Value.(type) {
		case cacheItem:
			return cacheItemVal.value, wasInCache
		default:
			//что делать, бросать ошибку? Правильным было бы конкретизировать интерфейс для ListItem.Value, а не вот это вот все
		}
	}

	return nil, wasInCache
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = newItems(l.capacity)
}

func newItems(capacity int) map[Key]*ListItem {
	return make(map[Key]*ListItem, capacity)
}

func updateValue(key Key, value interface{}, item *ListItem, l *lruCache) {
	item.Value = cacheItem{key, value}
	l.queue.MoveToFront(item)
}

func pushValue(key Key, value interface{}, l *lruCache) {
	if len(l.items) == l.capacity {
		deleteOldestItem(l)
	}
	l.queue.PushFront(cacheItem{key, value})
	l.items[key] = l.queue.Front()
}

func deleteOldestItem(l *lruCache) {
	itemForDel := l.queue.Back()
	l.queue.Remove(itemForDel)
	switch cacheItemVal := itemForDel.Value.(type) {
	case cacheItem:
		delete(l.items, cacheItemVal.key)
	default:
		//что делать, бросать ошибку? Правильным было бы конкретизировать интерфейс для ListItem.Value, а не вот это вот все
	}
}
