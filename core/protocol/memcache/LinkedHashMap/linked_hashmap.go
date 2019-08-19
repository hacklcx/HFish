package LinkedHashMap

import (
	"sync"
)

type LinkedHashMapNode struct {
	linklistNode *LinkListNode
	val          interface{}
}

type LinkedHashMap struct {
	linklist *LinkList
	hashmap  map[string]interface{}
	mutex    *sync.RWMutex
}

func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{
		linklist: NewLinkList(),
		hashmap:  make(map[string]interface{}),
		mutex: &sync.RWMutex{},
	}
}

func (this *LinkedHashMap) Lock() {
	this.mutex.Lock()
}

func (this *LinkedHashMap) Unlock() {
	this.mutex.Unlock()
}

func (this *LinkedHashMap) RLock() {
	this.mutex.RLock()
}

func (this *LinkedHashMap) RUnlock() {
	this.mutex.RUnlock()
}

func (this *LinkedHashMap) Add(key string, val interface{}) bool {
	_, isExists := this.hashmap[key]
	if isExists {
		return false
	}

	linkListNode := this.linklist.AddToTail(key)
	this.hashmap[key] = &LinkedHashMapNode{
		linklistNode: linkListNode,
		val:          val,
	}

	return true
}

func (this *LinkedHashMap) Get(key string) interface{} {
	originLinkedHashMapNode, isExists := this.hashmap[key]
	if !isExists {
		return nil
	}

	return (originLinkedHashMapNode.(*LinkedHashMapNode)).val
}

func (this *LinkedHashMap) Len() int {
	return len(this.hashmap)
}


func (this *LinkedHashMap) Remove(key string) (bool, interface{}) {
	originLinkedHashMapNode, isExists := this.hashmap[key]
	if !isExists {
		return false, nil
	}

	linkedHashMapNode := originLinkedHashMapNode.(*LinkedHashMapNode)

	delete(this.hashmap, key)
	this.linklist.RemoveNode(linkedHashMapNode.linklistNode)
	return true, linkedHashMapNode.val
}

func (this *LinkedHashMap) GetLinkList() *LinkList {
	return this.linklist
}