package structures

import "fmt"

const ArraySize = 7

// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}

// Insert: Add key to hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Exists: Return true if key is in hash table
func (h *HashTable) Exists(key string) bool {
	index := hash(key)
	return h.array[index].exists(key)
}

// Delete key from hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Bucket (linked list) structure
type bucket struct {
	head *bucketNode
}

// Bucket node: A linked list node holding a key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert node into bucket
func (b *bucket) insert(k string) {
	if !b.exists(k) {
		node := &bucketNode{key: k, next: b.head}
		b.head = node
	} else {
		fmt.Println("already exists")
	}
}

// Return true if key exists in bucket
func (b *bucket) exists(k string) bool {
	node := b.head
	for node != nil {
		if node.key == k {
			return true
		}
		node = node.next
	}
	return false
}

// Delete node with specified key, from bucket
func (b *bucket) delete(k string) {
	// Empty list
	if b.head == nil {
		return
	}

	// First node matches; update head
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prev := b.head
	for prev.next != nil {
		if prev.next.key == k {
			prev.next = prev.next.next
			return
		}
		prev = prev.next
	}
}

// Hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

// Init creates a bucket in each hash table array slot
func Init() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}
	return hashTable
}

func main() {
	hashTable := Init()
	list := []string{
		"Fido",
		"Spot",
		"Rover",
		"Mittens",
		"Paws",
		"Claws",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println(hashTable.Exists("Spot"))
	hashTable.Delete(("Spot"))
	fmt.Println(hashTable.Exists("Spot"))
}
