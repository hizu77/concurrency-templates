package model

type State int

func (s State) String() string {
	switch s {
	case StateCreated:
		return "created"
	case StateStarted:
		return "started"
	case StateProcessed:
		return "processed"
	case StateCompleted:
		return "completed"
	default:
		return "undefined"
	}
}

const (
	StateCreated State = iota
	StateStarted
	StateProcessed
	StateCompleted
)

type Order struct {
	ID    int
	State State
}
