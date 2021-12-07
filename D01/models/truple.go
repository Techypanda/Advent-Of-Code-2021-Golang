package models

// Pointers are used to make sure we have ATLEAST 3, could also use 0 val but not sure if they expect 0 as valid
type Truple struct {
	First           *float64
	Second          *float64
	Third           *float64
	currentRotation int
}

func (t *Truple) SetLatestVal(val float64) {
	if t.currentRotation > 3 || t.currentRotation < 1 {
		t.currentRotation = 1
	}
	switch t.currentRotation {
	case 1:
		t.First = &val
	case 2:
		t.Second = &val
	case 3:
		t.Third = &val
	default:
		panic("current rotation is not 1 2 or 3")
	}
	t.currentRotation += 1
}

func (t *Truple) ComputeSum() *float64 {
	if t.First == nil || t.Second == nil || t.Third == nil {
		return nil
	}
	res := (*t.First + *t.Second + *t.Third)
	return &res
}
