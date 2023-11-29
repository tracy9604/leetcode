package contest12_11

import (
	"sort"
	"strconv"
)

func FindHighAccessEmployee(access_times [][]string) []string {
	employeeMap := make(map[string][]string)

	for i := 0; i < len(access_times); i++ {
		if _, found := employeeMap[access_times[i][0]]; !found {
			employeeMap[access_times[i][0]] = make([]string, 0)
		}
		employeeMap[access_times[i][0]] = append(employeeMap[access_times[i][0]], access_times[i][1])
	}

	keys := make([]string, 0)
	data := make(map[string]int)
	for key, val := range employeeMap {
		sort.Strings(val)
		for i := 0; i < len(val); i++ {
			startTimeHour, _ := strconv.Atoi(val[i][:2])
			startTimeMinute, _ := strconv.Atoi(val[i][2:])
			startTime, _ := strconv.Atoi(val[i])
			endTimeMinute := startTimeMinute + 59
			endTimeHour := startTimeHour
			if endTimeMinute >= 60 {
				endTimeMinute -= 60
				endTimeHour += 1
			}
			endTime, _ := strconv.Atoi(strconv.Itoa(endTimeHour) + strconv.Itoa(endTimeMinute))
			if endTimeMinute < 10 {
				endTime, _ = strconv.Atoi(strconv.Itoa(endTimeHour) + "0" + strconv.Itoa(endTimeMinute))
			}

			accessTime := 1

			for j := i + 1; j < len(val); j++ {
				item, _ := strconv.Atoi(val[j])
				if item >= startTime && item <= endTime {
					accessTime++
					if valKey, _ := data[key]; accessTime >= 3 && accessTime > valKey {
						data[key] = accessTime
					}
				}
			}
		}
	}
	for key := range data {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return data[keys[i]] > data[keys[j]]
	})
	return keys
}
