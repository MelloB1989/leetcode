package main

type Node struct {
	prev  *Node
	next  *Node
	value int
	key   int
}

func NewNode(k, v int) *Node {
	return &Node{
		key:   k,
		value: v,
	}
}

type LRUCache struct {
	capacity int
	store    map[int]*Node
	left     *Node
	right    *Node
}

func Constructor(capacity int) LRUCache {
	left := NewNode(0, 0)
	right := NewNode(0, 0)

	left.next, right.prev = right, left

	return LRUCache{
		capacity: capacity,
		store:    make(map[int]*Node),
		right:    right,
		left:     left,
	}
}

func (this *LRUCache) insertRight(node *Node) {
	latest := this.right.prev
	node.prev, node.next = latest, this.right
	latest.next, this.right.prev = node, node
}

func (c *LRUCache) removeNode(node *Node) {
	if node == nil || node.prev == nil || node.next == nil {
		return
	}
	node.prev.next, node.next.prev = node.next, node.prev
	node.prev, node.next = nil, nil
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.store[key]; ok {
		this.removeNode(v)
		this.insertRight(v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.store[key]; ok {
		v.value = value
		this.removeNode(v)
		this.insertRight(v)
		return
	}
	node := NewNode(key, value)
	this.store[key] = node
	this.insertRight(node)
	if len(this.store) > this.capacity {
		lru := this.left.next
		this.removeNode(lru)
		delete(this.store, lru.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
