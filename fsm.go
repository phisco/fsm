package fsm

type FSM struct {
	currentState State
}

func NewFSM(s State) *FSM {
	return &FSM{s}
}

type State interface {
	next() (nextState State, err error)
}

func (f *FSM) getCurrentState() State {
	return f.currentState
}

func (f *FSM) tic() error {
	nextState, err := f.currentState.next()
	if err != nil {
		return err
	}
	f.currentState = nextState
	return nil
}
