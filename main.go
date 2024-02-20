package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type Imageedit struct {
	oldimg image.Image
	newimg *image.NRGBA
	pixels int
}

func main() {
	arguments, imageedit, exit := getArgs()
	if !exit {
		processImage(arguments, imageedit)
	}
	os.Exit(0)
}

func processImage(arguments []string, imageedit Imageedit) {
	file, err := os.Open(arguments[0])
	if err != nil {
		fmt.Println("Error opening infile")
		return
	}
	defer file.Close()
	imageedit.oldimg, err = png.Decode(file)
	if err != nil {
		fmt.Println("Cannot decode file")
		return
	}
	imageedit.newimg = image.NewNRGBA(image.Rectangle{image.Point{imageedit.oldimg.Bounds().Min.X, imageedit.oldimg.Bounds().Min.Y}, image.Point{imageedit.oldimg.Bounds().Max.X, imageedit.oldimg.Bounds().Max.Y}})
	reflect.TypeOf(reflect.ValueOf(&imageedit).MethodByName(arguments[2]).Call([]reflect.Value{}))
	newfile, err := os.Create(arguments[1])
	if err != nil {
		fmt.Println("Error Creating new outfile")
		return
	}
	defer newfile.Close()
	err = png.Encode(newfile, imageedit.newimg)
	if err != nil {
		fmt.Println("Error Encoding new image")
		return
	}
	fmt.Println("New Image Created!")
}

func getArgs() ([]string, Imageedit, bool) {
	arguments := []string{"", "", ""}
	exit := false
	var imageedit Imageedit
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "-h") || strings.Contains(arg, "--help") {
			fmt.Println("Usage:\n      ImageEdit [args] infile=[path/filename.png] outfile=[path/filename.png] function=[Flipx | Flipy | ...] pixels=[int]\n\nArguments:\n      infile      : path to photo to edit\n      outfile     : path to save new edited photo\n      function    : name of edit function\n                    [Flipx]        [Flipy]           [Roundrobincolumns]\n                    [Rotate]       [Roundrobiny]     [Pixelate]\n                    [Roundrobinx]  [Roundrobinrows]  [Rgbfilter]\n      pixels      : number of pixels to edit\n      help        : print usage instructions\n\nExample:\n      C:/user> ImageEdit infile=./filetoedit.png outfile=./newfilename.png function=RRR pixels=50")
			exit = true
			return arguments, imageedit, exit
		}
		arglist := strings.Split(arg, "=")
		if len(arglist) > 1 {
			switch arglist[0] {
			case "infile":
				arguments[0] = arglist[1]
				if !validateInfile(arglist[1]) {
					exit = true
					return arguments, imageedit, exit
				}
			case "outfile":
				arguments[1] = arglist[1]
				if !validateOutfile(arglist[1]) {
					exit = true
					return arguments, imageedit, exit
				}
			case "function":
				arguments[2] = arglist[1]
				if !validateFunction(arglist[1]) {
					exit = true
					return arguments, imageedit, exit
				}
			case "pixels":
				temp, err := strconv.Atoi(arglist[1])
				if err != nil || temp < 1 || temp > 1000 {
					imageedit.pixels = 1
				} else {
					imageedit.pixels = temp
				}
			}
		}
	}
	if len(arguments) < 3 || imageedit.pixels < 1 {
		exit = true
	}
	return arguments, imageedit, exit
}

func validateInfile(infile string) bool {
	_, err := filepath.EvalSymlinks(infile)
	if err != nil {
		return false
	}
	if infile == "" {
		return false
	}
	inExt := filepath.Ext(infile)
	if inExt == ".png" {
		return true
	} else {
		return false
	}
}

func validateOutfile(outfile string) bool {
	if outfile == "" {
		return false
	}
	pathout := filepath.Dir(outfile)
	if strings.Contains(strconv.QuoteRuneToASCII(os.PathSeparator), pathout) {
		return false
	}
	outExt := filepath.Ext(outfile)
	if outExt == ".png" {
		return true
	} else {
		return false
	}
}

func validateFunction(function string) bool {
	functions := reflect.TypeOf(&Imageedit{})
	_, result := functions.MethodByName(function)
	return result
}

func (imageedit *Imageedit) Flipy() *Imageedit {
	// flip over Y axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.oldimg.Bounds().Max.X - i - 1
			imageedit.newimg.Set(i, j, imageedit.oldimg.At(k, j))
		}
	}
	return imageedit
}

func (imageedit *Imageedit) Flipx() *Imageedit {
	// flip over X axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.oldimg.Bounds().Max.Y - j - 1
			imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, k))
		}
	}
	return imageedit
}

func (imageedit *Imageedit) Rotate() *Imageedit {
	// flip over both axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.oldimg.Bounds().Max.X - i - 1
			l := imageedit.oldimg.Bounds().Max.Y - j - 1
			imageedit.newimg.Set(i, j, imageedit.oldimg.At(k, l))
		}
	}
	return imageedit
}

func (imageedit *Imageedit) Roundrobiny() *Imageedit {
	// round robin over Y axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			if i < imageedit.oldimg.Bounds().Min.X+imageedit.pixels {
				k := imageedit.oldimg.Bounds().Max.X - imageedit.pixels + i - 1
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(k, j))
			} else {
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i-imageedit.pixels, j))
			}
		}
	}
	return imageedit
}

func (imageedit *Imageedit) Roundrobinx() *Imageedit {
	// round robin over X axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			if j < imageedit.oldimg.Bounds().Min.Y+imageedit.pixels {
				l := imageedit.oldimg.Bounds().Max.Y - imageedit.pixels + j - 1
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, l))
			} else {
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j-imageedit.pixels))
			}
		}
	}
	return imageedit
}

func (imageedit *Imageedit) Roundrobinrows() *Imageedit {
	// round robin every other pixels size over x axis; rows
	counter := 0
	for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
		counter += 1
		if counter <= imageedit.pixels {
			for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
				if i < imageedit.oldimg.Bounds().Min.X+imageedit.pixels {
					k := imageedit.oldimg.Bounds().Max.X - imageedit.pixels + i - 1
					imageedit.newimg.Set(i, j, imageedit.oldimg.At(k, j))
				} else {
					imageedit.newimg.Set(i, j, imageedit.oldimg.At(i-imageedit.pixels, j))
				}
			}
		} else if counter < imageedit.pixels*2 {
			for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j))
			}
		} else {
			for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j))
			}
			counter = 0
		}
	}
	return imageedit
}

func (imageedit *Imageedit) Roundrobincolumns() *Imageedit {
	// round robin every other pixels size over y axis; columns
	counter := 0
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		counter += 1
		if counter < imageedit.pixels {
			for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
				if j < imageedit.oldimg.Bounds().Min.Y+imageedit.pixels {
					l := imageedit.oldimg.Bounds().Max.Y - imageedit.pixels + j - 1
					imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, l))
				} else {
					imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j-imageedit.pixels))
				}
			}
		} else if counter < imageedit.pixels*2 {
			for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j))
			}
		} else {
			for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j))
			}
			counter = 0
		}
	}
	return imageedit
}

func (imageedit *Imageedit) Pixelate() *Imageedit {
	for i := imageedit.newimg.Bounds().Min.X; i < imageedit.newimg.Bounds().Max.X; i += imageedit.pixels {
		for j := imageedit.newimg.Bounds().Min.Y; j < imageedit.newimg.Bounds().Max.Y; j += imageedit.pixels {
			var sample color.Color
			if i < imageedit.newimg.Bounds().Max.X-imageedit.pixels/2 && j < imageedit.newimg.Bounds().Max.Y-imageedit.pixels/2 {
				sample = imageedit.oldimg.At(i+imageedit.pixels/2, j+imageedit.pixels/2)
			} else {
				sample = imageedit.oldimg.At(i, j)
			}
			for k := i; k < i+imageedit.pixels; k += 1 {
				for l := j; l < j+imageedit.pixels; l += 1 {
					imageedit.newimg.Set(k, l, sample)
				}
			}
		}
	}
	return imageedit
}

func (imageedit *Imageedit) Rgbfilter() *Imageedit {
	for i := imageedit.newimg.Bounds().Min.X; i < imageedit.newimg.Bounds().Max.X; i += 1 {
		for j := imageedit.newimg.Bounds().Min.Y; j < imageedit.newimg.Bounds().Max.Y; j += 1 {
			imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j))
			sample := imageedit.newimg.NRGBAAt(i, j)
			// tint towards red
			if imageedit.pixels == 1 {
				sample.R = sample.R + sample.R/4
				sample.G = sample.G - sample.G/4
				sample.B = sample.B - sample.B/4
				// tint towards green
			} else if imageedit.pixels == 2 {
				sample.R = sample.R - sample.R/4
				sample.G = sample.G + sample.G/4
				sample.B = sample.B - sample.B/4
				// tint towards blue
			} else if imageedit.pixels == 3 {
				sample.R = sample.R - sample.R/4
				sample.G = sample.G - sample.G/4
				sample.B = sample.B + sample.B/4
			}
			imageedit.newimg.SetNRGBA(i, j, sample)
		}
	}
	return imageedit
}
