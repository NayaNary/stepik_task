package main

import (
	"fmt"
	"stepik_task/baseConstruction/interfaces"
	"stepik_task/baseConstruction/interfaces/objs"
)

func main() {
  usage := objs.NewUsage("Сервис 1")
  usage.Increment()
  fmt.Println(usage.Value())
  fmt.Println(usage.Service())
}

func measure(g interfaces.Geometry) {
	fmt.Printf("%T: %+v\n", g, g)
	fmt.Println("area:", g.Area())
	fmt.Println("perimiter:", g.Perim())
}
