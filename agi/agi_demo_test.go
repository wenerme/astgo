package agi_test

import "github.com/wenerme/astgo/agi"

func ExampleRun() {
	agi.Run(func(session *agi.Session) {
		client := session.Client()
		client.Answer()
		client.StreamFile("activated", "#")
		client.SetVariable("AGISTATUS", "SUCCESS")
		client.Hangup()
	})
}
