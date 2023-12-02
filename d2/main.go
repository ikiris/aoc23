package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func p1(r io.Reader) (int64, error) {
	gamemax := map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	s := bufio.NewScanner(r)
	re := regexp.MustCompile(`Game (\d+):((?:(?:((?:((?: \d+ \w+)+),?)+);?))+)`)
	re2 := regexp.MustCompile(`((((?: \d+ \w+)+),?)+);?`)
	re3 := regexp.MustCompile(`(\d+) (\w+)[,]?`)
	games := make(map[int64]bool)
	for s.Scan() {
		stuff := re.FindStringSubmatch(s.Text())
		if len(stuff) == 0 {
			return 0, fmt.Errorf("failed to parse %s", s.Text())
		}
		game, err := strconv.ParseInt(stuff[1], 10, 64)
		if err != nil {
			return 0, err
		}
		games[game] = false
		stuffb := re2.FindAllStringSubmatch(stuff[2], -1)
		for _, v := range stuffb {
			stuffc := re3.FindAllStringSubmatch(v[1], -1)
			for _, i := range stuffc {
				c, err := strconv.ParseInt(i[1], 10, 64)
				if err != nil {
					return 0, err
				}
				color := i[2]
				if gamemax[color] < c {
					games[game] = true
				}
			}
		}
	}
	var ttl int64
	for i, v := range games {
		if v {
			continue
		}
		ttl += int64(i)
	}
	return ttl, nil
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	re := regexp.MustCompile(`Game (\d+):((?:(?:((?:((?: \d+ \w+)+),?)+);?))+)`)
	re2 := regexp.MustCompile(`((((?: \d+ \w+)+),?)+);?`)
	re3 := regexp.MustCompile(`(\d+) (\w+)[,]?`)
	gamesmax := map[int64]map[string]int64{}
	for s.Scan() {
		stuff := re.FindStringSubmatch(s.Text())
		if len(stuff) == 0 {
			return 0, fmt.Errorf("failed to parse %s", s.Text())
		}
		game, err := strconv.ParseInt(stuff[1], 10, 64)
		if err != nil {
			return 0, err
		}
		gamesmax[game] = map[string]int64{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		stuffb := re2.FindAllStringSubmatch(stuff[2], -1)
		for _, v := range stuffb {
			stuffc := re3.FindAllStringSubmatch(v[1], -1)
			for _, i := range stuffc {
				c, err := strconv.ParseInt(i[1], 10, 64)
				if err != nil {
					return 0, err
				}
				color := i[2]
				if gamesmax[game][color] < c {
					gamesmax[game][color] = c
				}
			}
		}
	}
	var ttl int64
	for _, v := range gamesmax {
		ttl += v["red"] * v["green"] * v["blue"]
	}
	return ttl, nil
}

func main() {}
