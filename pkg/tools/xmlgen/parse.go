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
