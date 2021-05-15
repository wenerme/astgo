package xmlgen

import (
	"encoding/xml"
	"github.com/iancoleman/strcase"
	"log"
	"strings"
)

type AgiDocModel struct {
	//Text        string           `xml:",chardata"`
	Name        string           `xml:"name,attr" json:"name,omitempty"`
	Language    string           `xml:"language,attr" json:"language,omitempty"`
	Synopsis    string           `xml:"synopsis" json:"synopsis,omitempty"`
	Syntax      SyntaxDocModel   `xml:"syntax" json:"syntax,omitempty"`
	SeeAlso     SeeAlsoModel     `xml:"see-also" json:"seeAlso,omitempty"`
	Description DescriptionModel `xml:"description" json:"description,omitempty"`
}
type SyntaxDocModel struct {
	//Text      string                    `xml:",chardata"`
	Parameter []SyntaxParameterDocModel `xml:"parameter" json:"parameter,omitempty"`
}
type SyntaxParameterDocModel struct {
	//Text     string        `xml:",chardata"`
	Name     string         `xml:"name,attr" json:"name,omitempty"`
	Required bool           `xml:"required,attr" json:"required,omitempty"`
	Para     ParaRaw        `xml:"para" json:"para,omitempty"`
	EnumList *EnumListModel `xml:"enumlist" `
}
type DocModel struct {
	XMLName xml.Name      `xml:"docs" json:"xmlName,omitempty"`
	Agi     []AgiDocModel `xml:"agi" json:"agi,omitempty"`
}

// ParaRaw contain ELEMENT astcli|literal|emphasis|filename|directory|replaceable|variable
type ParaRaw struct {
	Data string `xml:",innerxml" json:"data,omitempty"`
}
type DescriptionModel struct {
	//Text         string            `xml:",chardata"`
	Para         []ParaRaw          `xml:"para" json:"para,omitempty"`
	EnumList     *EnumListModel     `xml:"enumlist" json:"enumList,omitempty"`
	VariableList *VariableListModel `xml:"variablelist" json:"variableList,omitempty"`
}

type VariableListVariableValueModel struct {
	Text string `xml:",chardata" json:"text,omitempty"`
	Name string `xml:"name,attr" json:"name,omitempty"`
}
type VariableListVariableModel struct {
	Name  string                           `xml:"name,attr" json:"name,omitempty"`
	Para  ParaRaw                          `xml:"para" json:"para,omitempty"`
	Value []VariableListVariableValueModel `xml:"value" json:"value,omitempty"`
}
type VariableListModel struct {
	Variable []VariableListVariableModel `xml:"variable" json:"variable,omitempty"`
}
type SeeAlsoModel struct {
	Ref []SeeAlsoRefModel `json:"ref,omitempty"`
}
type SeeAlsoRefModel struct {
	//Text string `xml:",chardata"`
	Type string `xml:"type,attr" json:"type,omitempty"`
}

type EnumModel struct {
	//Text string  `xml:",chardata"`
	Name string  `xml:"name,attr" json:"name,omitempty"`
	Para ParaRaw `xml:"para" json:"para,omitempty"`
}
type EnumListModel struct {
	//Text string      `xml:",chardata"`
	Enum []EnumModel `xml:"enum" json:"enum,omitempty"`
}

func (d *DocModel) AgiModel(in AgiDocModel) *AGICommand {
	o := &AGICommand{
		Name:     in.Name,
		Synopsis: strings.TrimSpace(in.Synopsis),
		Syntax:   d.Syntax(&in.Syntax),
	}

	return o
}
func (d *DocModel) Syntax(in *SyntaxDocModel) *SyntaxDoc {
	o := &SyntaxDoc{}
	for _, v := range in.Parameter {
		parameter := d.Parameter(&v)
		if parameter.Name == "" {
			log.Println("drop param", in)
			o.Missing = true
			continue
		}
		o.Params = append(o.Params, parameter)
	}
	return o
}
func (*DocModel) Parameter(in *SyntaxParameterDocModel) *SyntaxParamDoc {
	o := &SyntaxParamDoc{
		RawName:  in.Name,
		Name:     fieldName(in.Name), // name not matter
		Required: in.Required,
	}
	log.Println("Param", in.Name, o.Name)
	return o
}

func fieldName(name string) string {
	switch name {
	case "offsetms":
		return "OffsetMS"
	case "keytree":
		return "KeyTree"
	}
	// s=slicense
	name = strings.ReplaceAll(name, "=", "-")
	if len(name) > 4 {
		name = strings.ReplaceAll(name, "name", "Name")
	}
	// xstrings VariableName -> Variablename
	return strcase.ToCamel(name)
}
