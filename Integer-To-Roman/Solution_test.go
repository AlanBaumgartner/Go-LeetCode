package intToRoman

import "testing"

func Test_intToRoman(t *testing.T) {
	test := "L"
	roman := intToRoman(50)
	if test != roman {
		t.Errorf("Failed - Actual: %10v, Result: %10v", test, roman)
	} else {
		t.Logf("Passed - Actual: %10v, Result: %10v", test, roman)
	}

	test = "XLI"
	roman = intToRoman(41)
	if test != roman {
		t.Errorf("Failed - Actual: %10v, Result: %10v", test, roman)
	} else {
		t.Logf("Passed - Actual: %10v, Result: %10v", test, roman)
	}

	test = "MMMCMXCIX"
	roman = intToRoman(3999)
	if test != roman {
		t.Errorf("Failed - Actual: %10v, Result: %10v", test, roman)
	} else {
		t.Logf("Passed - Actual: %10v, Result: %10v", test, roman)
	}

	test = "CDXCV"
	roman = intToRoman(495)
	if test != roman {
		t.Errorf("Failed - Actual: %10v, Result: %10v", test, roman)
	} else {
		t.Logf("Passed - Actual: %10v, Result: %10v", test, roman)
	}

}
