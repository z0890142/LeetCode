package GasStation

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	cur := 0
	index := 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		cur += gas[i] - cost[i]
		if cur < 0 {
			index = i + 1
			cur = 0
		}
	}
	if total < 0 {
		index = -1
	}

	return index
}
