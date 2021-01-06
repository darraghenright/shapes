package shapes

import (
	"fmt"
)

// Circle struct
type Circle struct {
	Radius float64
}

// String method
func (c Circle) String() string {
	return fmt.Sprintf("Circle(%0.2f)", c.Radius)
}

// Area method
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
