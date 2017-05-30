package interp_test

import (
	"testing"

	"github.com/cheekybits/is"

	"github.com/JoshuaKolden/interp"
)

func TestInterp(t *testing.T) {

	t.Run("Clamp", func(t *testing.T) {
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
			r := interp.Clamp(test.T, test.X, test.Y)
			is.Equal(test.R, r)
		}
	})

	t.Run("Step", func(t *testing.T) {
		is := is.New(t)

		tests := []struct {
			T float64
			X float64
			R float64
		}{
			{-1, 0, 0},
			{1, 0, 1},
			{2, 0, 1},
			{2, 1, 1},
			{1, 2, 0},
			{0.5, 1, 0},
			{1.5, 1, 1},
			{-0.5, -1, 1},
			{-1.5, -1, 0},
		}

		for _, test := range tests {
			r := interp.Step(test.T, test.X)
			is.Equal(test.R, r)
		}
	})

	t.Run("Mix", func(t *testing.T) {
		is := is.New(t)

		tests := []struct {
			T float64
			X float64
			Y float64
			R float64
		}{
			{0, 0, 1, 0},
			{1, 0, 1, 1},
			{0, 27.5, 33.7, 27.5},
			{-0.5, 27.5, 33.7, 27.5},
			{1, 27.5, 33.7, 33.7},
			{1.5, 27.5, 33.7, 33.7},
			{0.5, 27.5, 33.7, 30.6},
			{0.5, 33.7, 27.5, 30.6},
			{0, 33.7, 27.5, 33.7},
			{1, 33.7, 27.5, 27.5},
			{0, -12.8362, -5, -12.8362},
			{1, -12.8362, -5, -5},
			{0, -5, -12.8362, -5},
			{1, -5, -12.8362, -12.8362},
		}

		for _, test := range tests {
			r := interp.Mix(test.T, test.X, test.Y)
			is.Equal(test.R, r)
		}

	})

	t.Run("Map", func(t *testing.T) {

		is := is.New(t)
		tests := []struct {
			R float64
			X float64
			Y float64
			T float64
		}{
			{0, 0, 1, 0},
			{1, 0, 1, 1},
			{0, 27.5, 33.7, 27.5},
			{0, 27.5, 33.7, 5.5},
			{1, 27.5, 33.7, 33.7},
			{1, 27.5, 33.7, 42.7},
			{0.5, 27.5, 33.7, 30.6},
			{0.5, 33.7, 27.5, 30.6},
			{0, 33.7, 27.5, 33.7},
			{1, 33.7, 27.5, 27.5},
			{0, -12.8362, -5, -12.8362},
			{1, -12.8362, -5, -5},
			{0, -5, -12.8362, -5},
			{1, -5, -12.8362, -12.8362},
		}

		for _, test := range tests {
			r := interp.Map(test.T, test.X, test.Y)
			is.Equal(test.R, r)
		}
		// Map(t float64, x float64, y float64)
	})

	t.Run("Linearstep", func(t *testing.T) {
		is := is.New(t)
		tests := []struct {
			T float64
			R float64
		}{
			{-1, 0},
			{2, 1},
			{0.5, 0.5},
			{2, 1},
		}

		for _, test := range tests {
			r := interp.Linearstep(test.T)
			is.Equal(test.R, r)
		}

	})

	t.Run("Smoothstep", func(t *testing.T) {
		is := is.New(t)
		tests := []struct {
			T float64
			R float64
		}{
			{-1, 0},
			{2, 1},
			{0.5, 0.5},
			{2, 1},
			{0.75, 0.84375},
			{0.25, 0.15625},
		}

		for _, test := range tests {
			r := interp.Smoothstep(test.T)
			is.Equal(test.R, r)
		}
	})

	t.Run("Smoothmix", func(t *testing.T) {
		is := is.New(t)
		tests := []struct {
			R float64
			X float64
			Y float64
			T float64
		}{
			{0, 0, 1, 0},
			{1, 0, 1, 1},
			{0, 27.5, 33.7, 27.5},
			{0, 27.5, 33.7, 5.5},
			{1, 27.5, 33.7, 33.7},
			{1, 27.5, 33.7, 42.7},
			{0.5, 27.5, 33.7, 30.6},
			{0.5, 33.7, 27.5, 30.6},
			{0, 33.7, 27.5, 33.7},
			{1, 33.7, 27.5, 27.5},
			{0, -12.8362, -5, -12.8362},
			{1, -12.8362, -5, -5},
			{0, -5, -12.8362, -5},
			{1, -5, -12.8362, -12.8362},
			{0.84375, 0, 1, 0.75},
			{0.15625, 0, 1, 0.25},
		}

		for _, test := range tests {
			r := interp.Smoothmix(test.T, test.X, test.Y)
			is.Equal(test.R, r)
		}
	})

	t.Run("Easein", func(t *testing.T) {
		is := is.New(t)
		is.Fail("test not implemented")
		// Easein(t float64, exp float64)
	})

	t.Run("Easeout", func(t *testing.T) {
		is := is.New(t)
		is.Fail("test not implemented")
		// Easeout(t float64, exp float64)
	})

	t.Run("Easeinstep", func(t *testing.T) {
		is := is.New(t)
		is.Fail("test not implemented")
		// Easeinstep(t float64, exp float64)
	})

	t.Run("Easeoutstep", func(t *testing.T) {
		is := is.New(t)
		is.Fail("test not implemented")
		// Easeoutstep(t float64, exp float64)
	})

}
