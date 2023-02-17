package main

import (
	"encoding/json"
	"fmt"
	h3v3 "github.com/uber/h3-go/v3"
	"io/ioutil"
	"leetcode/code"
	"leetcode/h3_demo"
	"math"
	"os"
	"strings"
)

func GetResultFarmGeofence() {
	lat := 10.88675618334471
	long := 106.64220083059412

	fmt.Println(h3v3.ToString(h3v3.FromGeo(h3v3.GeoCoord{
		Latitude:  lat,
		Longitude: long,
	}, 9)))
	//indexes := h3_demo.GetIndex2(lat, long)
	//rs := h3_demo.GetGeoFromH3Indexes3(indexes)
	//file, _ := json.MarshalIndent(rs, "", " ")
	//_ = ioutil.WriteFile("/Users/hau.nguyen1/Documents/code/leetcode/output.json", file, 0644)
	originIndexesMap := make(map[string]int, 0)

	dat, err := os.ReadFile("/Users/hau.nguyen1/Documents/code/leetcode/data1.txt")
	if err != nil {
		panic(err)
	}

	tmpData := strings.Split(string(dat), ",")
	for _, item := range tmpData {
		originIndexesMap[item] = 0
	}

	fmt.Println("len origins = ", len(originIndexesMap))

	startIndex := "8965b574ac7ffff"

	rs := h3_demo.GetIndex(startIndex, originIndexesMap)

	//rs := h3_demo.GetGeoFromH3Indexes(indexes, remainIndexes, originIndexesMap)
	file, _ := json.MarshalIndent(rs, "", " ")

	_ = ioutil.WriteFile("/Users/hau.nguyen1/Documents/code/leetcode/output2.json", file, 0644)
	//
	//f, err := os.Create("/Users/hau.nguyen1/Documents/code/leetcode/output.json")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//_, err = f.WriteString(string(rsByte))
	//if err != nil {
	//	panic(err)
	//}
}

var validStatusMap = map[uint32][]uint32{
	1:  {2, 12, 4, 5, 67, 78, 10, 9},
	4:  {0, 13, 6, 7, 11, 20},
	0:  {14, 54, 52},
	40: {41, 0, 23},
}

var globalStatusMap = map[uint32][]uint32{}

func loadStatus(currentStatus uint32, status []uint32) []uint32 {
	if _, found := globalStatusMap[currentStatus]; found {
		return globalStatusMap[currentStatus]
	}
	globalStatusMap[currentStatus] = []uint32{}
	for _, item := range status {
		if _, found := validStatusMap[item]; found {
			globalStatusMap[currentStatus] = append(globalStatusMap[currentStatus], loadStatus(item, validStatusMap[item])...)
		} else {
			globalStatusMap[item] = []uint32{}
			globalStatusMap[currentStatus] = append(globalStatusMap[currentStatus], item)
		}
	}
	return globalStatusMap[currentStatus]
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(code.FindSpiralMatrix(matrix))
}
