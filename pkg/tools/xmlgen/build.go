package xmlgen

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func BuildAstDoc(doc *Docs, d *AstDoc) {
	byName := map[string]*Type{}
	byFieldName := map[string]map[string]*Field{}
	for _, e := range doc.Manager {
		typ, dup := byName[e.Name]
		if !dup {
			log.Println("duplicate action", e.Name)
			typ = &Type{
				Name: e.Name,
				Doc:  strings.TrimSpace(e.Synopsis),
			}
		}

		if byFieldName[typ.Name] == nil {
			byFieldName[typ.Name] = map[string]*Field{}
		}

		for _, p := range e.Syntax.Parameter {
			name := p.Name
			fieldName := goName(name)

			field, dup := byFieldName[typ.Name][fieldName]
			if dup {
				log.Println("duplicate field", typ.Name, fieldName)
				continue
			}

			field = &Field{
				Name:     fieldName,
				Required: p.Required != "",
				Default:  p.Default,
				Type:     &TypeInfo{Type: "string"},
			}

			var s []string
			for _, v := range p.Para {
				text := strings.TrimSpace(v.Text)
				if text != "" {
					s = append(s, text)
				} else {
					log.Println(v)
				}
			}
			field.Doc = strings.Join(s, "\n")

			if field.Name != name {
				field.StructTag = fmt.Sprintf(`json:"%s,omitempty"`, name)
			}

			byFieldName[typ.Name][fieldName] = field
			typ.Fields = append(typ.Fields, field)
		}

		if !dup {
			byName[typ.Name] = typ
			d.Actions = append(d.Actions, typ)
		}
	}
	{
		byName := map[string]*Type{}
		byFieldName := map[string]map[string]*Field{}
		for _, e := range doc.ManagerEvent {
			typeName := e.Name
			typ, dup := byName[typeName]
			e := e.ManagerEventInstance
			if !dup {
				log.Println("duplicate action", typeName)
				typ = &Type{
					Name: goName(typeName),
					Doc:  strings.TrimSpace(e.Synopsis),
				}
			}

			if byFieldName[typ.Name] == nil {
				byFieldName[typ.Name] = map[string]*Field{}
			}

			for _, p := range e.Syntax.Parameter {
				name := p.Name
				fieldName := goName(name)

				field, dup := byFieldName[typ.Name][fieldName]
				if dup {
					log.Println("duplicate field", typ.Name, fieldName)
					continue
				}

				field = &Field{
					Name:     fieldName,
					Required: p.Required != "",
					// Default:  p.Default,
					Type: &TypeInfo{Type: "string"},
				}

				var s []string
				for _, v := range p.Para {
					text := strings.TrimSpace(v.Text)
					if text != "" {
						s = append(s, text)
					} else {
						log.Println(v)
					}
				}
				field.Doc = strings.Join(s, "\n")

				if field.Name != name {
					field.StructTag = fmt.Sprintf(`json:"%s,omitempty"`, name)
				}

				byFieldName[typ.Name][fieldName] = field
				typ.Fields = append(typ.Fields, field)
			}

			if !dup {
				byName[typ.Name] = typ
				d.Events = append(d.Events, typ)
			}
		}
	}
	for _, v := range d.Actions {
		_, ok := byFieldName[v.Name]["ActionID"]
		if !ok {
			v.Fields = append(v.Fields, &Field{
				Name: "ActionID",
				Doc:  "ActionID for tx - compensate",
				Type: &TypeInfo{
					Type: "string",
				},
			})
		}
	}

}

var invalid = regexp.MustCompile(`[- ]|\(.*?\)`)
var valid = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*$")

func goName(s string) string {
	s = invalid.ReplaceAllLiteralString(s, "")
	if s[0] >= '0' && s[0] <= '9' {
		s = "Field" + s
	}
	// UnitAmount(0)
	if !valid.MatchString(s) {
		panic("invalid id: " + s)
	}
	return s
}
