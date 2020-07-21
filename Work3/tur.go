package main

import (
	"bufio"
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"image"
	"image/color"
	"io"
	"os"
	"strings"
)

//Упражнение: Stringers
type IPAddr [4]byte

func (host IPAddr) String() string {
	var sb strings.Builder
	for i, ip := range host {
		if i == len(host)-1 {
			sb.WriteString(fmt.Sprintf("%v", ip))
		} else {
			sb.WriteString(fmt.Sprintf("%v.", ip))
		}
	}
	return sb.String()
}

//Упражнение: ошибки
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.1f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, e
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z * z - x)/(2 * z)
	}
	return z, nil
}

//Упражнение: Reader
type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

//Упражнение: rot13Reader
type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n := len(b)
	if n <= 0 {
		return 0, nil
	}
	tmpReader := bufio.NewReader(rot.r)
	var err error
	for i := 0; i < n; i++ {
		buk, err := tmpReader.ReadByte()
		if err != nil {
			return i, err
		}
		switch {
		case buk >= 65 && buk <= 77:
			buk = buk + 13
		case buk >= 97 && buk <= 109:
			buk = buk + 13
		case buk >= 78 && buk <= 90:
			buk = buk - 13
		case buk >= 110 && buk <= 122:
			buk = buk - 13
		}
		b[i] = buk
	}
	return n, err
}

//Упражнение: изображения
type Image struct{
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		return uint8(x^y)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	//Упражнение: Stringers
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	////Упражнение: ошибки
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	//Упражнение: Reader
	reader.Validate(MyReader{})

	//Упражнение: rot13Reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot13 := rot13Reader{s}
	io.Copy(os.Stdout, &rot13)
	fmt.Println()

	//Упражнение: изображения
	m := Image{64, 64}
	pic.ShowImage(m)
}
