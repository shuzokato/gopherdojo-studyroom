// 画像変換

package convert

import (
	"bufio"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Convert image
func Convert() {
	fpaths := getFiles(".")
	for _, fpath := range fpaths {
		newFile(fpath)
		buf, file := readImage(fpath)
		defer file.Close()
		img, err := jpeg.Decode(buf)
		var newFileName string
		newFileName = fmt.Sprint(fpath[:len(fpath)-len(filepath.Ext(fpath))], ".png")
		out, err := os.Create(newFileName)
		if err != nil {
			panic(err)
		}
		defer out.Close()
		png.Encode(out, img)
		fmt.Printf("sucessfully converted '%s' to '%s'.\n", fpath, newFileName)
	}
}

// read input image
func readImage(fpath string) (*bufio.Reader, *os.File) {
	file, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(file)
	return buf, file
}

func getFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var fpaths []string
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			fpaths = append(fpaths, getFiles(fpath)...)
			continue
		}
		if filepath.Ext(file.Name()) == ".jpg" {
			fpaths = append(fpaths, fpath)
		}
	}
	return fpaths
}

type myFile struct {
	name     string
	fileType string
}

func newFile(fpath string) *myFile {
	file := new(myFile)
	file.name = filepath.Base(fpath)
	file.fileType = fpath[:len(fpath)-len(filepath.Ext(fpath))]
	fmt.Printf("file detected:%s\n", file.name)
	return file
}
