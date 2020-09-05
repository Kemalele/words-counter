package internal

import (
	"testing"
)



func TestIsBytesMatrixEqueal(t *testing.T) {
	testCases := []struct{
		a [][]byte
		b [][]byte
		result bool
	}{
		{[][]byte{{'a','b','c'},{'z','x','c'}},[][]byte{{'a','b','c'},{'z','x','c'}},true},
		{[][]byte{{'o','b','c'},{'z','x','c'}},[][]byte{{'a','b','c'},{'z','x','c'}},false},
		{[][]byte{{'a','b','c'},{'z','x','c'}},[][]byte{{'a','b','c'},{'z','o','c'}},false},
	}

	for _, test := range testCases {
		if res := isByteMatrixesEqual(test.a,test.b); res != test.result {
			t.Fatal("expected", test.result, "got", res)
		}else {
			t.Log("ok:   ",test.result, "==", res)
		}
	}
}