package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
}

var paletteMapping = map[string]int{
	"black": 0,
	"red":   1,
	"green": 2,
	"blue":  3,
}

func main() {
	fmt.Printf("Enter a number: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	entered := input.Text()
	fmt.Printf("You entered: %s\n", entered)
}

// ask user whet they want and call lissajous
// NOTE: don't forget that user needs to specify file where to write gif to
func CallLissajous() {
	palette := []color.Color{
		color.White,
		color.Black,
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 0, 255, 255},
	}
	paletteMapping := map[string]int{
		"black": 0,
		"red":   1,
		"green": 2,
		"blue":  3,
	}
	fmt.Println("Specify color for the GIF background(can enter a word or a number):")
}

func lissajous(out io.Writer, backgroundColor color.Color, imageColor color.Color) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	palette := []color.Color{backgroundColor, imageColor}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := rand.Intn(len(palette))
			for colorIndex == 0 {
				// Don't let the index be 0
				colorIndex = rand.Intn(len(palette))
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func printPalleteMapping(palleteMapping map[string]int) {
	for color, number := range palleteMapping {
		fmt.Printf("%s(%d), ", color, number) // I know map is unsorted, but I don't want to go into it right now
	}
	fmt.Println()
}

func getColorFromInput(palleteMapping map[string]int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
}
