package ebcdic

import (
	"fmt"
	"strconv"
)

var units = map[string]string{
	"{": "0",
	"A": "1",
	"B": "2",
	"C": "3",
	"D": "4",
	"E": "5",
	"F": "6",
	"G": "7",
	"H": "8",
	"I": "9",
	"}": "0",
	"J": "1",
	"K": "2",
	"L": "3",
	"M": "4",
	"N": "5",
	"O": "6",
	"P": "7",
	"Q": "8",
	"R": "9",
}

var signed = map[string]string{
	"{": "+",
	"A": "+",
	"B": "+",
	"C": "+",
	"D": "+",
	"E": "+",
	"F": "+",
	"G": "+",
	"H": "+",
	"I": "+",
	"}": "-",
	"J": "-",
	"K": "-",
	"L": "-",
	"M": "-",
	"N": "-",
	"O": "-",
	"P": "-",
	"Q": "-",
	"R": "-",
}

func Amount(s string) string {
	f := (float64(number(s)) / 100)
	return fmt.Sprintf("%s%0.2f", signed[sign(s)], f)
}

func number(s string) int {
	d, _ := strconv.Atoi(fullDigits(s))
	return d
}

func fullDigits(s string) string {
	return digits(s) + units[sign(s)]
}

func sign(s string) string {
	return s[len(s)-1:]
}

func digits(s string) string {
	return s[:len(s)-1]
}
