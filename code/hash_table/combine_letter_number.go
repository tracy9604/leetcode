package hash_table

import "strings"

func FindLetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	lettersMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	result := make([]string, 0)
	output := make([]string, 0)
	input := strings.Split(digits, "")
	findCombinations(lettersMap, input, output, 0, len(input), &result)
	return result
}

func findCombinations(keyPad map[string][]string, input []string, output []string, curr, n int, result *[]string) {
	if curr == n {
		*result = append(*result, strings.Join(output, ""))
		return
	}
	keys := keyPad[input[curr]]
	for i := 0; i < len(keys); i++ {
		output = append(output, keys[i])
		findCombinations(keyPad, input, output, curr+1, n, result)
		output = output[:(len(output) - 1)]

		if input[curr] == "0" || input[curr] == "1" {
			return
		}
	}
}
