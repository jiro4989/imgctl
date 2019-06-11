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

// GenerateImages は画像を重ねて書き込む。
// 画像は指定ディレクトリに連番のファイル名を生成して保存する。
// 戻り値として生成されたファイル名を返す。
func GenerateImages(outDir, saveFileNameFormat string, filePatterns [][]string) ([]string, error) {
	// 出力先ディレクトリの作成
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return nil, err
	}

	// 画像の幅が必要なので先行して1枚だけload
	src, err := imageio.ReadFile(filePatterns[0][0])
	if err != nil {
		return nil, err
	}
	width := src.Bounds().Size().X
	height := src.Bounds().Size().Y
	onFmt := outDir + "/" + saveFileNameFormat
	src = nil

	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	var generated []string
	for i, p := range filePatterns {
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
			mutex.Lock()
			generated = append(generated, on)
			mutex.Unlock()

			<-ch
		}(i, p)
	}

	wg.Wait()

	return generated, nil
}
