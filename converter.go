package romnum

import "errors"

type numeral struct {
	Num   int
	Value string
}

var (
	numerals = []numeral{
		0:  numeral{1000, "M"},
		1:  numeral{900, "CM"},
		2:  numeral{500, "D"},
		3:  numeral{100, "C"},
		4:  numeral{90, "XC"},
		5:  numeral{50, "L"},
		6:  numeral{40, "XL"},
		7:  numeral{10, "X"},
		8:  numeral{9, "IX"},
		9:  numeral{5, "V"},
		10: numeral{4, "IV"},
		11: numeral{1, "I"},
	}
)

//IsValid checks if the number is valid in Roman numeral
func IsValid(n int) bool {
	if n <= 0 || n >= 4000 {
		return false
	}

	return true
}

//Converter converts the int value to roman number and returns as string
func Convert(n int) (string, error) {
	var result string

	if !IsValid(n) {
		return "", errors.New("number does not exist in roman numerals")
	}

	for _, v := range numerals {
		for n >= v.Num {
			result += v.Value
			n = n - v.Num
		}
	}

	return result, nil
}
