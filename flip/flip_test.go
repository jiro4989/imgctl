package flip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipImages(t *testing.T) {
	const (
		inDir  = "../testdata/in"
		outDir = "../testdata/out/flip"
	)
	type TestData struct {
		desc   string
		files  []string
		expect []string
	}
	tds := []TestData{
		{
			desc:   "png",
			files:  []string{inDir + "/actor001.png", inDir + "/actor002.png"},
			expect: []string{outDir + "/actor001.png", outDir + "/actor002.png"},
		},
	}
	for _, v := range tds {
		got, err := FlipImages(outDir, v.files)
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
