package conversion

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func fileWalk(beforeFormat *string, afterFormat *string, directory *string) error {
	pathErr := filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == *beforeFormat {
			fmt.Println(info.Name())
			path := filepath.Join(*directory, info.Name())
			image, err := os.Open(path)
			if err != nil {
				fmt.Println("ファイルが開けませんでした。")
				return err
			}
			defer image.Close()

			m, imageErr := jpeg.Decode(image)
			if imageErr != nil {
				fmt.Println("画像を解析できませんでした。")
				return imageErr
			}
			newPath := filepath.Join(*directory, "newImage.png")
			pngImage, pngErr := os.Create(newPath)
			if pngErr != nil {
				fmt.Println("ファイルを作成できませんでした。")
				return pngErr
			}
			defer pngImage.Close()
			png.Encode(pngImage, m)
			return nil
		}
		return nil
	})

	if pathErr != nil {
		return pathErr
	}
	return nil
}