package agimodels

import "context"

type AnswerCommand struct {
}

func (AnswerCommand) Command(c context.Context) (string, error) {
	return "ANSWER", nil
}

type HangupCommand struct {
}

func (HangupCommand) Command(c context.Context) (string, error) {
	return "HANGUP", nil
}

type GetVariableCommand struct {
	Name string
}

func (cmd GetVariableCommand) Command(c context.Context) (string, error) {
	return "GET VARIABLE " + cmd.Name, nil
}
