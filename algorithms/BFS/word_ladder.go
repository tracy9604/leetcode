package BFS

func WordLadder(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, item := range wordList {
		wordSet[item] = true
	}
	if _, found := wordSet[endWord]; !found {
		return 0
	}

	queue := make([]string, 0)
	queue = append(queue, beginWord)
	depth := 0

	for len(queue) > 0 {
		depth++
		lSize := len(queue)
		for i := 0; i < lSize; i++ {
			top, tmpQueue := popStringQueue(queue)
			queue = tmpQueue

			if top == endWord {
				return depth + 1
			}

			for j := 0; j < len(top); j++ {
				temp := []byte(top)

				for ch := 'a'; ch < 'z'; ch++ {
					temp[j] = byte(ch)
					if string(temp) == top {
						continue
					}
					if _, found := wordSet[string(temp)]; !found {
						continue
					}
					if string(temp) == endWord {
						return depth + 1
					}

					queue = append(queue, string(temp))
					wordSet[string(temp)] = true
				}
			}
		}
	}

	return 0
}

func popStringQueue(queue []string) (string, []string) {
	if len(queue) == 0 {
		return "", queue
	}
	if len(queue) == 1 {
		return queue[0], []string{}
	}
	return queue[0], queue[1:]
}

func extendQueue(queue []string, wordSet map[string]bool, mapOne, mapTwo map[string]int) int {
	for len(queue) > 0 {
		top, tmpQueue := popStringQueue(queue)
		queue = tmpQueue
		currentStep := 0
		if val, check := mapOne[top]; check {
			currentStep = val
		}
		word := []byte(top)

		for idx, character := range word {
			originalChar := character
			for ch := 'a'; ch < 'z'; ch++ {
				word[idx] = byte(ch)

				_, found1 := wordSet[string(word)]
				_, found2 := mapOne[string(word)]
				if !found1 || found2 {
					continue
				}

				if val, found := mapTwo[string(word)]; found {
					return currentStep + val
				}

				queue = append(queue, string(word))
				mapOne[string(word)] = currentStep + 1
			}
			word[idx] = originalChar
		}
	}

	return -1
}
