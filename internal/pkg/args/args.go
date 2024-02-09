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

type Arguments struct {
	Infile   string
	Outfile  string
	Function string
	Pixels   int
}

func GetArgs() (arguments Arguments, exit bool) {
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
					arguments.Infile = arglist[1]
					if !arguments.validateInfile() {
						exit = true
					}
				case "outfile":
					arguments.Outfile = arglist[1]
					if !arguments.validateOutfile() {
						exit = true
					}
				case "function":
					arguments.Function = arglist[1]
					if !arguments.validateFunction(arglist[1]) {
						exit = true
					}
				case "pixels":
					temp, err := strconv.Atoi(arglist[1])
					if err != nil || temp < 1 || temp > 1000 {
						arguments.Pixels = 1
					} else {
						arguments.Pixels = temp
					}
				}
			}
		}
	}
	if len(arguments.Infile) < 1 || len(arguments.Outfile) < 1 || len(arguments.Function) < 1 || arguments.Pixels < 1 {
		exit = true
	}
	return arguments, exit
}

func (arguments Arguments) validateInfile() bool {
	result := true
	_, err := filepath.EvalSymlinks(arguments.Infile)
	if err != nil {
		result = false
	} else if arguments.Infile == "" {
		result = false
	} else {
		inExt := filepath.Ext(arguments.Infile)
		if inExt != ".png" {
			result = false
		}
	}
	return result
}

func (arguments Arguments) validateOutfile() bool {
	result := true
	if arguments.Outfile == "" {
		result = false
	} else {
		pathout := filepath.Dir(arguments.Outfile)
		if strings.Contains(strconv.QuoteRuneToASCII(os.PathSeparator), pathout) {
			result = false
		} else {
			outExt := filepath.Ext(arguments.Outfile)
			if strings.Compare(outExt, ".png") != 0 {
				result = false
			}
		}
	}
	return result
}

func (arguments Arguments) validateFunction(function string) bool {
	result := true
	functions := reflect.TypeOf(&imageedit.Imageedit{})
	_, result = functions.MethodByName(function)
	return result
}
