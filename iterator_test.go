package iterator

import "testing"

// IPair is a coordinates interface
type IPair interface {
	Equals(a IPair) bool
}

// BasePairs is an iterator underlying type
type BasePairs struct {
	slice []IPair
	i     int
}

// NewBasePairs returns a pointer to a new base type
func NewBasePairs(from []IPair) *BasePairs {
	return &BasePairs{slice: from}
}

// Next returns next coordinates element
func (d *BasePairs) Next() interface{} {
	d.i++
	return d.slice[d.i-1]
}

// HasNext returns true if an underlying slice has next element
func (d *BasePairs) HasNext() bool { return d.i < len(d.slice) }

// I returns a current iteration index
func (d *BasePairs) I() int { return d.i - 1 }

// SetI sets a current iteration index
func (d *BasePairs) SetI(i int) { d.i = i }

// Add adds an element to an underlying slice
func (d *BasePairs) Add(c interface{}) { d.slice = append(d.slice, c.(IPair)) }

// Pair is a coord pair
type Pair struct {
	x, y int
}

// Equals returns true if d equals to a
func (d Pair) Equals(a IPair) bool {
	return d.x == a.(Pair).x && d.y == a.(Pair).y
}

// SpecificPairs is a type which embedded base type
type SpecificPairs struct {
	*BasePairs
}

// NewSpecificPairs returns a new iterable descendant
func NewSpecificPairs(c []Pair) SpecificPairs {
	s := make([]IPair, len(c))
	for i := range c {
		s[i] = c[i]
	}
	return SpecificPairs{BasePairs: NewBasePairs(s)}
}

// TestIterator tests the iterator
func TestIterator(t *testing.T) {
	pairs := []Pair{{5, 5}, {4, 5}, {7, 8}}

	func(over Interface) {
		i := 0
		over.SetI(0)
		for over.HasNext() {
			nextElement := over.Next().(Pair)
			if over.I() != i {
				t.Fatalf("Expected i = %d, received i = %d", i, over.I())
			}
			if !nextElement.Equals(pairs[over.I()]) {
				t.Fatalf("Element not equals on iteration %d", i)
			}
			i++
		}
	}(NewSpecificPairs(pairs))
}
