package scale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScaleImages(t *testing.T) {
	const (
		inDir  = "../testdata/in"
		outDir = "../testdata/out/scale"
	)
	type TestData struct {
		desc      string
		files     []string
		scaleSize int
		expect    []string
	}
	tds := []TestData{
		{
			desc:      "png",
			files:     []string{inDir + "/actor001.png", inDir + "/actor002.png"},
			scaleSize: 100,
			expect:    []string{outDir + "/actor001.png", outDir + "/actor002.png"},
		},
		{
			desc:      "png",
			files:     []string{inDir + "/actor001.png", inDir + "/actor002.png"},
			scaleSize: 200,
			expect:    []string{outDir + "/actor001.png", outDir + "/actor002.png"},
		},
		{
			desc:      "png",
			files:     []string{inDir + "/actor001.png", inDir + "/actor002.png"},
			scaleSize: 50,
			expect:    []string{outDir + "/actor001.png", outDir + "/actor002.png"},
		},
	}
	for _, v := range tds {
		got, err := ScaleImages(outDir, v.files, v.scaleSize)
		assert.Nil(t, err, v.desc)
		var count int
		for _, g := range got {
			for _, e := range v.expect {
				if g == e {
					count++
				}
			}
		}
		assert.Equal(t, len(v.expect), count, v.desc)
	}
}
