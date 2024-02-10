package args

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/kalebpc/ImageEdit/internal/pkg/imageedit"
)

func GetArgs() (arguments []string, imageedit imageedit.Imageedit, exit bool) {
	exit = false
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "-h") || strings.Contains(arg, "--help") {
			fmt.Println("Usage:\n      ImageEdit [args] infile=[path/filename.png] outfile=[path/filename.png] function=[FX | FY | ...] pixels=[int]\n\nArguments:\n      infile      : path to photo to edit\n      outfile     : path to save new edited photo\n      function   : name of edit function\n                    [FX]   [FY]   [RRC]\n                    [FXY]  [RRY]  [PIX]\n                    [RRX]  [RRR]\n      pixels      : number of pixels to edit\n      help        : print usage instructions\n\nExample:\n      C:/user> ImageEdit infile=./filetoedit.png outfile=./newfilename.png function=RRR pixels=50")
			exit = true
		} else {
			arglist := strings.Split(arg, "=")
			if len(arglist) > 1 {
				switch arglist[0] {
				case "infile":
					arguments = append(arguments, arglist[1])
					if !validateInfile(arglist[1]) {
						exit = true
						break
					}
				case "outfile":
					arguments = append(arguments, arglist[1])
					if !validateOutfile(arglist[1]) {
						exit = true
						break
					}
				case "function":
					arguments = append(arguments, arglist[1])
					if !validateFunction(arglist[1]) {
						exit = true
						break
					}
				case "pixels":
					temp, err := strconv.Atoi(arglist[1])
					if err != nil || temp < 1 || temp > 1000 {
						imageedit.Pixels = 1
					} else {
						imageedit.Pixels = temp
					}
				}
			}
		}
	}
	if len(arguments) < 3 || imageedit.Pixels < 1 {
		exit = true
	}
	return arguments, imageedit, exit
}

func validateInfile(infile string) bool {
	result := true
	_, err := filepath.EvalSymlinks(infile)
	if err != nil {
		result = false
	} else if infile == "" {
		result = false
	} else {
		inExt := filepath.Ext(infile)
		if inExt != ".png" {
			result = false
		}
	}
	return result
}

func validateOutfile(outfile string) bool {
	result := true
	if outfile == "" {
		result = false
	} else {
		pathout := filepath.Dir(outfile)
		if strings.Contains(strconv.QuoteRuneToASCII(os.PathSeparator), pathout) {
			result = false
		} else {
			outExt := filepath.Ext(outfile)
			if strings.Compare(outExt, ".png") != 0 {
				result = false
			}
		}
	}
	return result
}

func validateFunction(function string) bool {
	result := true
	functions := reflect.TypeOf(&imageedit.Imageedit{})
	_, result = functions.MethodByName(function)
	return result
}
