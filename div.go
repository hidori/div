package calc

import (
	"github.com/pkg/errors"
)

func Div(dividend uint16, divisor uint16) (quotient uint16, remainder uint16, err error) {
	if divisor == 0 {
		return 0, 0, errors.New("divisor must be not equals to 0")
	}

	if divisor == 1 {
		return dividend, 0, nil
	}

	if dividend == 0 {
		return 0, 0, nil
	}

	if dividend == divisor {
		return 1, 0, nil
	}

	x := dividend
	y := divisor
	z := uint16(1)

	var (
		cx int
		cy int
	)

	for cx = 1; cx < 15; cx++ {
		if x>>cx == 0 {
			break
		}
	}

	for cy = 0; cy < cx; cy++ {
		if y&0x08000 != 0 {
			break
		}

		y <<= 1
		z <<= 1
	}

	for i := 0; i <= cy; i++ {
		t := int(x) - int(y)
		if t >= 0 {
			x = uint16(t)
			quotient |= z
		}

		y >>= 1
		z >>= 1
	}

	remainder = x

	return quotient, remainder, nil
}
