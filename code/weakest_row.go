package code

import "sort"

func FindWeakestRows(mat [][]int, k int) []int {
	tmpArr := make([]int, 0)
	result := make([]int, 0)
	tmpMap := make(map[int][]int)
	prevSoldiers := 0
	curSoldiers := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				curSoldiers++
			}
		}

		if i > 0 && prevSoldiers <= curSoldiers {
			if _, found := tmpMap[prevSoldiers]; !found {
				tmpArr = append(tmpArr, prevSoldiers)
			}
			tmpMap[prevSoldiers] = append(tmpMap[prevSoldiers], i-1)
		}

		if i == len(mat)-1 && prevSoldiers <= curSoldiers {
			if _, found := tmpMap[curSoldiers]; !found {
				tmpArr = append(tmpArr, curSoldiers)
			}
			tmpMap[curSoldiers] = append(tmpMap[curSoldiers], i)
			break
		}

		prevSoldiers = curSoldiers
		curSoldiers = 0
	}

	sort.Ints(tmpArr)

	for i := 0; i < len(tmpArr); i++ {
		if len(result) >= k {
			break
		}
		result = append(result, tmpMap[tmpArr[i]]...)
	}

	return result[:k]
}
