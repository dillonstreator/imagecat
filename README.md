# imagecat

Concatenate any number of images into a single image on x or y axis with optional alignment centering

## Installation

```sh
go get github.com/dillonstreator/imagecat/v2
```

## Usage

```go
import (
    "image"

    "github.com/dillonstreator/imagecat/v2"
)

func main () {

    images := []image.Image{
        // ... some images
    }

    img, err := imagecat.Concat(images, imagecat.WithAxis(imagecat.AxisY), imagecat.WithAlignment(imagecat.AlignmentCenter))
    if err != nil { /* handle error */ }

    // ... concat same images on different axis and different alignment
    img, err = imagecat.Concat(imagecat.WithAxis(imagecat.AxisX), imagecat.WithAlignment(imagecat.AlignmentNone))
    if err != nil { /* handle error */ }

}
```

## Examples

<div>
    <h3>Input images</h3>
    <img src="./resources/img1.jpeg" />
    <img src="./resources/img2.jpeg" />
    <img src="./resources/img3.jpeg" />
    <h3>concat on x axis with no alignment</h3>
    <img src="./resources/result.x.jpeg" />
    <h3>concat on x axis with center alignment</h3>
    <img src="./resources/result.x.center.jpeg" />
    <h3>concat on y axis with no alignment</h3>
    <img src="./resources/result.y.jpeg" />
    <h3>concat on y axis with center alignment</h3>
    <img src="./resources/result.y.center.jpeg" />
</div>


