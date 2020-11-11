package EvaluateReversePolishNotation

import "strconv"

func evalRPN(tokens []string) int {
	heap := []int{}
	for _, val := range tokens {

		switch val {
		case "+":
			tmp := heap[len(heap)-2] + heap[len(heap)-1]
			heap = heap[:len(heap)-2]
			heap = append(heap, tmp)
		case "-":
			tmp := heap[len(heap)-2] - heap[len(heap)-1]
			heap = heap[:len(heap)-2]
			heap = append(heap, tmp)
		case "*":
			tmp := heap[len(heap)-2] * heap[len(heap)-1]
			heap = heap[:len(heap)-2]
			heap = append(heap, tmp)
		case "/":
			tmp := heap[len(heap)-2] / heap[len(heap)-1]
			heap = heap[:len(heap)-2]
			heap = append(heap, tmp)
		default:
			num, _ := strconv.Atoi(val)
			heap = append(heap, num)
		}
	}
	return heap[0]
}
