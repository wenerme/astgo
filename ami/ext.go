package ami

type PongResponse struct {
	Pong string
}

type HasResponseType interface {
	ResponseType() interface{}
}

func (PingAction) ResponseType() interface{} {
	return &PongResponse{}
}

func init() {
}
