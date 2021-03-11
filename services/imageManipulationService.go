package services

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/alexander-grube/cryptogopher/utils"
	"github.com/disintegration/imaging"
)

// AdjustHue changes the hue of the image using the shift parameter (measured in degrees) and returns the adjusted image.
// The shift = 0 (or 360 / -360 / etc.) gives the original image.
// The shift = 180 (or -180) corresponds to a 180° degree rotation of the color wheel and thus gives the image with its hue inverted for each pixel.
//
// Examples:
//  dstImage = imaging.AdjustHue(srcImage, 90) // Shift Hue by 90°.
//  dstImage = imaging.AdjustHue(srcImage, -30) // Shift Hue by -30°.
//
func adjustHue(img image.Image, shift float64) *image.NRGBA {
	if math.Mod(shift, 360) == 0 {
		return imaging.Clone(img)
	}

	summand := shift / 360

	return imaging.AdjustFunc(img, func(c color.NRGBA) color.NRGBA {
		h, s, l := utils.RgbToHSL(c.R, c.G, c.B)
		h += summand
		h = math.Mod(h, 1)
		//Adding 1 because Golang's Modulo function behaves differently to similar operators in most other languages.
		if h < 0 {
			h++
		}
		r, g, b := utils.HslToRGB(h, s, l)
		return color.NRGBA{r, g, b, c.A}
	})
}

// ManipulateImg shifts the color of the image but the value of hueShift
func ManipulateImg(hueShift float64) {

	src, err := imaging.Open("services/testdata/gopher.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img1 := adjustHue(src, hueShift)

	dst := imaging.New(src.Bounds().Dx(), src.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))

	err = imaging.Save(dst, "services/testdata/gopher_out.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}
