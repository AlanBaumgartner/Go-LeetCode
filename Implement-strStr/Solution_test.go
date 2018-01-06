package strStr

import "testing"

func Test_strStr(t *testing.T) {
	test := 0
	result := strStr("Hello", "")
	if test != result {
		t.Errorf("Failed - Actual: %3v, Result: %3v", test, result)
	} else {
		t.Logf("Passed - Actual: %3v, Result: %3v", test, result)
	}

	test = 2
	result = strStr("Hello", "ll")
	if test != result {
		t.Errorf("Failed - Actual: %3v, Result: %3v", test, result)
	} else {
		t.Logf("Passed - Actual: %3v, Result: %3v", test, result)
	}

	test = 3
	result = strStr("Hello", "lo")
	if test != result {
		t.Errorf("Failed - Actual: %3v, Result: %3v", test, result)
	} else {
		t.Logf("Passed - Actual: %3v, Result: %3v", test, result)
	}

	test = 4
	result = strStr("Mississippi", "issip")
	if test != result {
		t.Errorf("Failed - Actual: %3v, Result: %3v", test, result)
	} else {
		t.Logf("Passed - Actual: %3v, Result: %3v", test, result)
	}
}
