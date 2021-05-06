package amimodels

type Event interface {
	EventTypeName() string
}
type Action interface {
	ActionTypeName() string
}
type HasActionID interface {
	GetActionID() string
	SetActionID(actionID string)
}
