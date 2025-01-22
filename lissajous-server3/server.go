package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // second color in palette
)

func main() {
	http.HandleFunc("/", handler)
	serverAddress := "localhost:5000"
	fmt.Printf("Listening at http://%s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := map[string]float64{
		"cycles":  float64(5),
		"res":     0.001,
		"size":    float64(100),
		"nframes": float64(64),
		"delay":   float64(8),
	}
	queryParams := r.URL.Query() // map[cycles:[123 1]]
	for queryKey := range queryParams {
		if !paramsContainKey(params, queryKey) {
			fmt.Printf("Unsupported query param: `%s`, skipping it\n", queryKey)
			continue
		}
		// queryParams[queryKey] [123 1]
		// queryParams.Get(queryKey) 123
		queryVals := queryParams[queryKey]
		queryVal := queryVals[len(queryVals)-1]
		floatQuery, err := strconv.ParseFloat(queryVal, 64)
		if err != nil {
			fmt.Printf("Error occurred during parsing of float64 query param `%s` with value: %s\n", queryKey, queryVal)
			continue
		}
		params[queryKey] = floatQuery
	}

	fmt.Println("Generated values from query params:", params)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: int(params["nframes"])}
	phase := 0.0 // phase difference
	for i := 0; i < int(params["nframes"]); i++ {
		rect := image.Rect(0, 0, 2*int(params["size"])+1, 2*int(params["size"])+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < params["cycles"]*2*math.Pi; t += params["res"] {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(params["size"]+x*params["size"]+0.5), int(params["size"]+y*params["size"]+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, int(params["delay"]))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}

func paramsContainKey(params map[string]float64, key string) bool {
	for paramKey := range params {
		if paramKey == key {
			return true
		}
	}
	return false
}
