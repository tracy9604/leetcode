package recursion

import "strconv"

func KthGrammarRecursion(n, currRow, target int, prev string) string {
	if currRow > n {
		return ""
	}
	dataRow := ""
	for i := 0; i < len(prev); i++ {
		if string(prev[i]) == "0" {
			dataRow += "01"
		} else if string(prev[i]) == "1" {
			dataRow += "10"
		}
	}
	if currRow == target {
		return dataRow
	}

	return KthGrammarRecursion(n, currRow+1, target, dataRow)

}

func KthGrammar(n int, k int) int {
	if k > n {
		return 0
	}
	tmp := KthGrammarRecursion(n, 1, k, "0")
	rs, _ := strconv.Atoi(tmp)
	return rs
}
