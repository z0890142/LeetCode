package main

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	curr := 0
	id := 0

	for i := 0; i < len(gas); i++ {
		curr += gas[i] - cost[i]
		total += gas[i] - cost[i]

		if curr < 0 {
			id = i + 1
			curr = 0
		}
	}

	if total < 0 {
		id = -1
	}

	return id
}
