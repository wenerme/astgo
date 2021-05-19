package xmlgen

//go:generate gomodifytags -file=doc_model.go -w -all -add-tags json -transform camelcase --skip-unexported -add-options json=omitempty

import (
	"encoding/xml"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"regexp"
	"strings"
)

type AgiDocModel struct {
	//Text        string           `xml:",chardata"`
	Name        string            `xml:"name,attr" json:"name,omitempty"`
	Language    string            `xml:"language,attr" json:"language,omitempty"`
	Synopsis    string            `xml:"synopsis" json:"synopsis,omitempty"`
	Syntax      *SyntaxDocModel   `xml:"syntax" json:"syntax,omitempty"`
	SeeAlso     *SeeAlsoModel     `xml:"see-also" json:"seeAlso,omitempty"`
	Description *DescriptionModel `xml:"description" json:"description,omitempty"`
}

type ManagerResponseListElementModel struct {
	Text         string               `xml:",chardata" json:"text,omitempty"`
	ManagerEvent []*ManagerEventModel `xml:"managerEvent" json:"managerEvent,omitempty"`
}
type ManagerResponseModel struct {
	Text         string                           `xml:",chardata" json:"text,omitempty"`
	ListElements *ManagerResponseListElementModel `xml:"list-elements" json:"listElements,omitempty"`
	ManagerEvent *ManagerEventModel               `xml:"managerEvent" json:"managerEvent,omitempty"`
}

func (m ManagerResponseModel) Events() []*ManagerEventModel {
	var o []*ManagerEventModel
	if m.ManagerEvent != nil {
		o = append(o, m.ManagerEvent)
	}
	if m.ListElements != nil {
		o = append(o, m.ListElements.ManagerEvent...)
	}
	return o
}

type ManagerActionModel struct {
	Name        string               `xml:"name,attr" json:"name,omitempty"`
	Language    string               `xml:"language,attr" json:"language,omitempty"`
	Synopsis    string               `xml:"synopsis" json:"synopsis,omitempty"`
	Syntax      *SyntaxDocModel      `xml:"syntax" json:"syntax,omitempty"`
	SeeAlso     *SeeAlsoModel        `xml:"see-also" json:"seeAlso,omitempty"`
	Description *DescriptionModel    `xml:"description" json:"description,omitempty"`
	Responses   ManagerResponseModel `xml:"responses" json:"responses,omitempty"`
}
type ManagerEventInstanceModel struct {
	Text        string            `xml:",chardata"`
	Class       string            `xml:"class,attr"`
	Synopsis    string            `xml:"synopsis"`
	Language    string            `xml:"language,attr" json:"language,omitempty"`
	Syntax      *SyntaxDocModel   `xml:"syntax" json:"syntax,omitempty"`
	SeeAlso     *SeeAlsoModel     `xml:"see-also" json:"seeAlso,omitempty"`
	Description *DescriptionModel `xml:"description" json:"description,omitempty"`
}
type ManagerEventModel struct {
	Name     string                    `xml:"name,attr" json:"name,omitempty"`
	Instance ManagerEventInstanceModel `xml:"managerEventInstance"`
}

type SyntaxDocModel struct {
	Parameter       []SyntaxParameterDocModel `xml:"parameter" json:"parameter,omitempty"`
	ChannelSnapshot []struct {
		Text   string `xml:",chardata"`
		Prefix string `xml:"prefix,attr"`
	} `xml:"channel_snapshot"`
	BridgeSnapshot []struct {
		Text   string `xml:",chardata"`
		Prefix string `xml:"prefix,attr"`
	} `xml:"bridge_snapshot"`
}
type SyntaxParameterDocModel struct {
	//Text     string        `xml:",chardata"`
	Name     string         `xml:"name,attr" json:"name,omitempty"`
	Required bool           `xml:"required,attr" json:"required,omitempty"`
	Para     ParaRaw        `xml:"para" json:"para,omitempty"`
	EnumList *EnumListModel `xml:"enumlist" json:"enumList,omitempty"`

	Note *struct {
		Text string `xml:",chardata"`
		Para struct {
			Text        string   `xml:",chardata"`
			Literal     []string `xml:"literal"`
			Filename    string   `xml:"filename"`
			Replaceable string   `xml:"replaceable"`
		} `xml:"para"`
	} `xml:"note"`
}
type DocModel struct {
	XMLName      xml.Name              `xml:"docs" json:"xmlName,omitempty"`
	Agi          []*AgiDocModel        `xml:"agi" json:"agi,omitempty"`
	Manager      []*ManagerActionModel `xml:"manager" json:"manager,omitempty"`
	ManagerEvent []*ManagerEventModel  `xml:"managerEvent" json:"managerEvent,omitempty"`
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
	Note         *struct {
		Text string `xml:",chardata"`
		Para struct {
			Text     string `xml:",chardata"`
			Filename string `xml:"filename"`
		} `xml:"para"`
	} `xml:"note"`
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
	Text string `xml:",chardata" json:"text,omitempty"`
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

func (d *DocModel) Synopsis(in string) string {
	s := strings.TrimSpace(in)
	s = deindent(s)
	return s
}
func (d *DocModel) ManagerModel(in *ManagerActionModel) *ManagerAction {
	o := &ManagerAction{
		Name:        in.Name,
		Synopsis:    d.Synopsis(in.Synopsis),
		Syntax:      d.Syntax(in.Syntax),
		SeeAlso:     d.SeeAlso(in.SeeAlso),
		Description: d.Description(in.Description),
	}
	resp := in.Responses.Events()

	for _, v := range resp {
		o.Responses = append(o.Responses, d.ManagerEventModel(v))
	}

	if len(resp) > 1 {
		fmt.Println("Multi response", len(resp), o.Name)
		for _, v := range resp {
			fmt.Print(" ", v.Name)
		}
		fmt.Println()
	}
	return o
}

func (d *DocModel) ManagerEventModel(in *ManagerEventModel) *ManagerEvent {
	o := &ManagerEvent{
		Name:     in.Name,
		Synopsis: d.Synopsis(in.Instance.Synopsis),
		Syntax:   d.Syntax(in.Instance.Syntax),
		SeeAlso:  d.SeeAlso(in.Instance.SeeAlso),
	}
	return o
}
func (d *DocModel) Description(in *DescriptionModel) string {
	if in == nil {
		return ""
	}
	var p []string
	for _, v := range in.Para {
		p = append(p, v.Data)
	}
	// ignore enum,note
	return formatParas(p)
}
func (d *DocModel) AGIModel(in *AgiDocModel) *AGICommand {
	o := &AGICommand{
		Name:        in.Name,
		Synopsis:    d.Synopsis(in.Synopsis),
		Syntax:      d.Syntax(in.Syntax),
		SeeAlso:     d.SeeAlso(in.SeeAlso),
		Description: d.Description(in.Description),
	}
	return o
}
func (d *DocModel) SeeAlso(in *SeeAlsoModel) (o []*SeeAlso) {
	if in == nil {
		return
	}
	for _, v := range in.Ref {
		o = append(o, &SeeAlso{
			Name: v.Text,
			Type: v.Type,
		})
	}
	return
}
func (d *DocModel) Syntax(in *SyntaxDocModel) *Syntax {
	o := &Syntax{}

	if in == nil {
		return o
	}
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
	"<variable>", "`", "</variable>", "`",
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
func (*DocModel) Parameter(in *SyntaxParameterDocModel) *Parameter {
	o := &Parameter{
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

func processParameter(p *Parameter) {
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
	case "":
	default:
		if name[0] >= '0' && name[0] <= '9' {
			name = "Field" + name
		}
	}

	p.Name = name
	p.Type = typ
}
