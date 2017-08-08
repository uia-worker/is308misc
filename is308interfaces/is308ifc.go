package is308ifc

// I Implementation of interface I
type I interface {
	M() string
}

// T the type for the interface?
type T struct {
	Name string
}

// M string implementation of the interace for T
func (t T) M() string {
	return t.Name
}
