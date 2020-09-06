package internal

import (
	"fmt"
	"testing"
)
var wordsTemplate = Words {
	[]word {
		{[]byte("hello"),0},
		{[]byte("bye"),0},
	},
	[]word{
		{[]byte("hello"),12},
		{[]byte("salam"),10},
		{[]byte("temporary"),1},
	},
	4,
}

func TestExists(t *testing.T) {
	testCases := []struct{
		words Words
		w 	word
		res int
	}{
		{wordsTemplate,word{[]byte("hello"), 0},0},
		{wordsTemplate,word{[]byte("nope"), 0},-1},
	}

	for _, test := range testCases {
		if i := test.words.exists(test.w); i != test.res {
			t.Fatal(fmt.Println("expected",test.res, "got",i))
		} else {
			t.Log(fmt.Println("ok:",test.res, "==", test.res))
		}
	}
}

func TestAddRecents(t *testing.T) {
	testCases := []struct{
		words Words
		w 	word
		wInd int
		res []word
	}{
		{wordsTemplate,word{[]byte("new"), 11},1,[]word{
			{[]byte("hello"),12},
			{[]byte("new"),11},
			{[]byte("salam"),10},
			{[]byte("temporary"),1},
		}},
	}

	for _, test := range testCases {
		if test.words.addRecent(test.w,test.wInd); !wordsSame(test.words.mostRecents,test.res) {
			t.Fatal(fmt.Sprintf("expected %v got %v",test.res,test.words.mostRecents))
		} else {
			t.Log(fmt.Println("ok:",test.res, "==", test.words.mostRecents))
		}
	}
}

func TestPopRecents(t *testing.T) {
	testCases := []struct {
		words Words
		w 	word
		wIndex int
		res []word
	}{
		{wordsTemplate,word{[]byte("salam"),11},2,[]word {
			{[]byte("hello"),12},
			{[]byte("salam"),11},
			{[]byte("temporary"),1},
		}},
	}

	testCases[0].words.addRecent(word{[]byte("salam"),11},1)

	for _, test := range testCases {
		if test.words.popRepeatingRecent(test.w,test.wIndex); !wordsSame(test.words.mostRecents,test.res) {
			t.Fatal(fmt.Sprintf("expected %v got %v",test.res,test.words.mostRecents))
		} else {
			t.Log(fmt.Println("ok:",test.res, "==", test.words.mostRecents))
		}
	}
}

//func TestSplit(t *testing.T) {
//	testCases := []struct{
//		data []byte
//		result [][]byte
//	}{
//		{[]byte("hello world"),[][]byte{[]byte("hello"),[]byte("world")}},
//		{[]byte("Wind ye down there, ye\nprouder, sadder souls!\n"),[][]byte{[]byte("wind"),[]byte("ye"),[]byte("down"),[]byte("there"),[]byte("ye"),[]byte("prouder"),[]byte("sadder"),[]byte("souls")}},
//		{[]byte("Hello--Kitty"),[][]byte{[]byte("hello"), []byte("kitty")}},
//	}
//
//
//
//	for _, test := range testCases {
//		if res := split(test.data,); {
//			t.Fatal(fmt.Sprintf("expected %s got %s", test.result, res))
//		}else {
//			t.Log(fmt.Sprintf("ok:   %s == %s",test.result, res))
//		}
//	}
//}
