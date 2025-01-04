// lissajou4 - creates a named .gif file with lissajous figures
// with custom background and primary colors provided by user
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

var colorEscapeCodes = map[string]string{
	"black": "\033[30m",
	"white": "\033[37m",
	"red":   "\033[31m",
	"green": "\033[32m",
	"blue":  "\033[33m",
}

const reset = "\033[0m"

func main() {
	GenerateLissajousGif()
}

// `GenerateLissajousGif` generates a lissajous animation and writes it into a user provided file,
// asks user for the background color and primary color (if they the same, asks user whether
// they are sure about their choice)
func GenerateLissajousGif() {
	var backgroundColor, primaryColor int
	var fileDescriptor *os.File
	fmt.Println("Generate a lissajous gif with custom colors")
	fmt.Println()
	backgroundColor = getColorFromInput(true)
	fmt.Println()
	primaryColor = getColorFromInput(false)
	fmt.Println()
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
	fmt.Println()
	fmt.Println("Your gif was generated!")
}

// `displayColors` displays available colors from the palette
func displayColors() {
	var sortedPalleteMappingKeys []int
	for key := range paletteMapping {
		sortedPalleteMappingKeys = append(sortedPalleteMappingKeys, key)
	}
	sort.Slice(sortedPalleteMappingKeys, func(i, j int) bool {
		return sortedPalleteMappingKeys[i] < sortedPalleteMappingKeys[j]
	})
	for index := range sortedPalleteMappingKeys {
		currColor := paletteMapping[index]
		fmt.Printf("%s%s(%d)%s", colorEscapeCodes[currColor], currColor, index, reset)
		if index == len(palette)-1 {
			fmt.Printf(": ")
		} else {
			fmt.Printf(", ")
		}
	}
}

// `getColorFromInput` gets a user's color choice from the available palette of colors
// bg bool parameter lets user know whether right now they are choosing color for the
// background or primary color
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

// `getFileDescriptorFromInput` gets a filename from user (which can be provided as a relative or full path)
// if file at the provided path already exists - file descriptor for it is returned,
// if not - new file at filename path is created and file descriptor for it is returned
func getFileDescriptorFromInput() *os.File {
	fmt.Printf("Enter a filename(relative or full path) to where gif will be written to(enter 'exit' to exit the program): ")
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

// `lissajous` generates a lissajous animation and writes it to out parameter as a gif
// bgColor(background color) and prColor(primary color) are passed as parameters
// to create a palette for the animation
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
