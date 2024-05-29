package daily

func MaxScoreWords(words []string, letters []byte, score []int) int {
	// count the frequency of each letter in the available letters
	letterCount := make(map[byte]int)
	for _, letter := range letters {
		letterCount[letter]++
	}

	// calculate the score of each word
	wordScores := make([]int, len(words))
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			wordScores[i] += score[word[j]-'a']
		}
	}

	var backtrack func(index int, currentScore int, letterCount map[byte]int) int

	backtrack = func(index int, currentScore int, letterCount map[byte]int) int {
		if index == len(words) {
			return currentScore
		}

		// skip the current word
		maxScore := backtrack(index+1, currentScore, letterCount)

		// try to include the current word if possible
		return maxScore
	}
	return 0
}
