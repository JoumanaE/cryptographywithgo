package ccypher

import "log"

func ShiftRune(r rune, shift int) (rune, error) {

	shift %= 26

	// lowercase
	if shift < 0 || shift > 26 {
		log.Fatal("The shift number has to be between 0 and 26")
	}
	if r >= 'a' && r <= 'z' {
		r += rune(shift)
	}
	if r > 'z' {
		r -= 26
	}
	// uppercase
	if r >= 'A' && r <= 'Z' {
		r += rune(shift)
	}
	if r > 'Z' {
		r -= 26
	}
	return r, nil
}
