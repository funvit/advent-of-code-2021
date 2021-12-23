package day20

import (
	"fmt"
	"log"

	"aoc"
)

// Note: if algo[0]==algo[-1] - it is bad algo.

type Point struct {
	X, Y int
}

type Image struct {
	pixels        map[Point]bool
	width, height int
}

func NewImage(w, h int) *Image {
	return &Image{
		pixels: make(map[Point]bool, w*h),
		width:  w,
		height: h,
	}
}

func (s *Image) Width() int {
	return s.width
}

func (s *Image) Height() int {
	return s.height
}

func (s *Image) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *Image) GetSize() Point {
	return Point{X: s.width, Y: s.height}
}

func (s *Image) SetPixel(x, y int, mark bool) {
	if x > s.width {
		panic("x is out of image width")
	}
	if y > s.height {
		panic("y is out of image height")
	}
	s.pixels[Point{X: x, Y: y}] = mark
}

func (s *Image) GetPixel(x, y int) bool {
	return s.pixels[Point{X: x, Y: y}]
}

func (s *Image) Clone() *Image {
	r := NewImage(s.width, s.height)
	for k, v := range s.pixels {
		r.pixels[k] = v
	}
	return r
}

func Part1(lines []string) int {

	image, algo := readInput(lines)
	log.Println("input image size:", image.GetSize())

	image = enhanceImageN(*image, algo, 2)
	log.Println("output image size:", image.GetSize())

	return len(image.pixels)
}

func Part2(lines []string) int {

	image, algo := readInput(lines)
	log.Println("input image size:", image.GetSize())

	image = enhanceImageN(*image, algo, 50)
	log.Println("output image size:", image.GetSize())

	return len(image.pixels)
}

func readInput(lines []string) (image *Image, algo string) {

	const imageLinesStart = 2

	//
	// read input data
	//
	algo = lines[0]
	image = NewImage(len(lines[2]), len(lines)-2)

	for y := 0; y < len(lines[imageLinesStart:]); y++ {
		line := lines[y+imageLinesStart]

		for x := 0; x < len(line); x++ {
			if line[x] != '#' {
				continue
			}
			image.SetPixel(x, y, true)
		}
	}

	return image, algo
}

func enhanceImageN(img Image, algo string, n int) *Image {

	// By default, infinity is dark. But if algo[0] is "on", this gonna lit
	// infinity after first algo apply pass.
	//
	// Now, if algo[-1] is "off", infinity must be dimmed after second.
	//
	// Repeat until done :)
	//
	// Too bad this little thing is not mentioned in the puzzle
	// description...

	for i := 0; i < n; i++ {
		var borderLitFlipFlop = algo[0] == '#' && algo[len(algo)-1] == '.'
		img = *enhanceImage(img, algo, borderLitFlipFlop && i%2 == 1)
	}

	return &img
}

func enhanceImage(img Image, algo string, infinityLit bool) *Image {

	/*
		Scanning:
		- "o" is infinityLit (dark by default)
		- "*" is a calculated point

		[o o o]
		[o * o]
		[o o .]. . . .
		     . . . . .
		     . . . . .
		     . . . . .
		     . . . . .

	*/

	outputImage := NewImage(img.Width()+2, img.Height()+2)

	for x := 0; x < outputImage.Width(); x++ {
		for y := 0; y < outputImage.Height(); y++ {
			idx := calcAlgoIndex(img, Point{X: x - 1, Y: y - 1}, infinityLit)
			m := algo[idx]
			if m == '#' {
				outputImage.SetPixel(x, y, true)
			}
		}
	}

	return outputImage
}

func calcAlgoIndex(img Image, at Point, infinityLit bool) int {

	pts := []Point{
		{at.X - 1, at.Y - 1},
		{at.X + 0, at.Y - 1},
		{at.X + 1, at.Y - 1},

		{at.X - 1, at.Y},
		{at.X + 0, at.Y},
		{at.X + 1, at.Y},

		{at.X - 1, at.Y + 1},
		{at.X + 0, at.Y + 1},
		{at.X + 1, at.Y + 1},
	}

	var r int
	for _, pt := range pts {
		r = r << 1
		if pt.X < 0 || pt.X > img.Width()-1 || pt.Y < 0 || pt.Y > img.Height()-1 {
			// out of bounds, use infinity lit
			if infinityLit {
				r++
			}
			continue
		}
		if img.GetPixel(pt.X, pt.Y) {
			r++
		}
	}

	return r
}

func drawImage(img Image) {

	for y := 0; y < img.Height(); y++ {
		for x := 0; x < img.Width(); x++ {
			if img.GetPixel(x, y) {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func init() {
	aoc.RegisterSolution(20, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
