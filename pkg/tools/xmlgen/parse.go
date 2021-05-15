package xmlgen

type Type struct {
	Name   string
	Doc    string
	Fields []*Field
}
type Field struct {
	Name        string
	Doc         string
	Default     string
	Required    bool
	Type        *TypeInfo
	StructTag   string
	Annotations interface{}
}
type Annotation interface {
	Name() string
}
type TypeInfo struct {
	Type string
}

func (ti *TypeInfo) String() string {
	return ti.Type
}

type AstDoc struct {
	Actions []*Type
	Events  []*Type
}

type AgiCommandDoc struct {
	Name        string
	Module      string
	Syntax      SyntaxDoc
	Synopsis    string
	Description string
	// synopsis?,syntax?,description?,see-also?
	/*
	   <see-also>
	     <ref type="application">AGI</ref>
	   </see-also>
	*/
}

type SyntaxDoc struct {
	Params []SyntaxParamDoc
}

type SyntaxParamDoc struct {
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
