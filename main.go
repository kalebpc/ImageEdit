package main

import (
	"fmt"
	"image/png"
	"os"
	"reflect"

	"github.com/kalebpc/ImageEdit/internal/pkg/args"
)

func main() {
	arguments, imageedit, exit := args.GetArgs()
	if !exit {
		file, err := os.Open(arguments[0])
		if err != nil {
			fmt.Println("Error opening infile")
		} else {
			defer file.Close()
			imageedit.Oldimg, err = png.Decode(file)
			if err != nil {
				fmt.Println("Cannot decode file")
			} else {
				imageedit.SetNewimg()
				reflect.TypeOf(reflect.ValueOf(&imageedit).MethodByName(arguments[2]).Call([]reflect.Value{}))
				newfile, err := os.Create(arguments[1])
				if err != nil {
					fmt.Println("Error Creating new outfile")
				} else {
					defer newfile.Close()
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
	os.Exit(0)
}
