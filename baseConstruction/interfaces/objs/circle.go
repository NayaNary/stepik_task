package objs

import (
	"math"
	"stepik_task/baseConstruction/interfaces"
)

type circle struct {
	radius float64
}

func NewCircle(r float64) interfaces.Geometry {
	return &circle{radius:r}
}

func (c *circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}
