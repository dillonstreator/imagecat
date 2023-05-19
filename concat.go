package imagecat

import (
	"fmt"
	"image"
	"image/draw"
	"math"
)

func Concat(images []image.Image, options ...OptionFn) (*image.RGBA, error) {
	config := newConfig(options...)

	totalHeight := 0
	totalWidth := 0
	maxHeight := 0
	maxWidth := 0

	for _, img := range images {
		height := img.Bounds().Dy()
		width := img.Bounds().Dx()

		totalHeight += height
		totalWidth += width

		if height > maxHeight {
			maxHeight = height
		}

		if width > maxWidth {
			maxWidth = width
		}
	}

	var img *image.RGBA

	if config.axis == AxisX {
		img = image.NewRGBA(image.Rect(0, 0, totalWidth, maxHeight))
	} else if config.axis == AxisY {
		img = image.NewRGBA(image.Rect(0, 0, maxWidth, totalHeight))
	} else {
		return nil, fmt.Errorf("unknown axis %v", config.axis)
	}

	xPos := 0
	yPos := 0

	for _, i := range images {
		height := i.Bounds().Dy()
		width := i.Bounds().Dx()

		if config.alignment == AlignmentCenter {
			if config.axis == AxisX {
				// pad yPos
				diff := float64(maxHeight - height)
				if diff > 0 {
					padding := int(math.Floor(diff / 2))
					yPos += padding
				}
			} else if config.axis == AxisY {
				// pad xPos
				diff := float64(maxWidth - width)
				if diff > 0 {
					padding := int(math.Floor(diff / 2))
					xPos += padding
				}
			}
		} else if config.alignment == AlignmentEnd {
			if config.axis == AxisX {
				// pad yPos
				diff := float64(maxHeight - height)
				if diff > 0 {
					padding := int(math.Floor(diff))
					yPos += padding
				}
			} else if config.axis == AxisY {
				// pad xPos
				diff := float64(maxWidth - width)
				if diff > 0 {
					padding := int(math.Floor(diff))
					xPos += padding
				}
			}
		}

		x := xPos + width
		y := yPos + height

		r := image.Rect(xPos, yPos, x, y)
		draw.Draw(img, r, i, image.Point{0, 0}, config.op)

		if config.axis == AxisX {
			xPos += i.Bounds().Dx()
			yPos = 0
		} else if config.axis == AxisY {
			yPos += i.Bounds().Dy()
			xPos = 0
		}
	}

	return img, nil
}
