# imagecat

Concatenate any number of images into a single image on x or y axis with optional alignment centering

## Installation

```sh
go get github.com/DillonStreator/imagecat
```

## Usage

```go
import (
    "image"

    "github.com/DillonStreator/imagecat"
)

func main () {

    images := []image.Image{
        // ... some images
    }

    concater := imagecat.NewConcater(images...)

    img, err := concater.Concat(imagecat.WithAxis(imagecat.ConcatAxisY), imagecat.WithAlignment(imagecat.ConcatAlignmentCenter))
    if err != nil { /* handle error */ }

    // ... do something with image

    // ... concat same images on different axis and different alignment
    img, err = concater.Concat(imagecat.WithAxis(imagecat.ConcatAxisX), imagecat.WithAlignment(imagecat.ConcatAlignmentNone))
    if err != nil { /* handle error */ }

    // ... do something with image

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


