package ebcdic

import (
	"testing"
)

func TestOnlyLastChar(t *testing.T) {
	amount := "00000000000123450{"
	last := sign(amount)
	if last != "{" {
		t.Error(last)
	}
}

func TestDigits(t *testing.T) {
	amount := "00000000000123450{"
	d := digits(amount)
	if d != "00000000000123450" {
		t.Error(d)
	}
}

func TestFullDigits(t *testing.T) {
	amount := "00000000000123450{"
	d := fullDigits(amount)
	if d != "000000000001234500" {
		t.Error(d)
	}
}

func TestNumber(t *testing.T) {
	amount := "00000000000123450{"
	d := number(amount)
	if d != 1234500 {
		t.Error(d)
	}
}

func TestEbcdicPositiveZero(t *testing.T) {
	amount := "00000000000123450{"
	s := Amount(amount)
	if s != "+12345.00" {
		t.Error(s)
	}
}

func TestEbcdicPositiveOne(t *testing.T) {
	amount := "00000000000123450A"
	s := Amount(amount)
	if s != "+12345.01" {
		t.Error(s)
	}
}

func TestEbcdicNegativeTwo(t *testing.T) {
	amount := "00000000000123450K"
	s := Amount(amount)
	if s != "-12345.02" {
		t.Error(s)
	}
}

// +0         {              C0               7B
// +1         A              C1               41
// +2         B              C2               42
// +3         C              C3               43
// +4         D              C4               44
// +5         E              C5               45
// +6         F              C6               46
// +7         G              C7               47
// +8         H              C8               48
// +9         I              C9               49

// -0         }              D0               7D
// -1         J              D1               4A
// -2         K              D2               4B
// -3         L              D3               4C
// -4         M              D4               4D
// -5         N              D5               4E
// -6         O              D6               4F
// -7         P              D7               50
// -8         Q              D8               51
// -9         R              D9               52
