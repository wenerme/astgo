package ami

// TODO
type AMI interface {
	Login(LoginAction) *GeneralResponse
	PJSIPShowEndpoint(PJSIPShowEndpointAction) (*EndpointDetailCompleteResponse, error)
	Ping() (string, error)
}

type _ami struct {
	c *_con
}

func ConnectManager(addr string) (*_ami, error) {
	c, err := connect(addr)
	if err != nil {
		return nil, err
	}
	a := &_ami{
		c: c,
	}
	return a, nil
}
