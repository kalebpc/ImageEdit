package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Arguments struct {
	infile    string
	outfile   string
	functions []string
	pixels    int
}

type Imageedit struct {
	oldimg image.Image
	newimg *image.NRGBA
}

func main() {
	arguments, exit := userInput()
	for !exit {
		file, err := openPng(arguments.infile)
		if err != nil {
			fmt.Println("Error opening infile")
		} else {
			defer file.Close()
			var imageedit Imageedit
			imageedit.oldimg, err = decodePng(file)
			if err != nil {
				fmt.Println("Cannot decode file")
			} else {
				min := image.Point{imageedit.oldimg.Bounds().Min.X, imageedit.oldimg.Bounds().Min.Y}
				max := image.Point{imageedit.oldimg.Bounds().Max.X, imageedit.oldimg.Bounds().Max.Y}
				imageedit.newimg = image.NewNRGBA(image.Rectangle{min, max})

				// Apply imageedit methods
				for i := 0; i < len(arguments.functions); i += 1 {
					switch function := arguments.functions[i]; function {
					case "FX":
						imageedit.FX()
					case "FY":
						imageedit.FY()
					case "FXY":
						imageedit.FXY()
					case "RRX":
						imageedit.RRX(arguments.pixels)
					case "RRY":
						imageedit.RRY(arguments.pixels)
					case "RRR":
						imageedit.RRR(arguments.pixels)
					case "RRC":
						imageedit.RRC(arguments.pixels)
					}
				}

				// create new file
				newfile, err := newImageFile(arguments.outfile)
				if err != nil {
					fmt.Println("Error Creating new outfile")
				} else {
					defer newfile.Close()
					// encode new file
					err = encodePng(newfile, imageedit.newimg)
					if err != nil {
						fmt.Println("Error Encoding new image")
					}
				}
			}
		}
		exit = true
	}
	os.Exit(0)
}

func openPng(Filein string) (file *os.File, err error) {
	file, err = os.Open(Filein)
	return file, err
}

func decodePng(file *os.File) (img image.Image, err error) {
	img, err = png.Decode(file)
	return img, err
}

func newImageFile(Fileout string) (newfile *os.File, err error) {
	newfile, err = os.Create(Fileout)
	return newfile, err
}

func encodePng(newfile *os.File, newimg *image.NRGBA) (err error) {
	err = png.Encode(newfile, newimg)
	return err
}

func userInput() (arguments Arguments, exit bool) {
	// No arguments or too many: Print Usage instructions
	exit = false
	if len(os.Args[1:]) <= 0 || len(os.Args[1:]) > 5 {
		fmt.Println("ImageEdit --help for usage instructions")
		exit = true
	}
	// Inspect os.Args
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "-h") || strings.Contains(arg, "--help") || strings.Contains(arg, "help") {
			fmt.Println("Usage:\n      ImageEdit [args] infile=[path/filename.png] outfile=[path/filename.png] functions=[FX | FY | ...] pixels=[int]\n\nArguments:\n      infile      : path to photo to edit\n      outfile     : path to save new edited photo\n      functions   : name of edit functions\n                    [FX]   [FY]   [RRC]\n                    [FXY]  [RRY]\n                    [RRX]  [RRR]\n      pixels      : number of pixels to edit\n      help        : print usage instructions\n\nExample:\n      C:/user> ImageEdit infile=./filetoedit.png outfile=./newfilename.png functions=RRR pixels=50")
			exit = true
		}
		arg := strings.Split(arg, "=")
		if len(arg) > 1 {
			if strings.Contains(arg[0], "infile") {
				arguments.infile = arg[1]
				if !arguments.validateInfile() {
					exit = true
				}
			}
			if strings.Contains(arg[0], "outfile") {
				arguments.outfile = arg[1]
				if !arguments.validateOutfile() {
					exit = true
				}
			}
			if strings.Contains(arg[0], "functions") {
				if arg[1] != "" {
					funcin := strings.Split(arg[1], "-")
					for i := 0; i < len(funcin); i += 1 {
						if arguments.validatefunctions(funcin[i]) {
							arguments.functions = append(arguments.functions, funcin[i])
						}
					}
				} else {
					exit = true
				}

			}
			if strings.Contains(arg[0], "pixels") {
				temp, err := strconv.Atoi(arg[1])
				if err != nil {
					arguments.pixels = 0
				} else {
					if temp > 0 {
						arguments.pixels = temp
					} else {
						arguments.pixels = 0
					}
				}
			}
		} else {
			exit = true
		}
	}
	if len(arguments.functions) < 1 || len(arguments.infile) < 1 || len(arguments.outfile) < 1 {
		exit = true
	}
	return arguments, exit
}

// Arguments methods
func (arguments Arguments) validateInfile() bool {
	result := true
	_, err := filepath.EvalSymlinks(arguments.infile)
	if err != nil {
		result = false
	}
	if arguments.infile == "" {
		result = false
	}
	inExt := filepath.Ext(arguments.infile)
	if inExt != ".png" {
		result = false
	}
	return result
}

func (arguments Arguments) validateOutfile() bool {
	result := true
	if arguments.outfile == "" {
		result = false
	}
	pathout := filepath.Dir(arguments.outfile)
	if strings.Contains(strconv.QuoteRuneToASCII(os.PathSeparator), pathout) {
		result = false
	}
	outExt := filepath.Ext(arguments.outfile)
	if strings.Compare(outExt, ".png") != 0 {
		result = false
	}
	if !strings.Contains(arguments.outfile, "/") {
		result = false
	}
	return result
}

func (arguments Arguments) validatefunctions(function string) bool {
	result := true
	validarguments := []string{"FX", "FY", "FXY", "RRX", "RRY", "RRR", "RRC"}
	counter := len(validarguments)
	for _, valid := range validarguments {
		if function == valid {
			counter -= 1
		}
	}
	if counter == len(validarguments) {
		result = false
	}
	return result
}

// Imageedit methods
func (imageedit *Imageedit) FY() {
	// flip imageedit.img over Y axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.oldimg.Bounds().Max.X - i - 1
			imageedit.newimg.Set(i, j, imageedit.oldimg.At(k, j))
		}
	}
}

func (imageedit *Imageedit) FX() {
	// flip imageedit.img over X axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.oldimg.Bounds().Max.Y - j - 1
			imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, k))
		}
	}
}

func (imageedit *Imageedit) FXY() {
	// flip imageedit.img over both axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.oldimg.Bounds().Max.X - i - 1
			l := imageedit.oldimg.Bounds().Max.Y - j - 1
			imageedit.newimg.Set(i, j, imageedit.oldimg.At(k, l))
		}
	}
}

func (imageedit *Imageedit) RRY(pixels int) {
	// round robin imageedit.img over Y axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			if i < imageedit.oldimg.Bounds().Min.X+pixels {
				k := imageedit.oldimg.Bounds().Max.X - pixels + i - 1
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(k, j))
			} else {
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i-pixels, j))
			}
		}
	}
}

func (imageedit *Imageedit) RRX(pixels int) {
	// round robin imageedit.img over X axis
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
			if j < imageedit.oldimg.Bounds().Min.Y+pixels {
				l := imageedit.oldimg.Bounds().Max.Y - pixels + j - 1
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, l))
			} else {
				imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j-pixels))
			}
		}
	}
}

func (imageedit *Imageedit) RRR(pixels int) {
	// round robin imageedit.img every other pixels size over x axis; rows
	counter := 0
	for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
		counter += 1
		if counter <= pixels {
			for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
				// round robin over x axis
				if i < imageedit.oldimg.Bounds().Min.X+pixels {
					k := imageedit.oldimg.Bounds().Max.X - pixels + i - 1
					imageedit.newimg.Set(i, j, imageedit.oldimg.At(k, j))
				} else {
					imageedit.newimg.Set(i, j, imageedit.oldimg.At(i-pixels, j))
				}
			}
		} else if counter < pixels*2 {
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
}

func (imageedit *Imageedit) RRC(pixels int) {
	// round robin imageedit.img every other pixels size over y axis; columns
	counter := 0
	for i := imageedit.oldimg.Bounds().Min.X; i < imageedit.oldimg.Bounds().Max.X; i += 1 {
		counter += 1
		if counter < pixels {
			for j := imageedit.oldimg.Bounds().Min.Y; j < imageedit.oldimg.Bounds().Max.Y; j += 1 {
				// round robin over y axis
				if j < imageedit.oldimg.Bounds().Min.Y+pixels {
					l := imageedit.oldimg.Bounds().Max.Y - pixels + j - 1
					imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, l))
				} else {
					imageedit.newimg.Set(i, j, imageedit.oldimg.At(i, j-pixels))
				}
			}
		} else if counter < pixels*2 {
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
}
