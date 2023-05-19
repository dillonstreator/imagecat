<h1 align="center">imagecat</h1>

<p align="center">Concatenate images into a single image on x or y axis with optional alignment centering</p>

<p align="center">
  <a aria-label="GoDoc" href="https://pkg.go.dev/github.com/dillonstreator/imagecat/v3">
    <img alt="GoDoc badge" src="https://godoc.org/github.com/dillonstreator/go-badge?status.svg">
  </a>
  <a aria-label="GoDoc" href="https://codecov.io/gh/dillonstreator/imagecat">
    <img alt="GoDoc badge" src="https://codecov.io/gh/dillonstreator/imagecat/branch/main/graph/badge.svg?token=ML10BJJUZ6">
  </a>
  <a aria-label="GoReport" href="https://goreportcard.com/report/github.com/dillonstreator/imagecat/v3">
    <img alt="GoReport badge" src="https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat">
  </a>
  <a aria-label="GoReport" href="https://opensource.org/licenses/MIT">
    <img alt="GoReport badge" src="https://img.shields.io/badge/License-MIT-yellow.svg">
  </a>
</p>

## Installation

```sh
go get github.com/dillonstreator/imagecat/v3
```

## Usage

```go
import (
    "image"

    "github.com/dillonstreator/imagecat/v3"
)

func main () {

    images := []image.Image{
        // ... images
    }

    // concat on x-axis with no alignment
    img, err := imagecat.Concat(images)
    if err != nil { /* handle error */ }

    // concat on y-axis with center alignment
    img, err = imagecat.Concat(images, imagecat.WithAxis(imagecat.AxisY), imagecat.WithAlignment(imagecat.AlignmentCenter))
    if err != nil { /* handle error */ }

}
```

## Examples

<div>
    <h3>Input images</h3>
    <p float="left">
      <img src="./resources/img1.jpeg" width="100" />
      <img src="./resources/img2.jpeg" width="100" />
      <img src="./resources/img3.jpeg" width="100" />
    </p>
    <table>
      <tr>
        <th></th>
        <th>Start Alignment</th>
        <th>Center Alignment</th>
        <th>End Alignment</th>
      </tr>
      <tr>
        <th>X-axis</th>
        <td><img src="./resources/result.x.jpeg" width="150" /></td>
        <td><img src="./resources/result.x.center.jpeg" width="150" /></td>
        <td><img src="./resources/result.x.end.jpeg" width="150" /></td>
      </tr>
      <tr>
        <th>Y-axis</th>
        <td align="center"><img src="./resources/result.y.jpeg" height="150" /></td>
        <td align="center"><img src="./resources/result.y.center.jpeg" height="150" /></td>
        <td align="center"><img src="./resources/result.y.end.jpeg" height="150" /></td>
      </tr>
    </table>
</div>


