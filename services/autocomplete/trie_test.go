package autocomplete

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Insert(t *testing.T) {
	text := "Vu Hai Lam"
	trie := NewTrie()
	trie.Insert(text)

	if len(trie.Root.Children) != 1 {
		t.Errorf("Word not inserted correctly, %s", text[0:1])
	}
}

func TestTrie_Search(t *testing.T) {
	testcases := []struct {
		Name     string
		trie     *Trie
		input    []string
		text     string
		expected []string
	}{
		{
			Name:  "Happy case",
			input: []string{"hello", "helloo", "helllloo", "hela", "abcd", "cdef"},
			trie:  NewTrie(),
			text:  "bel",
			expected: []string{
				"hela",
				"helllloo",
				"hello",
				"helloo",
			},
		},
		{
			Name:  "Happy case with utf8",
			input: []string{"vũ hải lâm", "vũ lâm hải"},
			trie:  NewTrie(),
			text:  "vu",
			expected: []string{
				"vũ hải lâm",
				"vũ lâm hải",
			},
		},
		{
			Name:  "Happy case with case insensitive",
			input: []string{"Vũ Hải Lâm", "Vũ Lâm Hải"},
			trie:  NewTrie(),
			text:  "vu",
			expected: []string{
				"Vũ Hải Lâm",
				"Vũ Lâm Hải",
			},
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			test.trie.Insert(test.input...)
			res := test.trie.Search(test.text)
			assert.ElementsMatch(t, res, test.expected)
		})
	}
}
