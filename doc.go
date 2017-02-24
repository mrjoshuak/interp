// The interp package is a collection of interpolation functions.
//
// Most functions are controlled by the first argument t as it ranges between 0.0 and 1.0.
//
// With the exception of 'Clamp', functions in this package take one of three forms.
//
// _: Functions that do not end with 'step' or 'mix' return normalized values that are not clamped.
//
// _Step: Function that end with 'step', return normalized values that clamp between 0.0 and 1.0
//
// _Mix: Functions that end with 'mix', return results that are not normalized, and rage across the provided arguments
//
package interp
