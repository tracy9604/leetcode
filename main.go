package main

import (
	"encoding/json"
	"fmt"
	h3v3 "github.com/uber/h3-go/v3"
	"io/ioutil"
	"leetcode/code"
	"leetcode/h3_demo"
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

func main() {
	arr1 := []int{2, 1, 100, 3}
	arr2 := []int{-5, -2, 10, -3, 7}
	d := 6
	fmt.Println(code.FindDistanceValueTwoArrays(arr1, arr2, d))
}
