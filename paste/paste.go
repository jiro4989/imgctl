package paste

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"log"
	"math"
	"os"

	"github.com/jiro4989/imgctl/imageio"
)

func PasteImages(outDir, saveFileNameFormat string, files []string, row, col, width, height int) ([]string, error) {
	var (
		max       = row * col
		fnFmt     = fmt.Sprintf("%s/%s", outDir, saveFileNameFormat)
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

	dist := image.NewRGBA(image.Rect(0, 0, outWidth, outHeight))

	var pasted []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		inFile := sc.Text()

		src, err := imageio.ReadFile(inFile)
		if err != nil {
			log.Fatal(err)
		}

		pt := calcPos(cnt, row, col, width, height, max)

		// 画像を貼り付け
		rect := image.Rectangle{
			pt,
			image.Pt(pt.X+width, pt.Y+height),
		}
		draw.Draw(dist, rect, src, image.Pt(0, 0), draw.Over)

		// 画像の保存
		if (cnt+1)%max == 0 {
			on := fmt.Sprintf(fnFmt, fcnt)
			if err := imageio.WriteFile(on, dist); err != nil {
				log.Fatal(err)
			}

			dist = image.NewRGBA(image.Rect(0, 0, outWidth, outHeight))
			fcnt++

			pasted = append(pasted, on)
		}

		cnt++
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	// 空のファイルが生成されないようにチェック
	if 0 < cnt%row*col {
		on := fmt.Sprintf(fnFmt, fcnt)
		if err := imageio.WriteFile(on, dist); err != nil {
			log.Fatal(err)
		}
		fmt.Println(on)
	}

	return pasted, nil
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
