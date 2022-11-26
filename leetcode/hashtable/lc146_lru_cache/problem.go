/*
Problem
LRU Cache
https://leetcode.com/problems/lru-cache/

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:
- LRUCache(capacity int): Initialize the LRU cache with positive size capacity.
- get(key int) int: Return the value of the key if the key exists, otherwise -1.
- put(key int, value int) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
- The functions get and put must each run in O(1) average time complexity.

Inputs:
- Constructor: Cache capacity (int)
- Get: key (int)
- Put: key(int), value (int)
Outputs:
- Constructor: Copy (not pointer to) of LRUCache
- Get: value (int)
- Put: none

Rules, Requirements, Definitions
- Keys are ints, not strings.
- Cache can be no larger than specified capacity.
  - 1 <= capacity <= 3000
- "Used" means an element was retrieved with Get, updated with Put, or added with Put.

Examples, Test Cases, Edge Cases

Example: Start with this doubly linked list of capacity 3
head <-> 2 (MRU) <-> 3 (LRU) <-> tail

Put [1, val]: List is not full, so we can add 1 as MRU.
head <-> 1 (MRU) <-> 2 <-> 3 (LRU) <-> tail

Get 2: 2 becomes MRU
head <-> 2 (MRU) <-> 1 <-> 3 (LRU) <-> tail

Put 4: Evict LRU (3). 4 becomes MRU. 2 and 1 are shifted, with 1 becoming LRU.
head <-> 4 (MRU) <-> 2 <-> 1 (LRU) <-> tail

Put [1, val]: Update 1 and make it LRU. 2 becomes LRU.
head <-> 1 (MRU) <-> 4 <-> 2 (LRU) <-> tail

Put 5: Evict LRU (2). 5 becomes MRU. 4 becomes LRU.
head <-> 5 (MRU) <-> 1 <-> 4 (LRU) <-> tail

Data Structure
- Doubly linked list
- Hash table

Approach 1: Hash table + Doubly linked list

Algorithm
Store keys and values in nodes of a doubly-linked list. Its head is the MRU, and its tail the LRU. Use a hash table for O(1) lookup of node locations. Its keys are the same as the DLL, but its value are pointers to the nodes, NOT the data values.
When a key is retrieved with Get, updated with Put, or added with Put, move that key to the front (head), making it the MRU. The DLL should keep track of its length, to determine if it is full. If we need to evict the LRU element, remove the tail.

For cache of capacity N:
Time: O(1) for Get and Put
Space: O(2N) for hash table of size N and linked list of size N

Doubly linked list data structures
- Create type for DLL node (prev, next, key, value).

Constructor for LRUCache
- Initialize cache to empty hash table.
- Initialize capacity to provided value.
- Initialize head and tail to empty DLL nodes, and point them at each other.

Methods for LRUCache
Remove node
- Identify the left and right neighbors of the specified node.
- Disconnect the node by linking its neighbors together.

Insert node at head of list (as MRU)
- Identify list head and next node.
- Attach the new node to these two node.
- Break the links between head and next, and reconnect them to the new node.

Get
- Check if key in hash table. If not, return -1.
- Make the looked-up node the MRU by removing it from the list and re-inserting it at the front.

Put
- Create a new DLL node with the provided key and value.
- Check if key in hash table:
	- If present:
		- Remove it, since we are going to re-insert it later at the front.
	- If absent:
  		- If length of linked list >= capacity (out of space), we need to evict the LRU:
			- Look up the LRU node at the tail.
			- Delete the LRU's k-v pair from the hash table.
			- Remove the LRU node.
- Add the new node to the front of the linked list.
- Add or update the key in hash table, using the supplied value.
*/

package leetcode

type DLLNode struct {
	key  int
	val  int
	next *DLLNode
	prev *DLLNode
}

type LRUCache struct {
	cache    map[int]*DLLNode
	capacity int
	head     *DLLNode // MRU
	tail     *DLLNode // LRU
}

func Constructor(capacity int) LRUCache {
	lRUCache := LRUCache{
		cache:    map[int]*DLLNode{},
		capacity: capacity,
		head:     &DLLNode{},
		tail:     &DLLNode{},
	}

	lRUCache.head.next = lRUCache.tail
	lRUCache.tail.prev = lRUCache.head

	return lRUCache
}

// Remove node from list
func (this *LRUCache) remove(node *DLLNode) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

// Insert node at head of list
func (this *LRUCache) insert(node *DLLNode) {
	prev, next := this.head, this.head.next
	node.prev, node.next = prev, next
	prev.next, next.prev = node, node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// Make node MRU
		this.remove(node)
		this.insert(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// Key is present; remove it so we can replace it with a new node with updated value
		this.remove(node)
	} else {
		// Key is absent; make space for it by evicting LRU.
		if len(this.cache) >= this.capacity {
			lruNode := this.tail.prev
			delete(this.cache, lruNode.key)
			this.remove(lruNode)
		}
	}

	newNode := &DLLNode{key: key, val: value}
	this.insert(newNode)
	this.cache[key] = newNode
}
