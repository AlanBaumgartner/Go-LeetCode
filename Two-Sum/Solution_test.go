package twoSum

import "testing"

func Test_twoSum(t *testing.T) {
	pass := true
	test := []int{0, 1}
	result := twoSum([]int{1, 2, 3, 4, 5, 6}, 3)
	if len(test) != len(result) {
		pass = false
	} else {
		for i, v := range test {
			if v != result[i] {
				pass = false
			}
		}
	}
	if pass {
		t.Logf("Passed - Actual: %v, Result: %v", test, result)
	} else {
		t.Errorf("Failed - Actual: %v, Result: %v", test, result)
	}

	pass = true
	test = []int{2, 3}
	result = twoSum([]int{1, 1, 0, 0, 0, 0}, 0)
	if len(test) != len(result) {
		pass = false
	} else {
		for i, v := range test {
			if v != result[i] {
				pass = false
			}
		}
	}
	if pass {
		t.Logf("Passed - Actual: %v, Result: %v", test, result)
	} else {
		t.Errorf("Failed - Actual: %v, Result: %v", test, result)
	}

	pass = true
	test = []int{2, 3}
	result = twoSum([]int{5, 5, 1, 1, 1, 5}, 2)
	if len(test) != len(result) {
		pass = false
	} else {
		for i, v := range test {
			if v != result[i] {
				pass = false
			}
		}
	}
	if pass {
		t.Logf("Passed - Actual: %v, Result: %v", test, result)
	} else {
		t.Errorf("Failed - Actual: %v, Result: %v", test, result)
	}
}
