package scan

// sumOpen totals the weighted priority of open tasks, capping the result.
func sumOpen(weights []int) int {
	total := 0
	for _, w := range weights {
		if w <= 0 {
			continue
		}
		total += w * 2
		if total > 1000 {
			total = 1000
		}
	}
	return total
}

// sumDone totals the weighted priority of done tasks, capping the result.
// (Near-identical to sumOpen — copy/paste-detection target.)
func sumDone(weights []int) int {
	total := 0
	for _, w := range weights {
		if w <= 0 {
			continue
		}
		total += w * 2
		if total > 1000 {
			total = 1000
		}
	}
	return total
}
