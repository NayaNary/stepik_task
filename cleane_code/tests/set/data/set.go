package data

// IntSet implements a mathematical set
// of integers, elements are unique.
type IntSet struct {
    elems map[int]bool
}

// MakeIntSet creates an empty set.
func MakeIntSet() IntSet {
    elems := make(map[int]bool, 0)
    return IntSet{elems}
}

// Contains reports whether an element is within the set.
func (s IntSet) Contains(elem int) bool {
    if _, ok := s.elems[elem]; ok{
		return true
	}
    return false
}

// Add adds an element to the set.
func (s IntSet) Add(elem int) bool {
    if s.Contains(elem) {
        return false
    }
    s.elems[elem] = true
    return true
}