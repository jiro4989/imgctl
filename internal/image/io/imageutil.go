package io

import (
	"image"
	"image/png"
	"log"
	"os"
)

// ReadImage はファイル名を指定するとPNG画像としてよみとって返却します
func ReadImage(fn string) (img image.Image, err error) {
	w, err := os.Open(fn)
	defer w.Close()
	if err != nil {
		log.Println(err)
		return
	}
	return png.Decode(w)
}

// WriteImage は指定したファイルに画像データを書き込みます。
func WriteImage(fn string, img image.Image) (err error) {
	w, err := os.Create(fn)
	defer w.Close()
	if err != nil {
		log.Println(err)
		return
	}

	return png.Encode(w, img)
}
