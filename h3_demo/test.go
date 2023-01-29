package h3_demo

import (
	"errors"
	"fmt"
	h3v3 "github.com/uber/h3-go/v3"
	"math"
	"os"
	"strings"
)

type GeoGeometry struct {
	Type        string         `json:"type"`
	Coordinates [][][2]float64 `json:"coordinates"`
}

type GeoFeatures struct {
	Type       string      `json:"type"`
	Properties Properties  `json:"properties,omitempty"`
	Geometry   GeoGeometry `json:"geometry"`
}

type Properties struct {
	StrokeWidth string `json:"stroke-width,omitempty"`
	Stroke      string `json:"stroke,omitempty"`
	FillColor   string `json:"fill,omitempty"`
}

type Geo struct {
	Type     string        `json:"type"`
	Features []GeoFeatures `json:"features"`
}

var colors = []string{
	"#9c4f49",
	"#338f83",
	"#1f71a3",
	"#752c78",
	"#782c3f",
	"#b0414c",
	"#c9bd5f",
}

func GetIndex(startIndex string, originIndexesMap map[string]int) Geo {
	indexesMap := make(map[h3v3.H3Index]int)
	centersMap := make(map[h3v3.H3Index]int)

	start := h3v3.FromString(startIndex)
	centers := make([]h3v3.H3Index, 0)
	centersMap[start] = 1
	centers = append(centers, start)

	maxGeo := int(math.Ceil(float64(len(originIndexesMap)) / float64(7)))
	fmt.Println("maxGeo = ", maxGeo)

	refinedGeo := Geo{
		Type:     "FeatureCollection",
		Features: make([]GeoFeatures, 0),
	}

	i := 0
	for {
		fmt.Println("i = ", i)
		if i >= maxGeo || i >= len(centers) {
			break
		}

		rs := findNextCenter(centers[i], indexesMap)

		// filter the duplicated center indexes
		for _, item := range rs {
			if _, found := originIndexesMap[h3v3.ToString(item)]; found {
				if _, exists := centersMap[item]; !exists {
					centersMap[item] = 0
					centers = append(centers, item)
					geoFeature, err := getGeoFromIndex(item, originIndexesMap, "")
					if err != nil {
						continue
					}
					refinedGeo.Features = append(refinedGeo.Features, geoFeature)
				}
				continue
			}

			k1 := h3v3.KRing(item, 1)
			for _, itemK1 := range k1 {
				if _, exists := originIndexesMap[h3v3.ToString(itemK1)]; exists {
					geoFeature, err := getGeoFromIndex(item, originIndexesMap, "#f5428a")
					if err != nil {
						continue
					}
					refinedGeo.Features = append(refinedGeo.Features, geoFeature)
					break
				}
			}
		}
		i++
	}

	remainIndexes := make([]string, 0)
	for key, val := range originIndexesMap {
		if val == 0 {
			remainIndexes = append(remainIndexes, key)
			originIndexesMap[key] = 1
		}
	}

	polygons := getCoordinatesFromH3Indexes(remainIndexes)

	properties := Properties{
		StrokeWidth: "3",
		Stroke:      "#000000",
		FillColor:   "#f5f242",
	}

	refinedGeo.Features = append(refinedGeo.Features, GeoFeatures{
		Type:       "Feature",
		Properties: properties,
		Geometry: GeoGeometry{
			Type:        "Polygon",
			Coordinates: [][][2]float64{polygons[0]},
		},
	})

	fmt.Println("len remain =", len(remainIndexes))

	return refinedGeo
}

//func GetIndex2(lat, long float64) []h3v3.H3Index {
//	indexesMap := make(map[h3v3.H3Index]int)
//	centersMap := make(map[h3v3.H3Index]int)
//
//	start := h3v3.FromGeo(h3v3.GeoCoord{
//		Latitude:  lat,
//		Longitude: long,
//	}, 9)
//	centers := make([]h3v3.H3Index, 0)
//	centersMap[start] = 1
//	centers = append(centers, start)
//
//	i := 0
//	for {
//		fmt.Println("i = ", i)
//
//		if len(centers) >= 50 || i >= len(centers) {
//			break
//		}
//
//		rs, rsMap := findNextCenter(centers[i], indexesMap)
//
//		// filter the duplicated center indexes
//		for _, item := range rs {
//			if _, exists := centersMap[item]; !exists {
//				centersMap[item] = 0
//				centers = append(centers, item)
//			}
//		}
//
//		// store the visited indexes
//		for key := range rsMap {
//			indexesMap[key] = 1
//		}
//
//		i++
//	}
//
//	return centers
//}

func findNextCenter(start h3v3.H3Index, indexesMap map[h3v3.H3Index]int) []h3v3.H3Index {
	result := []h3v3.H3Index{start}

	k2 := h3v3.KRing(start, 2)
	for _, index := range k2 {
		indexesMap[index] = 1
	}

	// k(0) = 1
	// k(1) = 6 + 1 = 7
	// k(2) = 6 * 2 + k(1) = 19
	// k(3) = 6 * 3 + k(2) = 37

	// tìm center index tiếp theo trong k(3) của start
	// 1. đánh dấu các index trong k(2) của start là đã được thăm
	// 2. xét 2 index liền kề trong k(2) của start
	// - lấy các điểm giao k(1) của 2 index này
	// - điểm nào chưa được thăm sẽ là center index tiếp theo
	// 3. tìm được center index: đánh dấu hết k(1), k(2) của điểm này là đã thăm

	for i := 7; i < len(k2); i += 2 {
		intersectIndexes := findIntersectIndex(k2[i], k2[i+1])

		for _, item := range intersectIndexes {
			if _, visited := indexesMap[item]; !visited {
				result = append(result, item)
			}
		}
	}
	return result
}

func findIntersectIndex(index1, index2 h3v3.H3Index) []h3v3.H3Index {
	k1Index1 := h3v3.KRing(index1, 1)
	k1Index2 := h3v3.KRing(index2, 1)
	result := make([]h3v3.H3Index, 0)

	k1Index1Map := make(map[h3v3.H3Index]int)

	for _, index := range k1Index1 {
		k1Index1Map[index] = 1
	}

	for _, index := range k1Index2 {
		if _, found := k1Index1Map[index]; found {
			result = append(result, index)
		}
	}

	return result
}

func getGeoFromIndex(index h3v3.H3Index, originIndexesMap map[string]int, fill string) (GeoFeatures, error) {
	k1 := h3v3.KRing(index, 1)
	neighbors := make([]string, 0)
	for _, item := range k1 {
		if val, found := originIndexesMap[h3v3.ToString(item)]; found && val == 0 {
			neighbors = append(neighbors, h3v3.ToString(item))
			originIndexesMap[h3v3.ToString(item)] = 1
		}
	}

	if len(neighbors) == 0 {
		return GeoFeatures{}, errors.New("neighbors are not found")
	}

	polygons := getCoordinatesFromH3Indexes(neighbors)

	properties := Properties{
		StrokeWidth: "3",
		Stroke:      "#000000",
	}

	if len(fill) > 0 {
		properties.FillColor = fill
	}

	return GeoFeatures{
		Type:       "Feature",
		Properties: properties,
		Geometry: GeoGeometry{
			Type:        "Polygon",
			Coordinates: [][][2]float64{polygons[0]},
		},
	}, nil
}

func GetGeoFromH3Indexes(indexes []h3v3.H3Index, remainIndexes []h3v3.H3Index, originIndexesMap map[string]int) Geo {
	refinedGeo := Geo{
		Type:     "FeatureCollection",
		Features: make([]GeoFeatures, 0),
	}

	for _, index := range remainIndexes {
		k1 := h3v3.KRing(index, 1)
		neighbors := make([]string, 0)
		for _, item := range k1 {
			if val, found := originIndexesMap[h3v3.ToString(item)]; found && val == 0 {
				neighbors = append(neighbors, h3v3.ToString(item))
				originIndexesMap[h3v3.ToString(item)] = 1
			}
		}

		if len(neighbors) == 0 {
			continue
		}

		polygons := getCoordinatesFromH3Indexes(neighbors)

		properties := Properties{
			StrokeWidth: "3",
			Stroke:      "#000000",
		}

		refinedGeo.Features = append(refinedGeo.Features, GeoFeatures{
			Type:       "Feature",
			Properties: properties,
			Geometry: GeoGeometry{
				Type:        "Polygon",
				Coordinates: [][][2]float64{polygons[0]},
			},
		})
	}

	for i, index := range indexes {
		neighbors := h3v3.KRing(index, 1)
		strNeighbors := make([]string, 0)
		for _, n := range neighbors {
			if _, found := originIndexesMap[h3v3.ToString(n)]; found {
				strNeighbors = append(strNeighbors, h3v3.ToString(n))
			}
		}

		if len(strNeighbors) == 0 {
			continue
		}

		polygons := getCoordinatesFromH3Indexes(strNeighbors)

		properties := Properties{
			StrokeWidth: "3",
			Stroke:      "#000000",
		}

		if i == 0 {
			properties.FillColor = colors[0]
		}

		refinedGeo.Features = append(refinedGeo.Features, GeoFeatures{
			Type:       "Feature",
			Properties: properties,
			Geometry: GeoGeometry{
				Type:        "Polygon",
				Coordinates: [][][2]float64{polygons[0]},
			},
		})
	}

	return refinedGeo
}

func GetGeoFromH3Indexes3(indexes []h3v3.H3Index) Geo {
	refinedGeo := Geo{
		Type:     "FeatureCollection",
		Features: make([]GeoFeatures, 0),
	}

	for i, index := range indexes {
		neighbors := h3v3.KRing(index, 1)
		strNeighbors := make([]string, 0)
		for _, n := range neighbors {
			strNeighbors = append(strNeighbors, h3v3.ToString(n))
		}

		polygons := getCoordinatesFromH3Indexes(strNeighbors)

		properties := Properties{
			StrokeWidth: "3",
			Stroke:      "#000000",
		}

		if i == 0 {
			properties.FillColor = colors[0]
		}

		for _, subPolygon := range polygons {
			refinedGeo.Features = append(refinedGeo.Features, GeoFeatures{
				Type:       "Feature",
				Properties: properties,
				Geometry: GeoGeometry{
					Type:        "Polygon",
					Coordinates: [][][2]float64{subPolygon},
				},
			})
			break
		}
	}

	return refinedGeo
}

func getCoordinatesFromH3Indexes(indexes []string) [][][2]float64 {
	h3Indexes := make([]h3v3.H3Index, 0, len(indexes))
	for _, h3IndexStr := range indexes {
		h3Indexes = append(h3Indexes, h3v3.FromString(h3IndexStr))
	}

	polygon := h3v3.SetToLinkedGeo(h3Indexes)
	coordinates := make([][][2]float64, 0)
	for polygon.First != nil {
		loop := polygon.First
		for loop != nil {
			coords := make([][2]float64, 0)
			curr := loop.First
			for curr != nil {
				coordinate := [2]float64{
					curr.Vertex.Longitude,
					curr.Vertex.Latitude,
				}
				coords = append(coords, coordinate)
				next := curr.Next
				curr = next
			}
			coordinates = append(coordinates, coords)
			next := loop.Next
			loop = next
		}
		next := polygon.Next
		if next == nil {
			break
		}
		polygon = *next
	}

	coordinates = append(coordinates, coordinates[0])
	return coordinates
}

func GetGeoFromH3Indexes2() Geo {
	indexes := make([]string, 0)
	dat, err := os.ReadFile("/Users/hau.nguyen1/Documents/code/leetcode/data.txt")
	if err != nil {
		panic(err)
	}

	tmpData := strings.Split(string(dat), ",")
	for _, item := range tmpData {
		indexes = append(indexes, item)
	}

	refinedGeo := Geo{
		Type:     "FeatureCollection",
		Features: make([]GeoFeatures, 0),
	}

	polygons := getCoordinatesFromH3Indexes(indexes)

	properties := Properties{
		StrokeWidth: "3",
		Stroke:      "#000000",
	}

	for _, subPolygon := range polygons {
		refinedGeo.Features = append(refinedGeo.Features, GeoFeatures{
			Type:       "Feature",
			Properties: properties,
			Geometry: GeoGeometry{
				Type:        "Polygon",
				Coordinates: [][][2]float64{subPolygon},
			},
		})
	}

	return refinedGeo
}
