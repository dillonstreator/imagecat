package imagecat

import "image/draw"

type Axis int
type Alignment int

const (
	// Specifies the axis with which the images will be concatenated on
	AxisX Axis = iota
	AxisY

	// Specifies the alignment that will be applied to the images
	AlignmentLeft Alignment = iota
	AlignmentCenter
	AlignmentEnd
)

type config struct {
	axis      Axis
	alignment Alignment
	op        draw.Op
}

func newConfig(options ...OptionFn) *config {
	cfg := &config{
		axis:      AxisX,
		alignment: AlignmentLeft,
		op:        draw.Over,
	}

	for _, option := range options {
		option(cfg)
	}

	return cfg
}

type OptionFn func(*config)

func WithAxis(axis Axis) OptionFn {
	return func(c *config) {
		c.axis = axis
	}
}

func WithAlignment(alignment Alignment) OptionFn {
	return func(c *config) {
		c.alignment = alignment
	}
}

func WithDrawOp(op draw.Op) OptionFn {
	return func(c *config) {
		c.op = op
	}
}
