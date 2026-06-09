package tokeniser

import "go/version"

type merges []string
type vocab map[string]uint64

type AddedToken struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	SingleWord  bool   `json:"single_word"`
	Lstrip      bool   `json:"lstrip"`
	Rstrip      bool   `json:"rstrip"`
	Normalized  bool   `json:"normalized"`
	Special     bool   `json:"special"`
}
type NormalizerStep struct {
	Type    string `json:"type"`
	Prepend string `json:"prepend,omitempty"`
	Pattern struct {
		String string `json:"String"`
	} `json:"pattern,omitempty"`
	Content string `json:"content,omitempty"`
}

type Context struct {
	Version     string           `json:"version"`
	// Use pointers so these can be 'nil' if the JSON says 'null'
	Truncation  interface{}      `json:"truncation"` 
	Padding     interface{}      `json:"padding"`
	AddedTokens []AddedToken     `json:"added_tokens"`
	Normalizer  struct {
		Type        string           `json:"type"`
		Normalizers []NormalizerStep `json:"normalizers"`
	} `json:"normalizer"`
    // We will add 'model' (which contains merges/vocab) next
}

func Tokenise(sentence string) (merges, vocab) {
	m := new(merges)
	v := new(vocab)

	words := []byte(sentence)


	for 

}
