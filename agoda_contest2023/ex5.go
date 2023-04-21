package agoda_contest2023

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type City struct {
	Name        string
	Days        int
	MoneyPerDay int
}

func Ex5Input() {
	reader := bufio.NewReader(os.Stdin)
	idx := 0
	N := 0
	totalDays := 0
	cities := make([]City, 0)
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if idx == 0 {
			N, _ = strconv.Atoi(text)
		} else {
			textSplits := strings.Split(text, " ")
			days, _ := strconv.Atoi(textSplits[1])
			moneyPerDay, _ := strconv.Atoi(textSplits[2])
			cities = append(cities, City{
				Name:        textSplits[0],
				Days:        days,
				MoneyPerDay: moneyPerDay,
			})
			totalDays += days
		}
		if idx == N {
			days, money := SolveEx5(cities, totalDays)
			fmt.Printf("%d %d", days, money)
			break
		}
		idx++
	}
}

func CountMoney(days int, cities []City, citiesMap map[string][]City, citiesDays map[string]int) (int, int) {
	rsMoney := 0
	rsDays := 0
	if days == 1 {
		for _, item := range cities {
			rsMoney += item.MoneyPerDay
			rsDays += 1
		}
		return rsDays, rsMoney
	} else {
		for key, val := range citiesDays {
			if val >= days {
				rsDays += days

				// count money
				daysToCount := days
				citiesItem := citiesMap[key]
				idx := 0
				for daysToCount > 0 {
					if citiesItem[idx].Days > daysToCount {
						rsMoney += daysToCount * citiesItem[idx].MoneyPerDay
						break
					}

					if idx >= len(citiesItem) {
						break
					}
					rsMoney += citiesItem[idx].Days * citiesItem[idx].MoneyPerDay
					daysToCount -= citiesItem[idx].Days
					idx++
				}
			}
		}
	}
	return rsDays, rsMoney
}

func SolveEx5(cities []City, totalDays int) (int, int) {
	citiesMap := make(map[string][]City)
	citiesDays := make(map[string]int)
	maxDay := int(math.Ceil(float64(totalDays) / float64(len(cities))))
	for _, item := range cities {
		citiesMap[item.Name] = append(citiesMap[item.Name], item)
		citiesDays[item.Name] += item.Days
	}

	for _, item := range citiesMap {
		sort.Slice(item, func(i, j int) bool {
			return item[i].MoneyPerDay < item[j].MoneyPerDay
		})
	}

	maxRsDay := 0
	minRsMoney := int(math.Pow10(5)) * len(cities)
	for i := 1; i <= maxDay; i++ {
		resultDay, resultMoney := CountMoney(i, cities, citiesMap, citiesDays)
		if resultDay >= maxRsDay {
			if resultMoney < minRsMoney {
				minRsMoney = resultMoney
			}
			maxRsDay = resultDay
		}
	}

	return maxRsDay, minRsMoney
}
