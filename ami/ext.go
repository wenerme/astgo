package ami

import "errors"

type GeneralResponse struct {
	// Error or Success
	Response string
	Message  string
}

func (self GeneralResponse) IsSuccess() bool {
	return self.Response == "Success"
}
func (self GeneralResponse) ToError() error {
	if self.IsSuccess() {
		return nil
	}
	return errors.New(self.Message)
}

type PingResponse struct {
	Ping string
}

type HasResponseType interface {
	ResponseType() interface{}
}

func (PingAction) ResponseType() interface{} {
	return &PingResponse{}
}

func (LoginAction) ResponseType() interface{} {
	return &GeneralResponse{}
}

func init() {
}
