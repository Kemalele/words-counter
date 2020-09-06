package internal

import (
	"bytes"
	)
type word struct {
	Data []byte
	Count int
}

type Words struct {
	Body 		     []word
	mostRecents      []word
	recentLimit 	 int
}
func InitWords(recentLimit int) Words {
	var words Words
	words.recentLimit = recentLimit
	words.mostRecents = make([]word,recentLimit)

	return words
}
func (words *Words) Append(w word) {
	// initializing mostRecents array
	if len(words.mostRecents) != words.recentLimit {
		words.mostRecents = make([]word,words.recentLimit)
	}

	if i := words.exists(w); i != -1 {
		words.Body[i].Count += 1
		words.mostRecent(words.Body[i])
	} else {
		w.Count += 1
		words.Body = append(words.Body, w)
		words.mostRecent(words.Body[len(words.Body) -1 ])
	}

}

func (words *Words) exists(w word) int {
	if len(words.Body) == 1 && bytes.Equal(words.Body[0].Data,w.Data){
		return 0
	}

	// front, back search
	for i, j := 0, len(words.Body) -1; i < j; i, j = i+1, j-1 {
		if bytes.Equal(words.Body[i].Data, w.Data) {
			return i
		} else if bytes.Equal(words.Body[j].Data, w.Data) {
			return j
		}
	}

	return -1
}

func (words *Words) mostRecent(w word) {
	// if input more recent - replace
	for i, word := range words.mostRecents {
		if w.Count > word.Count {
			words.popRepeatingRecent(w,i)
			words.addRecent(w,i)
			return
			// if word not more recent then before - stop
		} else if bytes.Equal(words.mostRecents[i].Data, w.Data) {
			return
		}
	}
}

func (words *Words) addRecent(w word, wIndex int) {
	var temp []word
	temp = append(temp, words.mostRecents[:wIndex]...)
	temp = append(temp,w)

	lastElements := words.mostRecents[wIndex : len(words.mostRecents)]
	temp = append(temp, lastElements...)

	words.mostRecents = temp[:words.recentLimit]
}

func (words *Words) popRepeatingRecent(w word, wIndex int) {
	var temp []word

	for i := wIndex; i < len(words.mostRecents);i++ {
		if bytes.Equal(w.Data, words.mostRecents[i].Data) {
			temp = append(temp,words.mostRecents[:i]...)
			if i + 1 < len(words.mostRecents) {
				lastElements := words.mostRecents[i+1:]
				temp = append(temp, lastElements...)
			}
			words.mostRecents = temp
		}
	}
}
