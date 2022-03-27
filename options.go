package imagecat

type ConcatAxis int
type ConcatAlignment int

const (
	ConcatAxisX ConcatAxis = iota
	ConcatAxisY

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
