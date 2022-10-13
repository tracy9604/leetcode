package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"leetcode/h3_demo"
	"os"
	"strings"
)

func main() {
	//lat := 21.018848843079834
	//long := 105.78389380540222
	//indexes := h3_demo.GetIndex2(lat, long)
	//rs := h3_demo.GetGeoFromH3Indexes3(indexes)
	//file, _ := json.MarshalIndent(rs, "", " ")
	//_ = ioutil.WriteFile("/Users/hau.nguyen1/Documents/code/leetcode/output.json", file, 0644)
	originIndexesMap := make(map[string]int, 0)

	dat, err := os.ReadFile("/Users/hau.nguyen1/Documents/code/leetcode/data.txt")
	if err != nil {
		panic(err)
	}

	tmpData := strings.Split(string(dat), ",")
	for _, item := range tmpData {
		originIndexesMap[item] = 0
	}

	fmt.Println("len origins = ", len(originIndexesMap))

	startIndex := "894142a5db3ffff"

	rs := h3_demo.GetIndex(startIndex, originIndexesMap)

	//rs := h3_demo.GetGeoFromH3Indexes(indexes, remainIndexes, originIndexesMap)
	file, _ := json.MarshalIndent(rs, "", " ")

	_ = ioutil.WriteFile("/Users/hau.nguyen1/Documents/code/leetcode/output.json", file, 0644)
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
