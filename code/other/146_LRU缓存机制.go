package other

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func ConstructorLRU(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode, 0),
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	cache.head.next = cache.tail
	cache.tail.pre = cache.head
	return cache
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	this.head.next = node
	node.next.pre = node
	node.pre = this.head
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key, value int) {
	node := initDLinkedNode(key, value)
	if n, ok := this.cache[key]; ok {
		n.value = value    // 注意这里要修改原来的node值，而不能直接把node塞到this.cache里
		this.moveToHead(n) // 注意这里要将n moveToHead，而不是node
	} else {
		this.addToHead(node)
		this.cache[key] = node
		if this.size == this.capacity {
			tailNode := this.removeTail()
			delete(this.cache, tailNode.key)
		} else {
			this.size++
		}
	}
}
