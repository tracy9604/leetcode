package h3_demo

import (
	"encoding/json"
	"math"

	h3v3 "github.com/uber/h3-go/v3"
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
	FillColor   string `json:"fill-color,omitempty"`
}

type Geo struct {
	Type     string        `json:"type"`
	Features []GeoFeatures `json:"features"`
}

func GetIndex(lat, long float64) []string {
	latLng := h3v3.GeoCoord{Latitude: lat, Longitude: long}
	resolution := 9 // between 0 (biggest cell) and 15 (smallest cell)

	cell := h3v3.FromGeo(latLng, resolution)

	neighbors2 := h3v3.KRing(cell, 2)
	neighbors3 := h3v3.KRing(cell, 3)
	result := make([]string, 0)

	neighbors2Map := make(map[h3v3.H3Index]h3v3.H3Index)
	for i := 1; i < len(neighbors2); i++ {
		if i == len(neighbors2)-1 {
			neighbors2Map[neighbors2[i]] = neighbors2[0]
			break
		}
		neighbors2Map[neighbors2[i]] = neighbors2Map[neighbors2[i+1]]
	}

	// k(0) = 1
	// k(1) = 6 + 1 = 7
	// k(2) = 6 * 2 + k(1) = 19
	// k(3) = 6 * 3 + k(2) = 37

	firstIndex := math.MinInt
	for i := len(neighbors2); i < len(neighbors3); i++ {
		if firstIndex >= 0 {
			break
		}
		k1Neighbors := h3v3.KRing(neighbors3[i], 1)
		k1NeighborsMap := make(map[h3v3.H3Index]h3v3.H3Index)
		for j := 1; j < len(k1Neighbors); j++ {
			if j == len(k1Neighbors)-1 {
				k1NeighborsMap[k1Neighbors[j]] = k1NeighborsMap[k1Neighbors[0]]
				break
			}
			k1NeighborsMap[k1Neighbors[j]] = k1NeighborsMap[k1Neighbors[j+1]]
		}

		for key, value := range k1NeighborsMap {
			if val, found := neighbors2Map[key]; found && value == val {
				result = append(result, h3v3.ToString(neighbors3[i]))
				firstIndex = i
				break
			}
		}
	}

	for i := 2; i <= 6; i++ {
		result = append(result, h3v3.ToString(neighbors3[firstIndex+3]))
	}

	return result
}

func GetStrGeoFromH3Indexes(indexes []string) string {
	g := GetGeoFromH3Indexes(indexes)
	strG, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return string(strG)
}

func GetGeoFromH3Indexes(indexes []string) Geo {

	refinedGeo := Geo{
		Type:     "FeatureCollection",
		Features: make([]GeoFeatures, 0),
	}

	for _, index := range indexes {
		polygons := getCoordinatesFromH3Indexes([]string{index})

		for _, subPolygon := range polygons {
			refinedGeo.Features = append(refinedGeo.Features, GeoFeatures{
				Type: "Feature",
				Properties: Properties{
					StrokeWidth: "3",
					Stroke:      "#000000",
				},
				Geometry: GeoGeometry{
					Type:        "Polygon",
					Coordinates: [][][2]float64{subPolygon},
				},
			})
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
