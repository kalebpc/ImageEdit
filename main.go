package main

import (
	"fmt"
	"image/png"
	"os"
	"reflect"

	"github.com/kalebpc/ImageEdit/internal/pkg/args"
	"github.com/kalebpc/ImageEdit/internal/pkg/imageedit"
)

func main() {
	arguments, imageedit, exit := args.GetArgs()
	if !exit {
		processImage(arguments, imageedit)
	}
	os.Exit(0)
}

func processImage(arguments args.Arguments, imageedit imageedit.Imageedit) {
	file, err := os.Open(arguments.Infile)
	if err != nil {
		fmt.Println("Error opening infile")
	} else {
		defer file.Close()
		imageedit.Oldimg, err = png.Decode(file)
		if err != nil {
			fmt.Println("Cannot decode file")
		} else {
			imageedit.SetNewimg()
			reflect.TypeOf(reflect.ValueOf(&imageedit).MethodByName(arguments.Function).Call([]reflect.Value{}))
			// create new file
			newfile, err := os.Create(arguments.Outfile)
			if err != nil {
				fmt.Println("Error Creating new outfile")
			} else {
				defer newfile.Close()
				// encode new file
				err = png.Encode(newfile, imageedit.Newimg)
				if err != nil {
					fmt.Println("Error Encoding new image")
				} else {
					fmt.Println("New Image Created!")
				}
			}
		}
	}
}
