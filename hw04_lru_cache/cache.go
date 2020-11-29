package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	capacity int               // - capacity
	queue    List              // - queue
	items    map[Key]*listItem // - items
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	item, status := l.items[key]

	if status {
		item.Value = cacheItem{key, value}
		l.items[key] = item
		l.queue.MoveToFront(item)
	} else {
		item = l.queue.PushFront(cacheItem{key, value})
		l.items[key] = item

		if l.queue.Len() > l.capacity {
			lastItem := l.queue.Back()
			l.queue.Remove(lastItem)
			delete(l.items, lastItem.Value.(cacheItem).key)
		}
	}
	return status
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	item, status := l.items[key]
	if !status {
		return nil, status
	}

	l.queue.MoveToFront(item)
	return item.Value.(cacheItem).value, status
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*listItem)
}

type cacheItem struct {
	key   Key         // - ключ, по которому он лежит в словаре
	value interface{} // - само значение
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem),
	}
}
