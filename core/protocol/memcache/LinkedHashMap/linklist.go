package LinkedHashMap

type LinkListNode struct {
	last *LinkListNode
	next *LinkListNode
	val  interface{}
}

func NewLinkListNode(last *LinkListNode, next *LinkListNode, val interface{}) *LinkListNode {
	node := &LinkListNode{
		last: last,
		next: next,
		val:  val,
	}
	return node
}

func (this *LinkListNode) SetLast(node *LinkListNode) {
	this.last = node
}

func (this *LinkListNode) SetNext(node *LinkListNode) {
	this.next = node
}

func (this *LinkListNode) GetLast() *LinkListNode {
	return this.last
}

func (this *LinkListNode) GetNext() *LinkListNode {
	return this.next
}

func (this *LinkListNode) GetVal() interface{} {
	return this.val
}

func (this *LinkListNode) IsHead() bool {
	return this.last == nil
}

func (this *LinkListNode) IsTail() bool {
	return this.next == nil
}

type LinkList struct {
	head   *LinkListNode
	tail   *LinkListNode
	length int
}

func NewLinkList() *LinkList {
	return &LinkList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (this *LinkList) GetHead() *LinkListNode {
	return this.head
}

func (this *LinkList) GetTail() *LinkListNode {
	return this.tail
}

func (this *LinkList) AddToHead(val interface{}) *LinkListNode {
	if this.head == nil && this.tail == nil {
		return this.addFirstNode(val)

	}
	node := NewLinkListNode(nil, this.head, val)
	this.head.SetLast(node)
	this.head = node
	this.length++
	return node
}

func (this *LinkList) AddToTail(val interface{}) *LinkListNode {
	if this.head == nil && this.tail == nil {
		return this.addFirstNode(val)

	}
	node := NewLinkListNode(this.tail, nil, val)
	this.tail.SetNext(node)
	this.tail = node
	this.length++
	return node
}

func (this *LinkList) RemoveNode(node *LinkListNode) {
	defer func() {
		this.length--
	}()

	/* LinkList中只有1个元素 */
	if node.IsHead() && node.IsTail() {
		this.head = nil
		this.tail = nil
		return
	}

	/* 节点是头节点 */
	if node.IsHead() {
		nextNode := node.GetNext()
		this.head = nextNode
		nextNode.SetLast(nil)
		node.SetNext(nil)
		return
	}

	/* 节点是尾节点 */
	if node.IsTail() {
		lastNode := node.GetLast()
		this.tail = lastNode
		lastNode.SetNext(nil)
		node.SetLast(nil)
		return
	}

	lastNode := node.GetLast()
	nextNode := node.GetNext()

	lastNode.SetNext(nextNode)
	nextNode.SetLast(lastNode)
	node.SetLast(nil)
	node.SetNext(nil)
}

func (this *LinkList) GetLength() int {
	return this.length
}

func (this *LinkList) addFirstNode(val interface{}) *LinkListNode {
	node := NewLinkListNode(nil, nil, val)
	this.head = node
	this.tail = node
	this.length++
	return node
}
