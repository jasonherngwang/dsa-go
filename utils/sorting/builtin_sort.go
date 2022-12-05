package main

import (
	"fmt"
	"sort"
)

// Custom type on which we can implement required sorting methods
type byLength []string

// Called once to determine slice length
func (s byLength) Len() int {
	return len(s)
}

// Called N*log(N) times
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Called N*log(N) times
// Less() defines comparison logic for determining sorting order.
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// Built-in sort
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// Custom sort-by
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
