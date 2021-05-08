package solution

// contains is a helper function that checks if element e is in the slice
func contains(slice []int, e int) bool {
	for _, v := range slice {
		if v == e {
			return true
		}
	}
	return false
}

func sumToN(n int) int {
	total := n

	for i := 0; i < n; i++ {
		total += i
	}
	return total
}
