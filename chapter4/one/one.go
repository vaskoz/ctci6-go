package one

// DiGraph represents a directed graph where each vertex is labeled by an integer.
type DiGraph map[int]map[int]struct{}

// RouteExists answers the question "Does a route exist between two given vertices?".
func RouteExists(graph DiGraph, start, end int) bool {
	var stack []int
	discovered := make(map[int]struct{})
	stack = append(stack, start)
	for len(stack) > 0 {
		var v int
		v, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if _, found := discovered[v]; !found {
			discovered[v] = struct{}{}
			for k := range graph[v] {
				if k == end {
					return true
				}
				stack = append(stack, k)
			}
		}
	}
	return false
}
