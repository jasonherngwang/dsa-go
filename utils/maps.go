package utils

// Tally up counts of elements in a slice
// Map key must be comparable (==, !=)
func Tally[K comparable](slice []K) map[K]int {
	tally := make(map[K]int)

	for _, elem := range slice {
		if _, ok := tally[elem]; !ok {
			tally[elem] = 0
		}
		tally[elem]++
	}

	return tally
}

// Convert slice to a set (unique values), implemented as a map
// From CodeWars problem "Array Diff"
func SliceToSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool)

	for _, elem := range slice {
		set[elem] = true
	}

	return set
}
