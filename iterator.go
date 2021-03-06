package iterator

// Interface is an interface for iterator
type Interface interface {
	// Next should return the next element
	Next() interface{}
	// HasNext should return true if we have next element
	HasNext() bool
	// I should return an iteration index
	I() int
	// Add should add c to an underlying storage
	Add(c interface{})
	// SetI should set an iteration index to i
	SetI(i int)
}
