package paste

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPasteImages(t *testing.T) {
// 	const (
// 		inDir  = "../testdata/in"
// 		outDir = "../testdata/out/flip"
// 	)
// 	type TestData struct {
// 		desc   string
// 		files  []string
// 		expect []string
// 	}
// 	tds := []TestData{
// 		{
// 			desc:   "png",
// 			files:  []string{inDir + "/actor001.png", inDir + "/actor002.png"},
// 			expect: []string{outDir + "/actor001.png", outDir + "/actor002.png"},
// 		},
// 	}
// 	for _, v := range tds {
// 		got, err := PasteImages(outDir, v.files)
// 		assert.Nil(t, err, v.desc)
// 		var count int
// 		for _, g := range got {
// 			for _, e := range v.expect {
// 				if g == e {
// 					count++
// 				}
// 			}
// 		}
// 		assert.Equal(t, len(v.expect), count, v.desc)
// 	}
// }

func TestCalcPos(t *testing.T) {
	type TestData struct {
		desc             string
		i, r, c, w, h, m int
		expect           image.Point
	}
	tds := []TestData{
		{
			desc: "A:1",
			i:    0, r: 2, c: 4, w: 144, h: 144, m: 8,
			expect: image.Pt(0, 0),
		},
		{
			desc: "B:1",
			i:    1, r: 2, c: 4, w: 144, h: 144, m: 8,
			expect: image.Pt(144, 0),
		},
		{
			desc: "D:1",
			i:    3, r: 2, c: 4, w: 144, h: 144, m: 8,
			expect: image.Pt(432, 0),
		},
		{
			desc: "A:2",
			i:    4, r: 2, c: 4, w: 144, h: 144, m: 8,
			expect: image.Pt(0, 144),
		},
		{
			desc: "D:2",
			i:    7, r: 2, c: 4, w: 144, h: 144, m: 8,
			expect: image.Pt(432, 144),
		},
		{
			desc: "A:1",
			i:    8, r: 2, c: 4, w: 144, h: 144, m: 8,
			expect: image.Pt(0, 0),
		},
		{
			desc: "D:2",
			i:    15, r: 2, c: 4, w: 144, h: 144, m: 8,
			expect: image.Pt(432, 144),
		},
		{
			desc: "D:2",
			i:    15, r: 2, c: 4, w: 10, h: 10, m: 8,
			expect: image.Pt(30, 10),
		},
	}
	for _, v := range tds {
		got := calcPos(v.i, v.r, v.c, v.w, v.h, v.m)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
