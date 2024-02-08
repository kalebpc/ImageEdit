package args

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kalebpc/ImageEdit/pkg/ui"
)

type Arguments struct {
	Infile   string
	Outfile  string
	Function string
	Pixels   int
}

func GetArgs() (arguments Arguments, exit bool) {
	exit = false
	if len(os.Args[1:]) <= 0 {
		// Start cl ui
		arguments.Infile = ui.StartUi()
		currentdir, _ := filepath.Split(arguments.Infile)
		fmt.Printf("%s", "Enter name of Outfile: ")
		num, err := fmt.Scanf("%s\n", &arguments.Outfile)
		if num > 1 || err != nil {
			exit = true
		} else {
			arguments.Outfile = currentdir + arguments.Outfile
			fmt.Printf("%s", "Enter name of function to run: ")
			num, err := fmt.Scanf("%s\n", &arguments.Function)
			if num > 1 || err != nil || !arguments.validatefunction(arguments.Function) || !arguments.validateOutfile() || !arguments.validateInfile() {
				exit = true
			} else {
				fmt.Printf("%s", "Enter number of pixels: ")
				num, err := fmt.Scanf("%d\n", &arguments.Pixels)
				if err != nil {
					exit = true
				} else if num > 1 {
					arguments.Pixels = 0
				}
			}
		}
	} else {
		// Inspect os.Args
		for _, arg := range os.Args[1:] {
			if strings.Contains(arg, "-h") || strings.Contains(arg, "--help") || strings.Contains(arg, "help") {
				fmt.Println("Usage:\n      ImageEdit [args] infile=[path/filename.png] outfile=[path/filename.png] function=[FX | FY | ...] pixels=[int]\n\nArguments:\n      infile      : path to photo to edit\n      outfile     : path to save new edited photo\n      function   : name of edit function\n                    [FX]   [FY]   [RRC]\n                    [FXY]  [RRY]  [PIX]\n                    [RRX]  [RRR]\n      pixels      : number of pixels to edit\n      help        : print usage instructions\n\nExample:\n      C:/user> ImageEdit infile=./filetoedit.png outfile=./newfilename.png function=RRR pixels=50")
				exit = true
			}
			arg := strings.Split(arg, "=")
			if len(arg) > 1 {
				if strings.Contains(arg[0], "infile") {
					arguments.Infile = arg[1]
					if !arguments.validateInfile() {
						exit = true
					}
				}
				if strings.Contains(arg[0], "outfile") {
					arguments.Outfile = arg[1]
					if !arguments.validateOutfile() {
						exit = true
					}
				}
				if strings.Contains(arg[0], "function") {
					if arg[1] != "" {
						if arguments.validatefunction(arg[1]) {
							arguments.Function = arg[1]
						}
					} else {
						exit = true
					}
				}
				if strings.Contains(arg[0], "pixels") {
					temp, err := strconv.Atoi(arg[1])
					if err != nil {
						arguments.Pixels = 0
					} else {
						if temp > 0 {
							arguments.Pixels = temp
						} else {
							arguments.Pixels = 0
						}
					}
				}
			} else {
				exit = true
			}
		}
	}

	if len(arguments.Function) < 1 || len(arguments.Infile) < 1 || len(arguments.Outfile) < 1 {
		exit = true
	}
	return arguments, exit
}

func (arguments Arguments) validateInfile() bool {
	result := true
	_, err := filepath.EvalSymlinks(arguments.Infile)
	if err != nil {
		result = false
	}
	if arguments.Infile == "" {
		result = false
	}
	inExt := filepath.Ext(arguments.Infile)
	if inExt != ".png" {
		result = false
	}
	return result
}

func (arguments Arguments) validateOutfile() bool {
	result := true
	if arguments.Outfile == "" {
		result = false
	}
	pathout := filepath.Dir(arguments.Outfile)
	if strings.Contains(strconv.QuoteRuneToASCII(os.PathSeparator), pathout) {
		result = false
	}
	outExt := filepath.Ext(arguments.Outfile)
	if strings.Compare(outExt, ".png") != 0 {
		result = false
	}
	// if !strings.Contains(arguments.Outfile, "/") {
	// 	result = false
	// }
	return result
}

func (arguments Arguments) validatefunction(function string) bool {
	result := true
	validarguments := []string{"FX", "FY", "FXY", "RRX", "RRY", "RRR", "RRC", "PIX"}
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
