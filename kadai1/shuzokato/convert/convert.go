// 画像変換

package convert

import (
	"image/jpeg"
	"image/png"
	"os"
)

func Convert() {
	file, err := os.Open("./src/gopher.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)

	out, err := os.Create("./src/gopher.png")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	png.Encode(out, img)
}
