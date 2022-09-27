package baseconstruction

// Counter представляет целочисленный счетчик
type Counter struct {
	value uint
}

// increment увеличивает счетчик на единицу
func (c *Counter) increment() {
	c.value++
}

func (c *Counter) Value() uint {
	return c.value
}

func Inc(c *Counter) func() {
	return c.increment
}
