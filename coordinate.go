package ggworkshop

import "fmt"

// Coordinate is a struct that represents an X,Y coordinate.
type Coordinate struct {
	// X is the X-vector of the coordinate.
	X int
	// Y is the Y-vector of the coordinate.
	Y int
}

// String formats the Coordinate to a string. It formats to (X, Y)
func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}
