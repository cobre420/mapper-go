package domain

type StateVector interface {
	Vector(name string) *Vector
	State() *State
}

type stateVector struct {
	vectors map[string]*Vector
	state   *State
}

func (s *stateVector) Vector(name string) *Vector {
	v, ok := s.vectors[name]
	if !ok {
		v = &Vector{}
		s.vectors[name] = v
	}
	return v
}

func (s *stateVector) State() *State {
	return s.state
}

func NewStateVector() StateVector {
	result := &stateVector{
		vectors: make(map[string]*Vector),
		state:   &State{},
	}
	return result
}

type Vector struct {
	Text string
}

type State struct {
	Text string
}
