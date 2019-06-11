package all

import (
	"sort"

	"github.com/jiro4989/imgctl/config"
	"github.com/jiro4989/imgctl/crop"
	"github.com/jiro4989/imgctl/flip"
	"github.com/jiro4989/imgctl/generate"
	"github.com/jiro4989/imgctl/imageio"
	"github.com/jiro4989/imgctl/paste"
	"github.com/jiro4989/imgctl/scale"
)

func RunAll(conf config.Config) (result []string, err error) {
	srcDir := conf.SrcDir
	flipDir := conf.FlipDir

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
		scaledSrc, err = scale.ScaleImages(outDir+"/"+srcDir, generated, s)
		if err != nil {
			panic(err)
		}
		scaledFliped, err = scale.ScaleImages(outDir+"/"+flipDir, fliped, s)
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
		cropedSrc, err = crop.CropImages(outDir+"/"+srcDir, scaledSrc, x, y, w, h)
		if err != nil {
			panic(err)
		}
		flipedX := getFlipedX(scaledFliped, x, w)
		cropedFliped, err = crop.CropImages(outDir+"/"+flipDir, scaledFliped, flipedX, y, w, h)
		if err != nil {
			panic(err)
		}
	}
	sort.Strings(cropedSrc)
	sort.Strings(cropedFliped)
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
		pastedSrc, err = paste.PasteImages(outDir+"/"+srcDir, filenameFormat, cropedSrc, r, c, w, h)
		if err != nil {
			panic(err)
		}
		pastedFliped, err = paste.PasteImages(outDir+"/"+flipDir, filenameFormat, cropedFliped, r, c, w, h)
		if err != nil {
			panic(err)
		}
	}
	result = append(result, pastedSrc...)
	result = append(result, pastedFliped...)

	return
}

func getFlipedX(files []string, x, w int) int {
	if len(files) < 1 {
		return x
	}

	file := files[0]
	img, err := imageio.ReadFile(file)
	if err != nil {
		panic(err)
	}
	imgX := img.Bounds().Size().X
	return imgX - x - w
}
