package CourseSchedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, coures := range prerequisites {
		graph[coures[1]] = append(graph[coures[1]], coures[0])
		inDegree[coures[0]] += 1
	}

	bfs := make([]int, 0)
	for course, d := range inDegree {
		if d == 0 {
			bfs = append(bfs, course)
		}
	}

	for i := 0; i < len(bfs); i++ {
		course := bfs[i]
		for _, j := range graph[course] {
			inDegree[j] -= 1
			if inDegree[j] == 0 {
				bfs = append(bfs, j)
			}
		}
	}

	if len(bfs) == numCourses {
		return true
	}

	return false
}
