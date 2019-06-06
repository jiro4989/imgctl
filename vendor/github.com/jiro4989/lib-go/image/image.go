// image は画像を扱う関数を管理します。
package image

import (
	"image"
	"image/png"
	"log"
	"os"
)

// ReadFile はファイル名を指定するとPNG画像としてよみとって返却します
func ReadFile(fn string) (img image.Image, err error) {
	w, err := os.Open(fn)
	defer w.Close()
	if err != nil {
		log.Println(err)
		return
	}
	return png.Decode(w)
}

// WriteFile は指定したファイルに画像データを書き込みます。
func WriteFile(fn string, img image.Image) (err error) {
	w, err := os.Create(fn)
	defer w.Close()
	if err != nil {
		log.Println(err)
		return
	}
	return png.Encode(w, img)
}
