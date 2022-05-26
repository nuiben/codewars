package kata

import (
	"fmt"
	"strings"
)

func RGB(r, g, b int) string {
	if r < 0 {
		r = 0
	}
	if r > 255 {
		r = 255
	}
	if g < 0 {
		g = 0
	}
	if g > 255 {
		g = 255
	}
	if b < 0 {
		b = 0
	}
	if b > 255 {
		b = 255
	}

	r1 := r / 16
	r2 := r % 16
	r_value := decToHex(r1, r2)

	g1 := g / 16
	g2 := g % 16
	g_value := decToHex(g1, g2)

	b1 := b / 16
	b2 := b % 16
	b_value := decToHex(b1, b2)

	return r_value + g_value + b_value
}

func decToHex(letterOne, letterTwo int) string {

	h := strings.ToUpper(fmt.Sprintf("%x", letterOne))
	i := strings.ToUpper(fmt.Sprintf("%x", letterTwo))

	return h + i

}
