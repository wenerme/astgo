package amimodels

type Event interface {
	EventTypeName() string
}
type Action interface {
	ActionTypeName() string
	//GetActionID() string
	//SetActionID(actionID string) // need ptr
}
