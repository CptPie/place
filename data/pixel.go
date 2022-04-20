package data

import "fmt"

type Pixel struct {
	Color Color
	X     int
	Y     int
}

func (p Pixel) String() string {
	return fmt.Sprintf("(%d, %d): %s", p.X, p.Y, p.Color.Name)
}
