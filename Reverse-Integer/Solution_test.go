package reverse

import "testing"

func Test_reverse(t *testing.T) {
	test := 123
	result := reverse(321)
	if test != result {
		t.Errorf("Failed - Actual: %10v, Result: %10v", test, result)
	} else {
		t.Logf("Passed - Actual: %10v, Result: %10v", test, result)
	}

	test = -5005
	result = reverse(-5005)
	if test != result {
		t.Errorf("Failed - Actual: %10v, Result: %10v", test, result)
	} else {
		t.Logf("Passed - Actual: %10v, Result: %10v", test, result)
	}

	test = 7556655
	result = reverse(5566557)
	if test != result {
		t.Errorf("Failed - Actual: %10v, Result: %10v", test, result)
	} else {
		t.Logf("Passed - Actual: %10v, Result: %10v", test, result)
	}

	test = 1051
	result = reverse(1501)
	if test != result {
		t.Errorf("Failed - Actual: %10v, Result: %10v", test, result)
	} else {
		t.Logf("Passed - Actual: %10v, Result: %10v", test, result)
	}
}
