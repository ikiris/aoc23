package generic

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func AddMapVals[T Number, T2 comparable](m map[T2]T) T {
	var result T
	for _, n := range m {
		result += n
	}
	return result
}

func SumArray[T Number](m ...T) T {
	var result T
	for _, n := range m {
		result += n
	}
	return result
}

func AvgArray[T Number](nums ...T) T {
	r := float64(SumArray(nums...)) / float64(len(nums))
	return T(r)
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Min[T Number](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

type Trie struct {
	Val  int
	Kids map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		Kids: make(map[rune]*Trie),
	}
}

func (t *Trie) Add(s string, v int) {
	n := t
	for _, r := range s {
		k := n.Kids[r]
		if k == nil {
			k = NewTrie()
			n.Kids[r] = k
		}
		n = k
	}
	n.Val = v
}