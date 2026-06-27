package main 

type Cache[T comparable] struct {
	itemsMap map[T]*node[T]
	cap uint
	root *node[T]
	tail *node[T]
}

func NewCache[T comparable](cap uint) *Cache[T] {
	return &Cache[T]{ 
		itemsMap: make(map[T]*node[T], cap),
		cap: cap,
		root: nil,
		tail: nil,
	}
}

type node[T comparable] struct {
	value T 
	next *node[T]
	prev *node[T]
}

func newNode[T comparable](value T) *node[T] {
	return &node[T] {
		value: value, 
		next: nil,
		prev: nil,
	}
}

func remove[T comparable](val *node[T], c *Cache[T]) {
	delete(c.itemsMap, val.value)
	if val.prev == nil && val.next == nil {
		c.root = nil 
		c.tail = nil
		return
	}
	if val.prev != nil {
		val.prev.next = val.next
	} else {
		c.root = val.next
	}

	if val.next != nil {
		val.next.prev = val.prev
	} else {
		c.tail = val.prev
	}
}

func addToFront[T comparable](val *node[T], c *Cache[T]) {
	if c.root == nil {
		c.root = val
		c.tail = val
	} else {
		c.root.prev = val 
		val.next = c.root
		c.root = val
		c.root.prev = nil
	}
	c.itemsMap[val.value] = val
}

func (c *Cache[T])Set(key T){
	if c.cap == 0 {
		return
	}

	if val, ok := c.itemsMap[key]; ok  {
		if val == c.root {
			return 
		}
		remove(val, c)
		addToFront(val, c)

	} else if len(c.itemsMap) < int(c.cap){
		var new *node[T] = newNode(key)
		addToFront(new, c)
	} else {
		var new *node[T] = newNode(key)
		remove(c.tail, c)
		addToFront(new, c)
	}
}

func (c *Cache[T])Get(key T) (T, bool) {
	
}