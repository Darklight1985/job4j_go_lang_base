package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size int
	Head *Node
	Tail *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size: size,
	}
}

func (l *LruCache) Put(key string, value string) {
	tail := l.Tail
	if tail == nil {
		firstNode := Node{
			Key:   key,
			Value: value,
		}
		l.Tail = &firstNode
		l.Head = &firstNode
		return
	}
	head := l.Head

	node := Node{
		Key:   key,
		Value: value,
		Next:  head,
	}

	oldNode := node.Next
	if oldNode != nil {
		oldNode.Prev = &node
	}
	l.Head = &node

	size := l.size
	count := 0
	for i := l.Head; i != nil; i = i.Next {
		count++
	}
	if count > size {
		newTail := *tail.Prev
		newTail.Next = nil
		l.Tail = &newTail
	}
}

func (l *LruCache) Get(key string) *string {

	var node Node
	for i := l.Head; i != nil; i = i.Next {
		if i.Key == key {
			prev := i.Prev
			next := i.Next
			if prev != nil {
				prev.Next = next
				next.Prev = prev
			}
			if next == nil && i.Prev != nil {
				prev.Next = nil
			}
			i.Prev = nil
			l.Head.Prev = &node
			i.Next = l.Head
			l.Head = i
			return &i.Value
		}
	}
	return nil
}
