package amimodels

import "github.com/pkg/errors"

func (evt OriginateResponseEvent) Err() error {
	if evt.Response == "Failure" {
		return errors.Errorf("Originate failed: exten %v reason %v", evt.Exten, evt.Reason)
	}
	return nil
}
