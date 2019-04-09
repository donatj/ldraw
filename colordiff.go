package ldraw

import (
	"image/color"
)

//go:generate go run ./cmd/colorgen/main.go -- colors.go
//go:generate go fmt colors.go

// Color gets the color.Color associated with the given BrickColor
func (b BrickColor) Color() color.Color {
	c, ok := BrickColorMap[b]
	if !ok {
		panic("unhandled color")
	}

	return c
}

// NearestBrickColor converts a color.Color to the nearest BrickColor.
//
// The results of this function may be nondeterministic in cases where the
// color distances are exactly the same.
func NearestBrickColor(c color.Color) (b BrickColor) {
	diff := xcdiff(c, b.Color())

	for bb, bc := range BrickColorMap {
		ndiff := xcdiff(c, bc)
		if ndiff < diff {
			b = bb
			diff = ndiff
		}
	}

	return
}

func xcdiff(l, r color.Color) uint64 {
	lr, lg, lb, la := l.RGBA()
	rr, rg, rb, ra := r.RGBA()

	return abs(int64(lr)-int64(rr)) + abs(int64(lg)-int64(rg)) + abs(int64(lb)-int64(rb)) + abs(int64(la)-int64(ra))
}

func abs(n int64) uint64 {
	if n < 0 {
		return uint64(-n)
	}
	return uint64(n)
}
