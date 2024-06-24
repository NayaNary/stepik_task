package interfaces

type Geometry interface {
	Area() float64
	Perim() float64
}

type Counter interface {
	Increment()
	Value() uint
}

// element - интерфейс элемента последовательности
type Element interface{}

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type Iterator interface {
	// чтобы понять сигнатуры методов - посмотрите,
	// как они используются в функции max() ниже
	Next() bool
	Val() Element
}
