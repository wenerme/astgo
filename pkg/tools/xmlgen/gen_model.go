package xmlgen

import "github.com/huandu/xstrings"

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

func (a AGICommand) StructName() string {
	return xstrings.ToCamelCase(a.Name)
}

type SyntaxDoc struct {
	Params  []*SyntaxParamDoc
	Missing bool
}

type SyntaxParamDoc struct {
	RawName     string
	Name        string
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
