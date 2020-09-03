package internal

import (
	"fmt"
	"testing"
)

func TestSplitBytes(t *testing.T) {
	testCases := []struct{
		data []byte
		splitter byte
		result [][]byte
	}{
		{[]byte("hello world"),' ',[][]byte{[]byte("hello"),[]byte("world")}},
		{[]byte("Wind ye down there, ye\nprouder, sadder souls!\n"),' ',[][]byte{[]byte("wind"),[]byte("ye"),[]byte("down"),[]byte("there"),[]byte("ye"),[]byte("prouder"),[]byte("sadder"),[]byte("souls")}}}


	for _, test := range testCases {
		if res := SplitBytes(test.data,test.splitter); !isByteMatrixesEqual(res,test.result) {
			t.Fatal(fmt.Sprintf("expected %s got %s", test.result, res))
		}else {
			t.Log(fmt.Sprintf("ok:   %s == %s",test.result, res))
		}
	}
}

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