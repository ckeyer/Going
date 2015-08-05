package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	const (
		dx = 300
		dy = 500
	)

	// 需要保存的文件
	imgcounter := 123
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()

	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	c := color.RGBA{uint8(0), uint8(0), uint8(0), 255}

	for i := 0; i < dx; i++ {
		img.Set(80, i, c)
		img.Set(81, i, c)
		img.Set(82, i, c)
	}
	// for y := 0; y < dy; y++ {
	// 	for x := 0; x < dx; x++ {

	// 		if x%8 == 0 {
	// 			// 设置某个点的颜色，依次是 RGBA
	// 			img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
	// 		}
	// 	}
	// }

	// 以PNG格式保存文件
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}

}
