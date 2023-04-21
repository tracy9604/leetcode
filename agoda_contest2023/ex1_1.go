package agoda_contest2023

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Task struct {
	Priority     int
	Dependencies []int
}

func Ex11Input() {
	reader := bufio.NewReader(os.Stdin)
	tasksMap := make(map[int][]int)
	idx := 0
	N := 0
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		textSplits := strings.Split(text, " ")

		if idx == 0 {
			N, _ = strconv.Atoi(textSplits[1])
		} else {
			priority, _ := strconv.Atoi(textSplits[0])
			dependency, _ := strconv.Atoi(textSplits[1])
			if _, found := tasksMap[priority]; !found {
				tasksMap[priority] = make([]int, 0)
			}
			tasksMap[priority] = append(tasksMap[priority], dependency)
			if _, found := tasksMap[dependency]; !found {
				tasksMap[dependency] = make([]int, 0)
			}
		}
		if idx == N {
			fmt.Println(tasksMap)
			//fmt.Println(SolveEx11(tasksMap))
			break
		}
		idx++

	}
}

func SolveEx11(tasksMap map[int][]int) {
	rs := make([]int, 0)
	indenTasks := make(map[int][]int)
	depenTasks := make(map[int]map[int]int)
	for task, item := range tasksMap {
		if len(item) == 0 {
			indenTasks[task] = make([]int, 0)
			rs = append(rs, task)
		} else {
			subDependTask := make(map[int]int)
			for _, i1 := range item {
				subDependTask[i1] = 1
			}
			depenTasks[task] = subDependTask
		}
	}

	sort.Slice(rs, func(i, j int) bool {
		return i > j
	})

	//for task := range indenTasks {
	//	for t1, item := range depenTasks {
	//		if _, found := item[task]; found {
	//
	//		}
	//	}
	//}
}
