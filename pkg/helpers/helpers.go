package helpers

// Contains is a helper function that checks if element e is in the slice
func Contains(slice []int, e int) bool {
	for _, v := range slice {
		if v == e {
			return true
		}
	}
	return false
}

// SumToNo returns an int which is the sum from 1 to n.
func SumToN(n int) int {
	total := n

	for i := 0; i < n; i++ {
		total += i
	}
	return total
}
