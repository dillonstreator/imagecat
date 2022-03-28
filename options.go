package imagecat

type ConcatAxis int
type ConcatAlignment int

const (
	// Specifies the axis with which the images will be concatenated on
	ConcatAxisX ConcatAxis = iota
	ConcatAxisY

	// Specifies the alignment that will be applied to the images
	ConcatAlignmentNone ConcatAlignment = iota
	ConcatAlignmentCenter
)

type concatOptions struct {
	axis      ConcatAxis
	alignment ConcatAlignment
}

type opt func(*concatOptions)

func WithAxis(axis ConcatAxis) opt {
	return func(cc *concatOptions) {
		cc.axis = axis
	}
}

func WithAlignment(alignment ConcatAlignment) opt {
	return func(cc *concatOptions) {
		cc.alignment = alignment
	}
}
