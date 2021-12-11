package autocomplete

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"go.uber.org/zap"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	shortStringLevenshteinLimit  int = 0
	mediumStringLevenshteinLimit int = 1
	longStringLevenshteinLimit   int = 2

	shortStringThreshold  int = 0
	mediumStringThreshold int = 3
	longStringThreshold   int = 5
)

type Trie struct {
	Root *Node

	LevenshteinScheme    map[int]int
	LevenshteinIntervals []int
	OriginalDict         map[string][]string
	Logger               *zap.Logger
}

type Node struct {
	Children map[rune]*Node
	Word     string
}

type score int

func NewTrie() *Trie {
	trie := new(Trie)
	trie.Root = new(Node)
	trie.Root.Children = make(map[rune]*Node)
	trie.OriginalDict = make(map[string][]string)
	trie.DefaultLevenshtein()

	return trie
}

func (t *Trie) DefaultLevenshtein() *Trie {
	t.LevenshteinScheme = map[int]int{
		shortStringThreshold:  shortStringLevenshteinLimit,
		mediumStringThreshold: mediumStringLevenshteinLimit,
		longStringThreshold:   longStringLevenshteinLimit}
	t.LevenshteinIntervals = []int{longStringThreshold, mediumStringThreshold, shortStringThreshold}
	return t
}

func (t *Trie) Insert(texts ...string) {
	transformer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	for _, txt := range texts {
		t.insert(txt, transformer)
	}
}

func (t *Trie) insert(text string, transformer transform.Transformer) {
	if len(text) == 0 {
		return
	}
	normal, _, err := transform.String(transformer, text)
	if err != nil {
		t.Logger.Error(err.Error())
		return
	}
	normal = strings.ToLower(normal)
	t.OriginalDict[normal] = append(t.OriginalDict[normal], text)
	text = normal
	currentNode := t.Root
	for i, c := range text {
		child, ok := currentNode.Children[c]
		if !ok {
			child = new(Node)
			child.Children = make(map[rune]*Node)
			if i == len(text)-len(string(c)) {
				child.Word = text
			}

			currentNode.Children[c] = child
		}
		currentNode = child
	}
}

func (t *Trie) Search(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	transformer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	text, _, err := transform.String(transformer, text)
	if err != nil {
		t.Logger.Error(err.Error())
		return []string{}
	}
	text = strings.ToLower(text)

	dist := t.maxDist(text)

	collections := make(map[string]score)

	currentRow := make([]int, 0, len(text)+1)
	for i := 0; i <= len(text); i++ {
		currentRow = append(currentRow, i)
	}
	for l, n := range t.Root.Children {
		n.recursiveLevenshteinDistance(collections, l, text, currentRow, dist)
	}

	hits := make([]string, 0, len(collections))
	for key := range collections {
		hits = append(hits, key)
	}
	sort.Slice(hits, func(i, j int) bool {
		switch {
		case collections[hits[i]] != collections[hits[j]]:
			return collections[hits[i]] < collections[hits[j]]
		default:
			return hits[i] < hits[j]
		}
	})

	originals := make([]string, 0, len(hits)*2)
	for _, hit := range hits {
		originals = append(originals, t.OriginalDict[hit]...)
	}
	return originals
}

func (node *Node) recursiveLevenshteinDistance(collection map[string]score, letter rune, text string, previousRow []int, maxDist int) {
	fmt.Println("Letter: " + string(letter))
	columns := len(text)
	currRow := make([]int, 0, len(previousRow))
	currRow = append(currRow, previousRow[0]+1)

	var insertCost, deleteCost, replaceCost int
	for i := 1; i <= columns; i++ {
		insertCost = currRow[i-1] + 1
		deleteCost = previousRow[i] + 1
		if text[i-1] != uint8(letter) {
			replaceCost = previousRow[i-1] + 1
		} else {
			replaceCost = previousRow[i-1]
		}
		currRow = append(currRow, min(min(insertCost, deleteCost), replaceCost))
	}

	if currRow[columns] <= maxDist {
		if node.Word != "" {
			previousScore, ok := collection[node.Word]
			if !ok {
				collection[node.Word] = score(currRow[columns])
			} else {
				if currRow[columns] < int(previousScore) {
					collection[node.Word] = score(currRow[columns])
				}
			}
		}
		node.collectAllDescendentWords(collection, currRow[columns])
	}

	minVal := currRow[columns]
	for _, v := range currRow {
		minVal = min(v, minVal)
	}

	if minVal <= maxDist {
		for l, n := range node.Children {
			fmt.Println("l: "+string(l), n.Word)
			n.recursiveLevenshteinDistance(collection, l, text, currRow, maxDist)
		}
	}
}

func (node *Node) collectAllDescendentWords(collection map[string]score, distance int) {
	for _, n := range node.Children {
		if n.Word != "" {
			previousScore, ok := collection[n.Word]
			if !ok || distance < int(previousScore) {
				fmt.Println(n.Word, distance)
				collection[n.Word] = score(distance)
			}

		}
		n.collectAllDescendentWords(collection, distance)
	}
}

func (t *Trie) maxDist(s string) int {
	runes := []rune(s)
	for _, limit := range t.LevenshteinIntervals {
		if len(runes) >= limit {
			return t.LevenshteinScheme[limit]
		}
	}

	return 0
}
