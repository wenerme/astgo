package xmlgen

import (
	"github.com/huandu/xstrings"
	"strings"
)

type AstDoc struct {
	Actions []*Type
	Events  []*Type
}
type Model struct {
	AGICommands []*AGICommand
	Actions     []*ManagerAction
	Events      []*ManagerEvent
}
type AGICommand struct {
	Name        string
	Module      string
	Syntax      *Syntax
	Synopsis    string
	Description string
	SeeAlso     []*SeeAlso
}
type ManagerAction struct {
	Name        string
	Synopsis    string
	Syntax      *Syntax
	Description string
	Response    *ManagerEvent
	SeeAlso     []*SeeAlso
	Responses   []*ManagerEvent
}

func (m ManagerAction) StructName() string {
	return validGoTypeName(m.Name) + "Action"
}

type ManagerEvent struct {
	Name     string
	Synopsis string
	Syntax   *Syntax
	SeeAlso  []*SeeAlso
}

func (m ManagerEvent) StructName() string {
	return validGoTypeName(m.Name) + "Event"
}

type SeeAlso struct {
	Type string
	Name string
}

var nameFixReplace = strings.NewReplacer("Callerid", "CallerID")

func (a AGICommand) StructName() string {
	switch a.Name {
	case "set autohangup":
		return "SetAutoHangup"
	case "asyncagi break":
		return "AsyncAGIBreak"
	case "database deltree":
		return "DatabaseDelTree"
	}
	name := xstrings.ToCamelCase(a.Name)
	name = nameFixReplace.Replace(name)
	return name
}

type Syntax struct {
	Params  []*Parameter
	Missing bool
}

func (d Syntax) HasRequiredParam() bool {
	for _, v := range d.Params {
		if v.Required {
			return true
		}
	}
	return false
}

type Parameter struct {
	RawName     string
	Name        string
	Type        string
	Default     string
	Required    bool
	Description string
	/*
				enumlist
		    <parameter required="true">
		      <enumlist>
		        <enum>
		          <parameter name="on" literal="true" required="true"/>
		        </enum>
		        <enum>
		          <parameter name="off" literal="true" required="true"/>
		        </enum>
		      </enumlist>
		    </parameter>
	*/
}
