package humanize

import (
	"math"
	"strconv"
)

func stripTrailingZeros(s string) string {
	offset := len(s) - 1
	for offset > 0 {
		if s[offset] == '.' {
			offset--
			break
		}
		if s[offset] != '0' {
			break
		}
		offset--
	}
	return s[:offset+1]
}

// Ftoa converts a float to a string with no trailing zeros.
func Ftoa(num float64) string {
	return stripTrailingZeros(strconv.FormatFloat(num, 'f', 6, 64))
}

// FtoaWidth converts a float to a string with the specified width.
// Can optionally remove trailing zeros.
//
// e.g. FtoaWidth(math.Pi, 4, false) -> "3.14"
func FtoaWidth(num float64, width int, stripTrailing bool) string {
	exponent := math.Log10(num)
	digits := int(math.Floor(exponent)) + 1
	precision := 0
	if digits > 0 {
		// num >= 1
		// Subtract one to leave room for the decimal point.
		precision = width - digits - 1
	} else {
		// num < 1
		// Subtract two to leave room for "0."
		precision = width + digits - 2
	}
	if precision < 0 {
		// Not enough enough precision to display decimals.
		precision = 0
	}
	s := strconv.FormatFloat(num, 'f', precision, 64)
	if stripTrailing {
		s = stripTrailingZeros(s)
	}
	return s
}
