package xmlgen

import (
	"encoding/xml"
	"github.com/iancoleman/strcase"
	"log"
	"regexp"
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
	var p []string
	for _, v := range in.Description.Para {
		p = append(p, v.Data)
	}
	o.Description = formatParas(p)
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

var regDefaultTo = regexp.MustCompile(`Defaults to\s*<literal>(.*?)</literal>`)

func parseParaDefaultTo(s string) string {
	sub := regDefaultTo.FindStringSubmatch(s)
	if len(sub) > 1 {
		return sub[1]
	}
	return ""
}

var regIndent = regexp.MustCompile(`(?m)^\s+`)

func deindent(s string) string {
	return regIndent.ReplaceAllString(s, "")
}

var paraTagReplace = strings.NewReplacer(
	"<literal>", "`", "</literal>", "`",
	"<directory>", "`", "</directory>", "`",
	"<filename>", " ", "</filename>", " ",
	"<replaceable>", " ", "</replaceable>", " ",
	"\n", " ",
)

func formatParas(s []string) string {
	sb := strings.Builder{}
	for _, v := range s {
		sb.WriteString(formatPara(v))
		sb.WriteString("\n\n")
	}
	return strings.TrimSpace(sb.String())
}
func formatPara(s string) string {
	s = deindent(s)
	return paraTagReplace.Replace(s)
}
func (*DocModel) Parameter(in *SyntaxParameterDocModel) *SyntaxParamDoc {
	o := &SyntaxParamDoc{
		RawName:  in.Name,
		Required: in.Required,
	}
	para := strings.TrimSpace(in.Para.Data)
	if para != "" {
		o.Default = parseParaDefaultTo(para)
		if o.Default == "" {
			// description
			o.Description = formatPara(para)
		}
	}
	processParameter(o)
	return o
}

func processParameter(p *SyntaxParamDoc) {
	rawName := p.RawName
	var name string
	typ := "string"
	switch rawName {
	case "skipms":
		name = "SkipMS"
	case "offsetms":
		name = "OffsetMS"
	case "keytree":
		name = "KeyTree"
	}
	if name == "" {
		name = rawName
		// s=slicense
		name = strings.ReplaceAll(name, "=", "-")
		if len(name) > 4 {
			name = strings.ReplaceAll(name, "name", "Name")
		}
		// xstrings VariableName -> Variablename
		name = strcase.ToCamel(name)
	}
	switch name {
	case "SkipMS", "OffsetMS", "SampleOffset", "Priority", "Timeout":
		typ = "int"
	case "Time": // AutoHangup
		typ = "float64"
	}
	p.Name = name
	p.Type = typ
}
