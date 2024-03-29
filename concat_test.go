package imagecat

import (
	"bytes"
	_ "embed"
	"image"
	"image/draw"
	"image/jpeg"
	"testing"
)

//go:embed resources/img1.jpeg
var img1Bytes []byte

//go:embed resources/img2.jpeg
var img2Bytes []byte

//go:embed resources/img3.jpeg
var img3Bytes []byte

//go:embed resources/result.x.jpeg
var resultXBytes []byte

//go:embed resources/result.y.jpeg
var resultYBytes []byte

//go:embed resources/result.x.center.jpeg
var resultXCenterBytes []byte

//go:embed resources/result.y.center.jpeg
var resultYCenterBytes []byte

//go:embed resources/result.x.end.jpeg
var resultXEndBytes []byte

//go:embed resources/result.y.end.jpeg
var resultYEndBytes []byte

//go:embed resources/result.x.src.jpeg
var resultXSrcBytes []byte

func TestConcat(t *testing.T) {
	img1, err := jpeg.Decode(bytes.NewReader(img1Bytes))
	if err != nil {
		t.Error(err)
	}
	img2, err := jpeg.Decode(bytes.NewReader(img2Bytes))
	if err != nil {
		t.Error(err)
	}
	img3, err := jpeg.Decode(bytes.NewReader(img3Bytes))
	if err != nil {
		t.Error(err)
	}

	type args struct {
		images  []image.Image
		options []OptionFn
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "x-axis",
			args:    args{images: []image.Image{img1, img2, img3}},
			want:    resultXBytes,
			wantErr: false,
		},
		{
			name:    "y-axis",
			args:    args{images: []image.Image{img1, img2, img3}, options: []OptionFn{WithAxis(AxisY)}},
			want:    resultYBytes,
			wantErr: false,
		},
		{
			name:    "x-axis center",
			args:    args{images: []image.Image{img1, img2, img3}, options: []OptionFn{WithAlignment(AlignmentCenter)}},
			want:    resultXCenterBytes,
			wantErr: false,
		},
		{
			name:    "y-axis center",
			args:    args{images: []image.Image{img1, img2, img3}, options: []OptionFn{WithAxis(AxisY), WithAlignment(AlignmentCenter)}},
			want:    resultYCenterBytes,
			wantErr: false,
		},
		{
			name:    "x-axis end",
			args:    args{images: []image.Image{img1, img2, img3}, options: []OptionFn{WithAlignment(AlignmentEnd)}},
			want:    resultXEndBytes,
			wantErr: false,
		},
		{
			name:    "y-axis end",
			args:    args{images: []image.Image{img1, img2, img3}, options: []OptionFn{WithAxis(AxisY), WithAlignment(AlignmentEnd)}},
			want:    resultYEndBytes,
			wantErr: false,
		},
		{
			name:    "invalid axis",
			args:    args{images: []image.Image{img1, img2, img3}, options: []OptionFn{WithAxis(3), WithAlignment(AlignmentCenter)}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "draw op src",
			args:    args{images: []image.Image{img1, img2, img3}, options: []OptionFn{WithDrawOp(draw.Src)}},
			want:    resultXSrcBytes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Concat(tt.args.images, tt.args.options...)
			if err != nil {
				if tt.wantErr {
					if got != nil {
						t.Error("Concat() expected nil result with non-nil err")
					}

					return
				}

				t.Errorf("Concat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var b bytes.Buffer
			err = jpeg.Encode(&b, got, &jpeg.Options{Quality: 100})
			if err != nil {
				t.Error(err)
			}

			if !bytes.Equal(b.Bytes(), tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
