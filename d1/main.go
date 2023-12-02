package main

import (
	"bufio"
	"io"

	"github.com/ikiris/aoc23/generic"
)

type tracker struct {
	tracks []*generic.Trie
	trie   *generic.Trie
}

var listA = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var listB = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func NewTracker(t *generic.Trie) *tracker {
	return &tracker{
		trie: t,
	}
}

func p1(r io.Reader) (int64, error) {
	t := NewTracker(InitTrie(listA))
	s := bufio.NewScanner(r)
	var vals []int
	for s.Scan() {
		first, last := 0, 0
		for _, v := range s.Text() {
			i, ok := t.Get(v)
			if !ok {
				continue
			}
			if first == 0 {
				first = i
			}
			last = i
		}
		vals = append(vals, first*10+last)
	}
	ttl := int64(generic.SumArray(vals...))
	return ttl, nil
}

func p2(r io.Reader) (int64, error) {
	t := NewTracker(InitTrie(listB))
	s := bufio.NewScanner(r)
	var vals []int
	for s.Scan() {
		first, last := 0, 0
		for _, v := range s.Text() {
			i, ok := t.Get(v)
			if !ok {
				continue
			}
			if first == 0 {
				first = i
			}
			last = i
		}
		vals = append(vals, first*10+last)
	}
	ttl := int64(generic.SumArray(vals...))
	return ttl, nil
}

func main() {}

func InitTrie(list map[string]int) *generic.Trie {
	t := generic.NewTrie()
	for s, v := range list {
		t.Add(s, v)
	}
	return t
}

func (t *tracker) Get(r rune) (int, bool) {
	tracks := append(t.tracks, t.trie)
	var newTracks []*generic.Trie
	res, resBool := 0, false
	for _, v := range tracks {
		tr, ok := v.Kids[r]
		if !ok {
			continue
		}
		if len(tr.Kids) > 0 {
			newTracks = append(newTracks, tr)
			continue
		}
		res = tr.Val
		resBool = true
	}
	t.tracks = newTracks
	return res, resBool
}
