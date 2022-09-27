package objs

import (
	"stepik_task/baseConstruction/interfaces"
)

type rect struct {
	width, height float64
}

func NewRect(w, h float64) interfaces.Geometry {
	return &rect{width: w, height: h}
}

func (r rect) Area() float64 {
	return r.width * r.height
}

func (r rect) Perim() float64 {
	return 2*r.width + 2*r.height
}
