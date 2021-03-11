package services

import (
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

// ManipulateImg shifts the color of the image but the value of hueShift
func ManipulateImg(hueShift float64) {

	src, err := imaging.Open("services/testdata/gopher.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img1 := imaging.AdjustHue(src, hueShift)

	dst := imaging.New(src.Bounds().Dx(), src.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))

	err = imaging.Save(dst, "services/testdata/gopher_out.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}
