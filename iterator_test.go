package iterator

import "testing"

// IPair is a coordinates interface
type IPair interface {
	Equals(a IPair) bool
}

// BasePair is an iterator underlying type
type BasePair struct {
	slice []IPair
	i     int
}

// NewBasePair returns a pointer to a new base type
func NewBasePair(from []IPair) *BasePair {
	return &BasePair{slice: from}
}

// Pair is a coord pair
type Pair struct {
	x, y int
}

// Equals returns true if d equals to a
func (d Pair) Equals(a IPair) bool {
	return d.x == a.(Pair).x && d.y == a.(Pair).y
}

// SpecificPair is a type which embedded base type
type SpecificPair struct {
	*BasePair
}

// NewSpecificPair returns a new descendant
func NewSpecificPair(c []Pair) SpecificPair {
	s := make([]IPair, len(c))
	for i := range c {
		s[i] = c[i]
	}
	return SpecificPair{BasePair: NewBasePair(s)}
}

// Next returns next coordinates element
func (d *SpecificPair) Next() interface{} {
	d.i++
	return d.slice[d.i-1]
}

// HasNext returns true if an underlying slice has next element
func (d *SpecificPair) HasNext() bool { return d.i < len(d.slice) }

// I returns a current iteration index
func (d *SpecificPair) I() int { return d.i - 1 }

// Add adds an element to an underlying slice
func (d *SpecificPair) Add(c interface{}) { d.slice = append(d.slice, c.(IPair)) }

// TestIterator tests the iterator
func TestIterator(t *testing.T) {
	pairs := []Pair{{5, 5}, {4, 5}, {7, 8}}
	specificPairs := NewSpecificPair(pairs)

	i := 0
	for specificPairs.HasNext() {
		nextElement := specificPairs.Next().(Pair)
		if specificPairs.I() != i {
			t.Fatalf("Expected i = %d, received i = %d", i, specificPairs.I())
		}
		if !nextElement.Equals(pairs[specificPairs.I()]) {
			t.Fatalf("Element not equals on iteration %d", i)
		}
		i++
	}
}
