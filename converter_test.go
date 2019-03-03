package romnum

import "testing"

func TestConverter(t *testing.T) {

	numeralTestData := []numeral{
		0:  numeral{1, "I"},
		1:  numeral{2, "II"},
		2:  numeral{3, "III"},
		3:  numeral{4, "IV"},
		4:  numeral{5, "V"},
		5:  numeral{6, "VI"},
		6:  numeral{9, "IX"},
		7:  numeral{10, "X"},
		8:  numeral{20, "XX"},
		9:  numeral{100, "C"},
		10: numeral{400, "CD"},
		11: numeral{500, "D"},
		12: numeral{3999, "MMMCMXCIX"},
	}

	for _, v := range numeralTestData {
		result, err := Convert(v.Num)
		if err != nil {
			t.Errorf("expected %s, got %s", v.Value, err.Error())
		}
		if result != v.Value {
			t.Errorf("Fail: expected %s, got %s", v.Value, result)
		}
	}

}

func TestIsValid(t *testing.T) {
	if IsValid(0) != false {
		t.Errorf("Expected %v, got %v", false, IsValid(0))
	}
}
