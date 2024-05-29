package BFS

func CourseSchedule(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegrees := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		course := prerequisite[0]
		prerequisiteCourse := prerequisite[1]

		if _, found := graph[prerequisiteCourse]; !found {
			graph[prerequisiteCourse] = make([]int, 0)
		}
		graph[prerequisiteCourse] = append(graph[prerequisiteCourse], course)
		inDegrees[course]++
	}

	queue := make([]int, 0)

	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	processedCourse := 0

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		processedCourse++

		prerequisitesCourse, _ := graph[course]

		for _, neighbor := range prerequisitesCourse {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return processedCourse == numCourses
}
