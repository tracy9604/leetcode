package queue

import "strconv"

func popQueue(queue []string) []string {
	if len(queue) == 0 {
		return queue
	}
	if len(queue) == 1 {
		return []string{}
	}
	return queue[1:]
}

func IsOpenTheLock(deadends []string, target string) int {
	start := "0000"
	if start == target {
		return 0
	}
	avoid := make(map[string]bool)
	for _, item := range deadends {
		avoid[item] = true
		if start == item || target == item {
			return -1
		}
	}

	queue := make([]string, 0)
	queue = append(queue, start)

	count := 0
	for len(queue) != 0 {
		count++

		size := len(queue)
		for j := 0; j < size; j++ {
			curr := queue[0]
			queue = popQueue(queue)

			for i := 0; i < len(curr); i++ {
				item, _ := strconv.Atoi(string(curr[i]))

				// increase digit and check
				incrItem := item + 1
				if incrItem > 9 {
					incrItem = 0
				}
				curr = curr[:i] + strconv.Itoa(incrItem) + curr[i+1:]

				if _, found := avoid[curr]; !found {
					queue = append(queue, curr)
					avoid[curr] = true
				}

				if curr == target {
					return count
				}

				// decrease digit and check
				decrItem := item - 1
				if decrItem < 0 {
					decrItem = 9
				}
				curr = curr[:i] + strconv.Itoa(decrItem) + curr[i+1:]

				if _, found := avoid[curr]; !found {
					queue = append(queue, curr)
					avoid[curr] = true
				}

				if curr == target {
					return count
				}

				// restore
				curr = curr[:i] + strconv.Itoa(item) + curr[i+1:]
			}
		}
	}
	return -1
}
