package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateImages(t *testing.T) {
	const inDir = "../testdata/in"
	const outDir = "../testdata/out"
	const format = "actor%03d.png"
	type TestData struct {
		desc         string
		filePatterns [][]string
		expect       []string
	}
	tds := []TestData{
		{
			desc: "png",
			filePatterns: [][]string{
				{inDir + "/body.png", inDir + "/eyebrows.png", inDir + "/eye.png", inDir + "/mouse.png"},
				{inDir + "/body.png", inDir + "/eyebrows.png", inDir + "/eye.png", inDir + "/mouse.png", inDir + "/red.png"},
			},
			expect: []string{outDir + "/actor001.png", outDir + "/actor002.png"},
		},
	}
	for _, v := range tds {
		got, err := GenerateImages(outDir, format, v.filePatterns)
		assert.Nil(t, err, v.desc)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
