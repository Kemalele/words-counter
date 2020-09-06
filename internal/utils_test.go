package internal

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestItob(t *testing.T) {
	testCases := []struct{
		input int
		res []byte
	}{
		{1,[]byte(strconv.Itoa(1))},
		{10,[]byte(strconv.Itoa(10))},
	}

	for _, test := range testCases {
		if res := itob(test.input); !bytes.Equal(res, test.res) {
			t.Fatal(fmt.Sprintf("expected %v, got %v", test.res,res))
		} else {
			t.Log(fmt.Sprintf("ok: %v == %v",test.res,res))
		}
	}
}
