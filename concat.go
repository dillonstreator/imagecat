package imagecat

import (
	"fmt"
	"image"
	"image/draw"
	"math"
	"sync"
)

type Concater struct {
	images      []image.Image
	maxWidth    int
	maxHeight   int
	totalWidth  int
	totalHeight int

	mu sync.Mutex
}

//NewConcater creates a new concater
func NewConcater(images ...image.Image) *Concater {
	c := &Concater{}
	c.Add(images...)

	return c
}

//Add adds images to the concater for later concatenation
func (c *Concater) Add(images ...image.Image) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, img := range images {
		c.images = append(c.images, img)

		height := img.Bounds().Dy()
		width := img.Bounds().Dx()

		c.totalHeight += height
		c.totalWidth += width

		if height > c.maxHeight {
			c.maxHeight = height
		}

		if width > c.maxWidth {
			c.maxWidth = width
		}
	}
}

//Clear clears the images and any calculated max/total widths and heights
func (c *Concater) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.images = []image.Image{}
	c.maxHeight = 0
	c.maxWidth = 0
	c.totalHeight = 0
	c.totalWidth = 0
}

//Concat concatenates the added images with axis and alignment options
func (c *Concater) Concat(opts ...opt) (*image.RGBA, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	config := &concatOptions{
		axis:      ConcatAxisX,
		alignment: ConcatAlignmentNone,
	}
	for _, o := range opts {
		o(config)
	}

	var img *image.RGBA

	if config.axis == ConcatAxisX {
		img = image.NewRGBA(image.Rect(0, 0, c.totalWidth, c.maxHeight))
	} else if config.axis == ConcatAxisY {
		img = image.NewRGBA(image.Rect(0, 0, c.maxWidth, c.totalHeight))
	} else {
		return nil, fmt.Errorf("unkown axis %v", config.axis)
	}

	xPos := 0
	yPos := 0

	for _, i := range c.images {
		height := i.Bounds().Dy()
		width := i.Bounds().Dx()

		if config.alignment == ConcatAlignmentCenter {
			if config.axis == ConcatAxisX {
				// pad yPos
				diff := float64(c.maxHeight - height)
				if diff > 0 {
					padding := int(math.Floor(diff / 2))
					yPos += padding
				}
			} else if config.axis == ConcatAxisY {
				// pad xPos
				diff := float64(c.maxWidth - width)
				if diff > 0 {
					padding := int(math.Floor(diff / 2))
					xPos += padding
				}
			}
		}

		x := xPos + width
		y := yPos + height

		r := image.Rect(xPos, yPos, x, y)
		draw.Draw(img, r, i, image.Point{0, 0}, draw.Over)

		if config.axis == ConcatAxisX {
			xPos += i.Bounds().Dx()
			yPos = 0
		} else if config.axis == ConcatAxisY {
			yPos += i.Bounds().Dy()
			xPos = 0
		}
	}

	return img, nil
}
