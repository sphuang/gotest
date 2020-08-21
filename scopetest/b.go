package scopetest

// GetA can access non-export variable from other file
func GetA() int {
	return a
}

// GetB can access non-export function from other file
func GetB() int {
	return getb()
}
