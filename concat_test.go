package imagecat

import (
	"bytes"
	_ "embed"
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

func TestConcater_Concat(t *testing.T) {

	c := &Concater{}

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

	c.Add(img1, img2, img3)
	imgX, err := c.Concat()
	if err != nil {
		t.Error(err)
	}

	var b bytes.Buffer
	err = jpeg.Encode(&b, imgX, &jpeg.Options{Quality: 100})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(resultXBytes, b.Bytes()) {
		t.Error("x axis bytes not equal")
	}

	imgY, err := c.Concat(WithAxis(ConcatAxisY))
	if err != nil {
		t.Error(err)
	}

	b = bytes.Buffer{}
	err = jpeg.Encode(&b, imgY, &jpeg.Options{Quality: 100})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(resultYBytes, b.Bytes()) {
		t.Error("y axis bytes not equal")
	}

	imgYCenter, err := c.Concat(WithAxis(ConcatAxisY), WithAlignment(ConcatAlignmentCenter))
	if err != nil {
		t.Error(err)
	}

	b = bytes.Buffer{}
	err = jpeg.Encode(&b, imgYCenter, &jpeg.Options{Quality: 100})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(resultYCenterBytes, b.Bytes()) {
		t.Error("y axis centered bytes not equal")
	}

	imgXCenter, err := c.Concat(WithAxis(ConcatAxisX), WithAlignment(ConcatAlignmentCenter))
	if err != nil {
		t.Error(err)
	}

	b = bytes.Buffer{}
	err = jpeg.Encode(&b, imgXCenter, &jpeg.Options{Quality: 100})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(resultXCenterBytes, b.Bytes()) {
		t.Error("x axis centered bytes not equal")
	}
}
