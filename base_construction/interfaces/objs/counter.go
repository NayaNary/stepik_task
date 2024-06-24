package objs

type counter struct {
    val uint
}
func (c *counter) Increment() {
    c.val++
}
func (c *counter) Value() uint {
    return c.val
}