package xmlgen

import (
	"fmt"
	"reflect"
	"regexp"
)

func BuildModel(in *DocModel, out *Model) {
	for _, v := range in.Agi {
		out.AGICommands = append(out.AGICommands, in.AGIModel(v))
	}
	for _, v := range in.Manager {
		out.Actions = append(out.Actions, in.ManagerModel(v))
	}
	for _, v := range in.ManagerEvent {
		out.Events = append(out.Events, in.ManagerEventModel(v))
	}
	//
	Dedup(&out.Actions, func(i int) string {
		return out.Actions[i].Name
	})
	Dedup(&out.Events, func(i int) string {
		return out.Events[i].Name
	})
	for _, v := range out.Events {
		Dedup(&v.Syntax.Params, func(i int) string {
			return v.Syntax.Params[i].Name
		})
	}

	// add missing event, unify event
	events := map[string]*ManagerEvent{}
	for _, v := range out.Events {
		events[v.Name] = v
	}
	for _, action := range out.Actions {
		for k, v := range action.Responses {
			if e, ok := events[v.Name]; ok {
				action.Responses[k] = e
			} else {
				out.Events = append(out.Events, v)
				fmt.Println("Append Response Event", v.Name)
			}
		}
	}

	// add missing ActionID
	for _, v := range out.Actions {
		if len(v.Syntax.Params) == 0 {

		}
		found := false
		for _, v := range v.Syntax.Params {
			if found = v.Name == "ActionID"; found {
				break
			}
		}
		if !found {
			o := []*Parameter{{
				Name:        "ActionID",
				Type:        "string",
				Description: "ActionID for this transaction. Will be returned.",
			}}
			v.Syntax.Params = append(o, v.Syntax.Params...)
		}
	}
}

var regInvalid = regexp.MustCompile(`[- ]|\(.*?\)`)

func validGoTypeName(s string) string {
	s = invalid.ReplaceAllLiteralString(s, "")
	return s
}
func Dedup(arr interface{}, cond func(int) string) {
	m := map[string]bool{}
	Filter(arr, func(i int) bool {
		k := cond(i)
		if _, ok := m[k]; ok {
			return false
		}
		m[k] = true
		return true
	})
}

func Filter(arr interface{}, cond func(int) bool) {
	rv := reflect.ValueOf(arr).Elem()
	rt := rv.Type()

	neo := reflect.MakeSlice(rt, 0, 0)
	for i := 0; i < rv.Len(); i++ {
		if v := rv.Index(i); cond(i) {
			neo = reflect.Append(neo, v)
		}
	}
	rv.Set(neo)
}
