package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kalebpc/ImageEdit/internal/pkg/args"
	"github.com/kalebpc/ImageEdit/internal/pkg/imageedit"
)

func main() {
	arguments, exit := args.GetArgs()
	if !exit {
		processImage(arguments)
	}
	os.Exit(0)
}

func processImage(arguments args.Arguments) {
	file, err := os.Open(arguments.Infile)
	if err != nil {
		fmt.Println("Error opening infile")
	} else {
		defer file.Close()
		var imageedit imageedit.Imageedit
		imageedit.Oldimg, err = png.Decode(file)
		if err != nil {
			fmt.Println("Cannot decode file")
		} else {
			imageedit.Newimg = image.NewNRGBA(image.Rectangle{image.Point{imageedit.Oldimg.Bounds().Min.X, imageedit.Oldimg.Bounds().Min.Y}, image.Point{imageedit.Oldimg.Bounds().Max.X, imageedit.Oldimg.Bounds().Max.Y}})
			// Apply imageedit method
			switch arguments.Function {
			case "FX":
				// flip over x-axis
				imageedit.FX()
			case "FY":
				// flip over y-axis
				imageedit.FY()
			case "FXY":
				// rotate
				imageedit.FXY()
			case "RRX":
				// roundrobin around x-axis
				imageedit.RRX(arguments.Pixels)
			case "RRY":
				// roundrobin around y-axis
				imageedit.RRY(arguments.Pixels)
			case "RRR":
				// roundrobin `pixels` size rows
				imageedit.RRR(arguments.Pixels)
			case "RRC":
				// roundrobin `pixels` size columns
				imageedit.RRC(arguments.Pixels)
			case "PIX":
				// `pixels` size pixelate whole image
				imageedit.PIX(arguments.Pixels)
			}
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
	return
}
