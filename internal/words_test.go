package internal

import (
	"fmt"
	"testing"
)

func TestSplitBytes(t *testing.T) {
	testCases := []struct{
		data []byte
		result [][]byte
	}{
		{[]byte("hello world"),[][]byte{[]byte("hello"),[]byte("world")}},
		{[]byte("Wind ye down there, ye\nprouder, sadder souls!\n"),[][]byte{[]byte("wind"),[]byte("ye"),[]byte("down"),[]byte("there"),[]byte("ye"),[]byte("prouder"),[]byte("sadder"),[]byte("souls")}},
		{[]byte("Hello--Kitty"),[][]byte{[]byte("hello"), []byte("kitty")}},
	}



	for _, test := range testCases {
		if res := split(test.data,); !isByteMatrixesEqual(res,test.result) {
			t.Fatal(fmt.Sprintf("expected %s got %s", test.result, res))
		}else {
			t.Log(fmt.Sprintf("ok:   %s == %s",test.result, res))
		}
	}
}
