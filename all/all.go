package all

import (
	"github.com/jiro4989/imgctl/config"
	"github.com/jiro4989/imgctl/crop"
	"github.com/jiro4989/imgctl/flip"
	"github.com/jiro4989/imgctl/generate"
	"github.com/jiro4989/imgctl/paste"
	"github.com/jiro4989/imgctl/scale"
)

func RunAll(conf config.Config) (result []string, err error) {
	var generated []string
	{
		var (
			outDir         = conf.Generate.OutDir
			filenameFormat = conf.Generate.OutFileNameFormat
			patterns       = conf.Generate.SrcPatterns
		)

		generated, err = generate.GenerateImages(outDir, filenameFormat, patterns)
		if err != nil {
			panic(err)
		}
	}
	result = append(result, generated...)

	var fliped []string
	{
		outDir := conf.Flip.OutDir
		fliped, err = flip.FlipImages(outDir, generated)
		if err != nil {
			panic(err)
		}
	}
	result = append(result, fliped...)

	var scaledSrc, scaledFliped []string
	{
		var (
			s      = conf.Scale.Size
			outDir = conf.Scale.OutDir
		)
		scaledSrc, err = scale.ScaleImages(outDir+"/src", generated, s)
		if err != nil {
			panic(err)
		}
		scaledFliped, err = scale.ScaleImages(outDir+"/flip", fliped, s)
		if err != nil {
			panic(err)
		}
	}
	result = append(result, scaledSrc...)
	result = append(result, scaledFliped...)

	var cropedSrc, cropedFliped []string
	{
		var (
			x      = conf.Crop.X
			y      = conf.Crop.Y
			w      = conf.Crop.Width
			h      = conf.Crop.Height
			outDir = conf.Crop.OutDir
		)
		cropedSrc, err = crop.CropImages(outDir+"/src", scaledSrc, x, y, w, h)
		if err != nil {
			panic(err)
		}
		cropedFliped, err = crop.CropImages(outDir+"/flip", scaledFliped, x, y, w, h)
		if err != nil {
			panic(err)
		}
	}
	result = append(result, cropedSrc...)
	result = append(result, cropedFliped...)

	var pastedSrc, pastedFliped []string
	{
		var (
			r              = conf.Paste.Row
			c              = conf.Paste.Col
			w              = conf.Paste.Width
			h              = conf.Paste.Height
			outDir         = conf.Paste.OutDir
			filenameFormat = conf.Paste.OutFileNameFormat
		)
		pastedSrc, err = paste.PasteImages(outDir+"/src", filenameFormat, cropedSrc, r, c, w, h)
		if err != nil {
			panic(err)
		}
		pastedFliped, err = paste.PasteImages(outDir+"/flip", filenameFormat, cropedFliped, r, c, w, h)
		if err != nil {
			panic(err)
		}
	}
	result = append(result, pastedSrc...)
	result = append(result, pastedFliped...)

	return
}
