// https://github.com/austingebauer/go-leetcode/blob/master/lru_cache_146/solution_test.go

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createLRUCache() LRUCache {
	node1 := &DLLNode{
		key: 1,
		val: 1,
	}
	lRUCache := LRUCache{
		cache:    map[int]*DLLNode{},
		capacity: 1,
		head: &DLLNode{
			next: node1,
		},
		tail: &DLLNode{
			prev: node1,
		},
	}
	lRUCache.cache[1] = node1
	node1.next = lRUCache.tail
	node1.prev = lRUCache.head

	return lRUCache
}

func Test_LRUCacheGet(t *testing.T) {
	type input struct {
		lRUCache LRUCache
		key      int
	}

	type testCase struct {
		name     string
		input    input
		expected int
	}

	tests := []testCase{
		{
			name: "Get existing key",
			input: input{
				lRUCache: createLRUCache(),
				key:      1,
			},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.input.lRUCache.Get(test.input.key)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_LRUCachePutUpdate(t *testing.T) {
	type input struct {
		lRUCache LRUCache
		key      int
		val      int
	}

	type testCase struct {
		name     string
		input    input
		expected int
	}

	tests := []testCase{
		{
			name: "Update existing key",
			input: input{
				lRUCache: createLRUCache(),
				key:      1,
				val:      42,
			},
			expected: 42,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.input.lRUCache.Put(test.input.key, test.input.val)
			actual := test.input.lRUCache.Get(test.input.key)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_LRUCacheEvictLRU(t *testing.T) {
	type input struct {
		lRUCache LRUCache
		key      int
		val      int
	}

	type testCase struct {
		name     string
		input    input
		expected int
	}

	tests := []testCase{
		{
			name: "Successfully evicts LRU",
			input: input{
				lRUCache: createLRUCache(),
				key:      2,
				val:      2,
			},
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.input.lRUCache.Put(test.input.key, test.input.val)
			actual := test.input.lRUCache.Get(1)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_LRUCachePutInsert(t *testing.T) {
	type input struct {
		lRUCache LRUCache
		key      int
		val      int
	}

	type testCase struct {
		name     string
		input    input
		expected int
	}

	tests := []testCase{
		{
			name: "Successfully inserts node",
			input: input{
				lRUCache: createLRUCache(),
				key:      2,
				val:      2,
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.input.lRUCache.Put(test.input.key, test.input.val)
			actual := test.input.lRUCache.Get(test.input.key)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_LRUCache_1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1)) // returns 1

	cache.Put(3, 3)                   // evicts key 2
	assert.Equal(t, -1, cache.Get(2)) // returns -1 (not found)

	cache.Put(4, 4)                   // evicts key 1
	assert.Equal(t, -1, cache.Get(1)) // returns -1 (not found)
	assert.Equal(t, 3, cache.Get(3))  // returns 3
	assert.Equal(t, 4, cache.Get(4))  // returns 4
}

func Test_LRUCache_2(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	assert.Equal(t, 1, cache.Get(2))
}

/*
["LRUCache","put","get","put","get","get"]
[[1],[2,1],[2],[3,2],[2],[3]]
*/
func Test_LRUCache_3(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	assert.Equal(t, 1, cache.Get(2))

	cache.Put(3, 2)
	assert.Equal(t, -1, cache.Get(2))
	assert.Equal(t, 2, cache.Get(3))
}

/*
["LRUCache","put","put","get","put","put","get"]
[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
*/
func Test_LRUCache_4(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	assert.Equal(t, 2, cache.Get(2))
}

/*
["LRUCache","put","put","put","put","get","get"]
[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
Output:   [null,null,null,null,null,1,-1]
Expected: [null,null,null,null,null,-1,3]
*/
func Test_LRUCache_5(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3) // evict 2->1
	cache.Put(4, 1) // evict 1->1

	// left with: 2->3, 4->1
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(2))
}

/*
["LRUCache","put","put","get","put","put","get"]
[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
Output:   [null,null,null,null,null,1,-1]
Expected: [null,null,null,null,null,-1,3]
*/
func Test_LRUCache_6(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	assert.Equal(t, 2, cache.Get(2))
	cache.Put(1, 1)
	cache.Put(4, 1)
	assert.Equal(t, -1, cache.Get(2))
}

/*
["LRUCache","get","put","get","put","put","get","get"]
[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
*/
func Test_LRUCache_7(t *testing.T) {
	cache := Constructor(2)
	assert.Equal(t, -1, cache.Get(2))
	cache.Put(2, 6)
	assert.Equal(t, -1, cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2) // evicts 2->6
	assert.Equal(t, 2, cache.Get(1))
	assert.Equal(t, 6, cache.Get(2))
}

/*
["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
Output: [null,null,null,null,null,4,3,2,-1,null,-1,2,-1,4,5]
Expected: [null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
*/
func Test_LRUCache_8(t *testing.T) {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	assert.Equal(t, 4, cache.Get(4))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 2, cache.Get(2))
	assert.Equal(t, -1, cache.Get(1))
	cache.Put(5, 5)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 2, cache.Get(2))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, -1, cache.Get(4))
	assert.Equal(t, 5, cache.Get(5))
}
