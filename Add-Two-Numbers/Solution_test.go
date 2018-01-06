package addTwoNumbers

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	pass := true
	test := &ListNode{7, &ListNode{7, &ListNode{7, nil}}}
	result := addTwoNumbers(&ListNode{4, &ListNode{5, &ListNode{6, nil}}}, &ListNode{3, &ListNode{2, &ListNode{1, nil}}})
	for test.Next != nil {
		if test.Val == result.Val {
			test = test.Next
			result = result.Next
		} else {
			pass = false
		}
	}
	if pass {
		t.Logf("Passed")
	} else {
		t.Errorf("Failed")
	}

	pass = true
	test = &ListNode{7, &ListNode{7, &ListNode{7, nil}}}
	result = addTwoNumbers(&ListNode{4, &ListNode{5, &ListNode{6, nil}}}, &ListNode{3, &ListNode{2, &ListNode{1, nil}}})
	for test.Next != nil {
		if test.Val == result.Val {
			test = test.Next
			result = result.Next
		} else {
			pass = false
		}
	}
	if pass {
		t.Logf("Passed")
	} else {
		t.Errorf("Failed")
	}

	pass = true
	test = &ListNode{0, &ListNode{4, &ListNode{0, &ListNode{1, nil}}}}
	result = addTwoNumbers(&ListNode{5, &ListNode{9, &ListNode{1, nil}}}, &ListNode{5, &ListNode{4, &ListNode{8, nil}}})
	for test.Next != nil {
		if test.Val == result.Val {
			test = test.Next
			result = result.Next
		} else {
			pass = false
		}
	}
	if pass {
		t.Logf("Passed")
	} else {
		t.Errorf("Failed")
	}
}
