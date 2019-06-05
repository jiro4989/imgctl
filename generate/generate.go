package generate

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"runtime"
	"sync"

	"github.com/jiro4989/imgctl/imageio"
)

type Config struct {
	SaveFileNameFormat string
	SaveDir            string
}

// WriteImage は画像を重ねて書き込む。
// 画像は指定ディレクトリに連番のファイル名を生成して保存する。
func WriteImage(filePatterns [][]string, config Config) error {
	// 出力先ディレクトリの作成
	outDir := config.SaveDir
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return err
	}

	// 画像の幅が必要なので先行して1枚だけload
	src, err := imageio.ReadFile(config.Pattern[0][0])
	if err != nil {
		return err
	}
	width := src.Bounds().Size().X
	height := src.Bounds().Size().Y
	onFmt := outDir + "/" + config.SaveFileNameFormat
	src = nil

	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup

	for i, p := range cfg.Image.Pattern {
		wg.Add(1)
		dist := image.NewRGBA(image.Rect(0, 0, width, height))

		go func(i int, p []string) {
			defer wg.Done()
			ch <- 1

			for _, f := range p {
				src, err := imageio.ReadFile(f)
				if err != nil {
					panic(err)
				}

				// 画像を貼り付け
				rect := image.Rectangle{
					image.ZP,
					image.Pt(width, height),
				}
				draw.Draw(dist, rect, src, image.ZP, draw.Over)
			}

			on := fmt.Sprintf(onFmt, (i + 1))
			if err := imageio.WriteFile(on, dist); err != nil {
				panic(err)
			}
			fmt.Println(on)

			<-ch
		}(i, p)
	}

	wg.Wait()
}
