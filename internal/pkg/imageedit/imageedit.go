package imageedit

import "image"

type Imageedit struct {
	Oldimg image.Image
	Newimg *image.NRGBA
	Pixels int
}

func (imageedit *Imageedit) SetNewimg() {
	imageedit.Newimg = image.NewNRGBA(image.Rectangle{image.Point{imageedit.Oldimg.Bounds().Min.X, imageedit.Oldimg.Bounds().Min.Y}, image.Point{imageedit.Oldimg.Bounds().Max.X, imageedit.Oldimg.Bounds().Max.Y}})
}

func (imageedit *Imageedit) FY() {
	// flip over Y axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.Oldimg.Bounds().Max.X - i - 1
			imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(k, j))
		}
	}
}

func (imageedit *Imageedit) FX() {
	// flip over X axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.Oldimg.Bounds().Max.Y - j - 1
			imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, k))
		}
	}
}

func (imageedit *Imageedit) FXY() {
	// flip over both axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.Oldimg.Bounds().Max.X - i - 1
			l := imageedit.Oldimg.Bounds().Max.Y - j - 1
			imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(k, l))
		}
	}
}

func (imageedit *Imageedit) RRY() {
	// round robin over Y axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			if i < imageedit.Oldimg.Bounds().Min.X+imageedit.Pixels {
				k := imageedit.Oldimg.Bounds().Max.X - imageedit.Pixels + i - 1
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(k, j))
			} else {
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i-imageedit.Pixels, j))
			}
		}
	}
}

func (imageedit *Imageedit) RRX() {
	// round robin over X axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			if j < imageedit.Oldimg.Bounds().Min.Y+imageedit.Pixels {
				l := imageedit.Oldimg.Bounds().Max.Y - imageedit.Pixels + j - 1
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, l))
			} else {
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, j-imageedit.Pixels))
			}
		}
	}
}

func (imageedit *Imageedit) RRR() {
	// round robin every other pixels size over x axis; rows
	counter := 0
	for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
		counter += 1
		if counter <= imageedit.Pixels {
			for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
				// round robin over x axis
				if i < imageedit.Oldimg.Bounds().Min.X+imageedit.Pixels {
					k := imageedit.Oldimg.Bounds().Max.X - imageedit.Pixels + i - 1
					imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(k, j))
				} else {
					imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i-imageedit.Pixels, j))
				}
			}
		} else if counter < imageedit.Pixels*2 {
			for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, j))
			}
		} else {
			for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, j))
			}
			counter = 0
		}
	}
}

func (imageedit *Imageedit) RRC() {
	// round robin every other pixels size over y axis; columns
	counter := 0
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		counter += 1
		if counter < imageedit.Pixels {
			for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
				// round robin over y axis
				if j < imageedit.Oldimg.Bounds().Min.Y+imageedit.Pixels {
					l := imageedit.Oldimg.Bounds().Max.Y - imageedit.Pixels + j - 1
					imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, l))
				} else {
					imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, j-imageedit.Pixels))
				}
			}
		} else if counter < imageedit.Pixels*2 {
			for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, j))
			}
		} else {
			for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, j))
			}
			counter = 0
		}
	}
}

func (imageedit *Imageedit) PIX() {
	for i := imageedit.Newimg.Bounds().Min.X; i < imageedit.Newimg.Bounds().Max.X; i += imageedit.Pixels {
		for j := imageedit.Newimg.Bounds().Min.Y; j < imageedit.Newimg.Bounds().Max.Y; j += imageedit.Pixels {
			sample := imageedit.Oldimg.At(i, j)
			for k := i; k < i+imageedit.Pixels; k += 1 {
				for l := j; l < j+imageedit.Pixels; l += 1 {
					imageedit.Newimg.Set(k, l, sample)
				}
			}
		}
	}
}
