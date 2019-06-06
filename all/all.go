package all

import (
	"github.com/jiro4989/imgctl/crop"
	"github.com/jiro4989/imgctl/flip"
	"github.com/jiro4989/imgctl/generate"
	"github.com/jiro4989/imgctl/paste"
	"github.com/jiro4989/imgctl/scale"
)

func RunAll() (result []string, err error) {
	outDir := "out"
	filenameFormat := "%03d.png"
	var patterns [][]string

	generated, err := generate.GenerateImages(outDir, filenameFormat, patterns)
	if err != nil {
		panic(err)
	}
	fliped, err := flip.FlipImages(outDir, generated)
	if err != nil {
		panic(err)
	}

	s := 120
	scaledSrc, err := scale.ScaleImages(outDir, generated, s)
	if err != nil {
		panic(err)
	}
	scaledFliped, err := scale.ScaleImages(outDir, fliped, s)
	if err != nil {
		panic(err)
	}

	x, y, w, h := 1, 1, 1, 1
	cropdSrc, err := crop.CropImages(outDir, scaledSrc, x, y, w, h)
	if err != nil {
		panic(err)
	}
	cropdFliped, err := crop.CropImages(outDir, scaledFliped, x, y, w, h)
	if err != nil {
		panic(err)
	}

	r, c := 2, 4
	pastedSrc, err := paste.PasteImages(outDir, filenameFormat, cropdSrc, r, c, w, h)
	if err != nil {
		panic(err)
	}
	pastedFliped, err := paste.PasteImages(outDir, filenameFormat, cropdFliped, r, c, w, h)
	if err != nil {
		panic(err)
	}

	result = append(result, pastedSrc...)
	result = append(result, pastedFliped...)

	return
}
