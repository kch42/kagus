package kagus

import (
	"fmt"
)

// RGB0x is an image/color implementation that allows you to write 24-Bit RGB colors using number literals, usually using the notation 0xRRGGBB
type RGB0x uint32

func (rgb RGB0x) String() string {
	return fmt.Sprintf("#%06x", rgb)
}

func (_rgb RGB0x) RGBA() (r, g, b, a uint32) {
	rgb := uint32(_rgb)
	r = (rgb & 0xff0000) >> 8
	g = rgb & 0xff00
	b = (rgb & 0xff) << 8
	a = 0xffff
	return
}
