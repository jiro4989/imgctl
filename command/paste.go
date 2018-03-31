package command

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/codegangsta/cli"
)

func CmdPaste(c *cli.Context) {
	var (
		row       = c.Int("r")
		col       = c.Int("c")
		width     = c.Int("width")
		height    = c.Int("height")
		outPre    = c.String("o")
		padFmt    = c.String("p")
		outDir    = c.String("d")
		max       = row * col
		fnFmt     = fmt.Sprintf("%s/%s%s.png", outDir, outPre, padFmt)
		outWidth  = width * col
		outHeight = height * row
	)
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	var (
		cnt  = 0
		fcnt = 1
	)

	oImg := image.NewRGBA(image.Rect(0, 0, outWidth, outHeight))

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		inFile := sc.Text()

		func() {
			r, err := os.Open(inFile)
			defer r.Close()
			if err != nil {
				log.Fatal(err)
			}

			img, err := png.Decode(r)
			if err != nil {
				log.Fatal(err)
			}

			pt := calcPos(cnt, row, col, width, height, max)

			// 画像を貼り付け
			rect := image.Rectangle{
				pt,
				image.Pt(pt.X+width, pt.Y+height),
			}
			draw.Draw(oImg, rect, img, image.Pt(0, 0), draw.Over)
		}()

		// 画像の保存
		if (cnt+1)%max == 0 {
			on := fmt.Sprintf(fnFmt, fcnt)
			func() {
				w, err := os.Create(on)
				defer w.Close()
				if err != nil {
					log.Fatal(err)
				}
				if err := png.Encode(w, oImg); err != nil {
					log.Fatal(err)
				}
			}()

			oImg = image.NewRGBA(image.Rect(0, 0, outWidth, outHeight))
			fcnt++

			fmt.Println(on)
		}

		cnt++
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	on := fmt.Sprintf(fnFmt, fcnt)
	w, err := os.Create(on)
	defer w.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(w, oImg); err != nil {
		log.Fatal(err)
	}
}

// calcPos は画像の貼り付け座標を計算して返す。
func calcPos(i, r, c, w, h, m int) image.Point {
	var (
		ti = float64(i)
		tc = float64(c)
		tw = float64(w)
		th = float64(h)
		tm = float64(m)
	)

	if m <= i {
		a, _ := math.Modf(ti / tm)
		ti -= a * tm
	}

	x := math.Mod(ti, tc) * tw
	y, _ := math.Modf(ti / tc)
	y *= th

	return image.Pt(int(x), int(y))
}
