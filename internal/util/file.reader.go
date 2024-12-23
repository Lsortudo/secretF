package util

// Filters even-numbered names
func FilterEven(people []string) []string {
	var result []string
	for i, person := range people {
		if i%2 == 0 { // Apenas índices pares
			result = append(result, person)
		}
	}
	return result
}
