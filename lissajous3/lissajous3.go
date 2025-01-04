package main

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

var palette = []color.Color{
	color.Black,
	color.White,
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
}

var paletteMapping = map[int]string{
	0: "black",
	1: "white",
	2: "red",
	3: "green",
	4: "blue",
}

func main() {
	GenerateLissajousGif()
}

func GenerateLissajousGif() {
	var backgroundColor, primaryColor int
	var fileDescriptor *os.File
	fmt.Println("Generate a lissajous gif with custom colors")
	fmt.Println()
	backgroundColor = getColorFromInput(true)
	fmt.Println()
	primaryColor = getColorFromInput(false)
	fmt.Println()
	// fmt.Printf("You chose %s and %s for the bg and pr, respectively\n", paletteMapping[backgroundColor], paletteMapping[primaryColor])
	for backgroundColor == primaryColor {
		var userChoice string // yes or no
		validChoices := []string{"yes", "y", "no", "n"}
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Printf("Background and Primary colors are the same, are you sure?%v ", validChoices)
			scanner.Scan()
			userChoice = strings.ToLower(scanner.Text())
			isChoiceValid := false
			for _, choice := range validChoices {
				if choice == userChoice {
					isChoiceValid = true
					break
				}
			}
			if !isChoiceValid {
				fmt.Printf("Invalid choice: %s\n", userChoice)
				continue
			}
			break
		}
		if userChoice == "yes" || userChoice == "y" {
			fmt.Println("Bg and Pr colors will be the same")
			break
		} else {
			primaryColor = getColorFromInput(false)
		}
	}
	fmt.Printf("You chose %s and %s for the bg and pr, respectively\n", paletteMapping[backgroundColor], paletteMapping[primaryColor])
	fmt.Println()
	fileDescriptor = getFileDescriptorFromInput()
	lissajous(fileDescriptor, palette[backgroundColor], palette[primaryColor])
}

func displayColors() {
	var sortedPalleteMappingKeys []int
	for key := range paletteMapping {
		sortedPalleteMappingKeys = append(sortedPalleteMappingKeys, key)
	}
	sort.Slice(sortedPalleteMappingKeys, func(i, j int) bool {
		return sortedPalleteMappingKeys[i] < sortedPalleteMappingKeys[j]
	})
	for index := range sortedPalleteMappingKeys {
		fmt.Printf("%s(%d)", paletteMapping[index], index)
		if index == len(palette)-1 {
			fmt.Printf(": ")
		} else {
			fmt.Printf(", ")
		}
	}
}

func getColorFromInput(bg bool) int {
	var context string
	if bg {
		context = "background"
	} else {
		context = "primary"
	}
	fmt.Printf("Choose %s color of the image (enter a number or color name, enter 'exit' to exit the program)\n", context)
	displayColors()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		if input == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		intInput, err := strconv.Atoi(input)
		if err != nil { // input is a string
			for key, val := range paletteMapping {
				if val == input {
					fmt.Printf("You chose %s(%d) %s color\n", val, key, context)
					return key
				}
			}
			fmt.Printf("Invalid choice: %s, choose from: ", input)
			displayColors()
		} else { // input is a digit
			for key, val := range paletteMapping {
				if key == intInput {
					fmt.Printf("You chose %s(%d) %s\n", val, key, context)
					return key
				}
			}
			fmt.Printf("Invalid choice: %d, choose from: ", intInput)
			displayColors()
		}
	}
}

func getFileDescriptorFromInput() *os.File {
	fmt.Printf("Enter a filename(relative or full path) to where gif will be written to(filename must end with .gif): ")
	scanner := bufio.NewScanner(os.Stdin)
	var fileDescriptor *os.File
	var filename string
	for {
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		filenames := strings.Split(input, " ")
		filename = filenames[0]
		if len(filenames) > 1 {
			fmt.Printf("Taking first filename %s, discarding others: %s\n", filename, strings.Join(filenames[1:], " "))
		}
		if filename == "" {
			fmt.Printf("Invalid filename, filename cannot be empty: ")
			continue
		}
		splittedFilename := strings.Split(filename, ".")
		if splittedFilename[len(splittedFilename)-1] != "gif" { // even if filename didn't have any dots at all, this condition is still valid
			fmt.Printf("Invalid filename '%s', filename must end with .gif: ", filename)
			continue
		}
		break
	}
	if _, err := os.Stat(filename); err == nil { // actually, apparently, os.Create can work with an existing file too, but it is useful to know about os.Stat
		if fileDescriptor, err = os.Open(filename); err != nil {
			fmt.Fprintf(os.Stderr, "lissajous3: %v\n", err)
		}
		fmt.Printf("Using existing file: %s\n", filename)
	} else if errors.Is(err, os.ErrNotExist) {
		if fileDescriptor, err = os.Create(filename); err != nil {
			fmt.Fprintf(os.Stderr, "lissajous3: %v\n", err)
		}
		fmt.Printf("Creating a new file: %s\n", filename)
	} else {
		fmt.Fprintf(os.Stderr, "lissajous3: %v\n", err)
	}
	return fileDescriptor
}

func lissajous(out io.Writer, bgColor color.Color, prColor color.Color) {
	palette := []color.Color{bgColor, prColor}
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1) // 1 is the index of primary color in the palette
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
