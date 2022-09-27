package interfaces

type Geometry interface{
    Area() float64
    Perim() float64
}

type Counter interface{
    Increment()
    Value() uint 
}