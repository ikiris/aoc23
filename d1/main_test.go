package main

import (
	"io"
	"testing"

	"github.com/ikiris/aoc21/generic/testgeneric"
)

func TestP1(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/input1.txt"),
			142,
			false,
		},
		{
			"d1",
			testgeneric.GetHandle(t, "testdata/input2.txt"),
			55108,
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p1(tc.data)
			if (err != nil) != tc.wantErr {
				t.Fatalf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("func got: %d, want: %d", got, tc.want)
			}
		})
	}
}

func TestP2(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"d1.51",
			testgeneric.GetHandle(t, "testdata/input3.txt"),
			281,
			false,
		},
		{
			"d1.52",
			testgeneric.GetHandle(t, "testdata/input2.txt"),
			56324,
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p2(tc.data)
			if (err != nil) != tc.wantErr {
				t.Fatalf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("func got: %d, want: %d", got, tc.want)
			}
		})
	}
}
