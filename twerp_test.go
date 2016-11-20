package twerp_test

import (
	"testing"

	"github.com/cheekybits/is"

	"github.com/JoshuaKolden/twerp"
)

func TestClamp(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		T float64
		X float64
		Y float64
		R float64
	}{
		{-1, 0, 1, 0},
		{2, 0, 1, 1},
		{0.5, 0, 1, 0.5},
		{-1, 1, 0, 0},
		{2, 1, 0, 1},
		{0.5, 1, 0, 0.5},
		{32.32, 32.2, 37.87, 32.32},
	}

	for _, test := range tests {
		r := twerp.Clamp(test.T, test.X, test.Y)
		is.Equal(test.R, r)
	}

}
