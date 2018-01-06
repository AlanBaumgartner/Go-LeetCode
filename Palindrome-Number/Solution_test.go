package isPalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	test := false
	result := isPalindrome(123)
	if test != result {
		t.Errorf("Failed - Actual: %7v, Result: %7v", test, result)
	} else {
		t.Logf("Passed - Actual: %7v, Result: %7v", test, result)
	}

	test = true
	result = isPalindrome(2155512)
	if test != result {
		t.Errorf("Failed - Actual: %7v, Result: %7v", test, result)
	} else {
		t.Logf("Passed - Actual: %7v, Result: %7v", test, result)
	}

	test = false
	result = isPalindrome(-555)
	if test != result {
		t.Errorf("Failed - Actual: %7v, Result: %7v", test, result)
	} else {
		t.Logf("Passed - Actual: %7v, Result: %7v", test, result)
	}

	test = true
	result = isPalindrome(420024)
	if test != result {
		t.Errorf("Failed - Actual: %7v, Result: %7v", test, result)
	} else {
		t.Logf("Passed - Actual: %7v, Result: %7v", test, result)
	}
}
