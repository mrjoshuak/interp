// The interp package is a collection of interpolation functions.
// Most functions are controlled by the first argument t as it
// ranges between 0.0 and 1.0.

// With the exception of 'Clamp', functions in this package take one of three forms.
//
// _: Functions that do not end with 'step' or 'mix' return normalized values that are not clamped.
// _Step: Function that end with 'step', return normalized values that clamp between 0.0 and 1.0
// _Mix: Functions that end with 'mix', return results that are not normalized, and rage across the provided arguments
//
package interp

import "math"

// Clamp returns t, constrained within the range of x and y.
func Clamp(t float64, x float64, y float64) (r float64) {
	if x > y {
		x, y = y, x
	}
	r = x
	if t > y {
		r = y
	} else if t > x {
		r = t
	}
	return r
}

// Step return 0 if t is less than x, otherwise it returns 1.
func Step(t float64, x float64) float64 {
	if t < x {
		return 0
	}
	return 1
}

// Mix returns x when t is 0 and y when t is 1.0, and blends between x and y between 0.0 and 1.0.
// Mix does not clamp results if t is outside the range 0.0 to 1.0, the results is a value along
// along the slope of x and y.
// In this context 't' is sometimes referred to as 'alpha'
func Mix(t float64, x float64, y float64) float64 {
	return x + t*(y-x) //t*y + (1-t)*x
}

// Map is the inverse of Mix, output is 0.0 when t == x and 1.0 when t == y
func Map(t float64, x float64, y float64) float64 {
	return (t - x) / (y - x)
}

// Linear mapping of t, clamped within 0.0 to 1.0
func Linearstep(t float64) (r float64) {
	t = Clamp(t, 0, 1)
	return t
}

// Easeinout mapping of t, clamped to the range 0 to 1
func Smoothstep(t float64) float64 {
	r := Linearstep(t)
	r = r * r * (3 - 2*r)
	return r
}

// Easeinout mapping of t, clamped to the range 0 to 1
func Smoothmix(t float64, x float64, y float64) (r float64) {
	r = Linearstep(t)
	r = r * r * (3 - 2*r)
	return
}

func Easein(t float64, exp float64) float64 {
	return 1 - math.Pow(1-t, exp)
}

func Easeout(t float64, exp float64) float64 {
	return math.Pow(t, exp)
}

func Easeinstep(t float64, exp float64) float64 {
	t = Clamp(t, 0, 1)
	return 1 - math.Pow(1-t, exp)
}

func Easeoutstep(t float64, exp float64) float64 {
	t = Clamp(t, 0, 1)
	return math.Pow(t, exp)
}
