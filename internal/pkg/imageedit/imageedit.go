package imageedit

import "image"

type Imageedit struct {
	Oldimg image.Image
	Newimg *image.NRGBA
}

func (imageedit *Imageedit) FY() {
	// flip imageedit.img over Y axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.Oldimg.Bounds().Max.X - i - 1
			imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(k, j))
		}
	}
}

func (imageedit *Imageedit) FX() {
	// flip imageedit.img over X axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.Oldimg.Bounds().Max.Y - j - 1
			imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, k))
		}
	}
}

func (imageedit *Imageedit) FXY() {
	// flip imageedit.img over both axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			k := imageedit.Oldimg.Bounds().Max.X - i - 1
			l := imageedit.Oldimg.Bounds().Max.Y - j - 1
			imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(k, l))
		}
	}
}

func (imageedit *Imageedit) RRY(pixels int) {
	// round robin imageedit.img over Y axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			if i < imageedit.Oldimg.Bounds().Min.X+pixels {
				k := imageedit.Oldimg.Bounds().Max.X - pixels + i - 1
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(k, j))
			} else {
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i-pixels, j))
			}
		}
	}
}

func (imageedit *Imageedit) RRX(pixels int) {
	// round robin imageedit.img over X axis
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
			if j < imageedit.Oldimg.Bounds().Min.Y+pixels {
				l := imageedit.Oldimg.Bounds().Max.Y - pixels + j - 1
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, l))
			} else {
				imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, j-pixels))
			}
		}
	}
}

func (imageedit *Imageedit) RRR(pixels int) {
	// round robin imageedit.img every other pixels size over x axis; rows
	counter := 0
	for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
		counter += 1
		if counter <= pixels {
			for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
				// round robin over x axis
				if i < imageedit.Oldimg.Bounds().Min.X+pixels {
					k := imageedit.Oldimg.Bounds().Max.X - pixels + i - 1
					imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(k, j))
				} else {
					imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i-pixels, j))
				}
			}
		} else if counter < pixels*2 {
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

func (imageedit *Imageedit) RRC(pixels int) {
	// round robin imageedit.img every other pixels size over y axis; columns
	counter := 0
	for i := imageedit.Oldimg.Bounds().Min.X; i < imageedit.Oldimg.Bounds().Max.X; i += 1 {
		counter += 1
		if counter < pixels {
			for j := imageedit.Oldimg.Bounds().Min.Y; j < imageedit.Oldimg.Bounds().Max.Y; j += 1 {
				// round robin over y axis
				if j < imageedit.Oldimg.Bounds().Min.Y+pixels {
					l := imageedit.Oldimg.Bounds().Max.Y - pixels + j - 1
					imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, l))
				} else {
					imageedit.Newimg.Set(i, j, imageedit.Oldimg.At(i, j-pixels))
				}
			}
		} else if counter < pixels*2 {
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

func (imageedit *Imageedit) PIX(pixels int) {
	for i := imageedit.Newimg.Bounds().Min.X; i < imageedit.Newimg.Bounds().Max.X; i += pixels {
		for j := imageedit.Newimg.Bounds().Min.Y; j < imageedit.Newimg.Bounds().Max.Y; j += pixels {
			sample := imageedit.Oldimg.At(i, j)
			for k := i; k < i+pixels; k += 1 {
				for l := j; l < j+pixels; l += 1 {
					imageedit.Newimg.Set(k, l, sample)
				}
			}
		}
	}
}
