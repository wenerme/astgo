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
}
type AGICommand struct {
	Name        string
	Module      string
	Syntax      *SyntaxDoc
	Synopsis    string
	Description string
	// synopsis?,syntax?,description?,see-also?
	/*
	   <see-also>
	     <ref type="application">AGI</ref>
	   </see-also>
	*/
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

type SyntaxDoc struct {
	Params  []*SyntaxParamDoc
	Missing bool
}

func (d SyntaxDoc) HasRequiredParam() bool {
	for _, v := range d.Params {
		if v.Required {
			return true
		}
	}
	return false
}

type SyntaxParamDoc struct {
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
