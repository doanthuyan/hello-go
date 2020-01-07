package mypic
import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
	"math/rand"
)
type MyImage struct{
	// Pix holds the image's pixels, in R, G, B, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}
func PicShow(){
	const (
		dx = 100
		dy = 100
	)
	data := CreateData(dx,dy)
	m := NewMyImage(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			if v != 0 {
				m.Pix[i] = uint8(rand.Intn(int(v)))
				m.Pix[i+1] = uint8(rand.Intn(int(v)))
				m.Pix[i+2] = uint8(rand.Intn(int(v)))
				m.Pix[i+3] = uint8(rand.Intn(int(v)))
			}else{
				m.Pix[i] = v
				m.Pix[i+1] = v
				m.Pix[i+2] = v
				m.Pix[i+3] = v
			}
		}
	}
	pic.ShowImage(m)
}
func CreateData(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}
	
	// Do something.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			if i + j > 0{	
				a[i][j] =  uint8(rand.Intn(i+ j))
			} else{
				a[i][j] = uint8(dx)
			}
		}
	}	
	return a
}
func (mi MyImage) Bounds() image.Rectangle{
	
	return mi.Rect
}
func (mi MyImage) ColorModel() color.Model{
	return color.RGBAModel
}
func (mi MyImage) At(x, y int) color.Color{
	return mi.RGBAAt(x, y)
}
func (mi MyImage) RGBAAt(x, y int) color.RGBA {
	if !(image.Point{x, y}.In(mi.Rect)) {
		return color.RGBA{}
	}
	i := mi.PixOffset(x, y)
	s := mi.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	return color.RGBA{s[0], s[1], s[2], s[3]}
}
// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (mi MyImage) PixOffset(x, y int) int {
	return (y-mi.Rect.Min.Y)*mi.Stride + (x-mi.Rect.Min.X)*4
}
// NewRGBA returns a new RGBA image with the given bounds.
func NewMyImage(r image.Rectangle) *MyImage {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 4*w*h)
	return &MyImage{buf, 4 * w, r}
}
// Opaque scans the entire image and reports whether it is fully opaque.
func (mi MyImage) Opaque() bool {
	/*if mi.Rect.Empty() {
		return true
	}
	i0, i1 := 3, mi.Rect.Dx()*4
	for y := mi.Rect.Min.Y; y < mi.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 4 {
			if mi.Pix[i] != 0xff {
				return false
			}
		}
		i0 += mi.Stride
		i1 += mi.Stride
	}
	return true*/
	return false
}
func (mi MyImage) SetRGBA(x, y int, c color.RGBA) {
	if !(image.Point{x, y}.In(mi.Rect)) {
		return
	}
	i := mi.PixOffset(x, y)
	s := mi.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = c.R
	s[1] = c.G
	s[2] = c.B
	s[3] = c.A
}
func (mi MyImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(mi.Rect)) {
		return
	}
	i := mi.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	s := mi.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = c1.R
	s[1] = c1.G
	s[2] = c1.B
	s[3] = c1.A
}