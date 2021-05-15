package xmlgen

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/pkg/errors"
	"go/format"
	"golang.org/x/tools/imports"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

import (
	"github.com/Masterminds/sprig"
)

//go:embed template/*.tmpl template/**/*.tmpl
var templateFS embed.FS

func MustParseTemplates() *template.Template {
	//ext := pick(sprig.TxtFuncMap(), "ternary", "title")
	tpl := template.Must(
		template.New("gqlgen").
			//Funcs(gen.Funcs).
			//Funcs(ext).
			Funcs(sprig.TxtFuncMap()).
			//Funcs(Funcs).
			ParseFS(templateFS, "template/*.tmpl", "template/**/*.tmpl"),
	)
	//return &gen.Template{
	//	Template: tpl,
	//	FuncMap:  gen.Funcs,
	//}
	return tpl
}

func pick(funcs template.FuncMap, names ...string) template.FuncMap {
	o := make(template.FuncMap)
	for _, v := range names {
		o[v] = funcs[v]
	}
	return o
}

func NewAMIGenerator() *Generator {
	return &Generator{
		Template:  MustParseTemplates(),
		Formatter: GoFormatter,
		Templates: []TemplateGenerate{
			&GenerateTemplate{
				Name:       "ami/actions",
				FormatName: "actions.go",
			},
			&GenerateTemplate{
				Name:       "ami/events",
				FormatName: "events.go",
			},
		},
	}
}
func NewAGIGenerator() *Generator {
	return &Generator{
		Template:  MustParseTemplates(),
		Formatter: GoFormatter,
		Templates: []TemplateGenerate{
			&GenerateTemplate{
				Name:       "agi/commands",
				FormatName: "commands.go",
			},
		},
	}
}

type Generator struct {
	Files     []File
	Template  *template.Template
	Templates []TemplateGenerate
	Debug     bool
	Formatter func(file *File) error
	Loader    func() error
}
type File struct {
	Name    string
	Content []byte
}
type WriteConfig struct {
	Target string
	DryRun bool
}

func (ge *Generator) Write(s WriteConfig) error {
	for _, v := range ge.Files {
		o := filepath.Join(s.Target, v.Name)
		// log.Println("write ", v.Name, " to ", o)
		if !s.DryRun {
			if err := os.WriteFile(o, v.Content, 0644); err != nil {
				return err
			}
		}
		if s.DryRun || ge.Debug {
			log.Println(o, "\n", strings.TrimSpace(string(v.Content)))
		}
	}
	return nil
}

func (ge *Generator) Generate(ctx context.Context, source interface{}) (err error) {
	if err = ge.init(); err != nil {
		return
	}
	var all []File
	for _, tmpl := range ge.Templates {
		var files []File
		files, err = tmpl.Generate(ctx, ge, source)
		if err != nil {
			return err
		}
		if ge.Formatter != nil {
			for i := range files {
				f := files[i]
				file := &files[i]
				if err = ge.Formatter(file); err != nil {
					if ge.Debug {
						log.Println(f.Name, err.Error(), "\n", lineNumber(string(f.Content)))
					}
					return errors.Wrapf(err, "format %v", file.Name)
				}
			}
		}
		all = append(all, files...)
	}
	ge.Files = append(ge.Files, all...)
	return
}

func (ge *Generator) init() error {
	if ge.Debug {
		log.Println("generating")
	}
	if ge.Loader != nil {
		if ge.Debug {
			log.Println("generator loading")
		}
		if err := ge.Loader(); err != nil {
			return err
		}
	}
	return nil
}

type TemplateGenerate interface {
	Generate(ctx context.Context, ge *Generator, source interface{}) ([]File, error)
}
type GenerateTemplate struct {
	Name         string // template name.
	DeriveSource func(ctx context.Context, rootSource interface{}) interface{}
	Skip         func(ctx context.Context, source interface{}) bool   // skip condition (storage constraints or gated by a feature-flag).
	Format       func(ctx context.Context, source interface{}) string // file name format.
	FormatName   string                                               // for single
}

func (tmpl *GenerateTemplate) Generate(ctx context.Context, ge *Generator, source interface{}) (files []File, err error) {
	if tmpl.DeriveSource != nil {
		source = tmpl.DeriveSource(ctx, source)
	}
	all, ok := source.([]interface{})
	if !ok {
		all = []interface{}{source}
	}

	for _, source := range all {
		f, err := tmpl.one(ctx, ge, source)
		if err != nil {
			return nil, err
		}
		if f.Name != "" {
			files = append(files, f)
		}
	}
	return
}
func (tmpl *GenerateTemplate) one(ctx context.Context, ge *Generator, source interface{}) (file File, err error) {
	if tmpl.Skip != nil {
		if tmpl.Skip(ctx, source) {
			return
		}
	}

	b := bytes.NewBuffer(nil)
	if err = ge.Template.ExecuteTemplate(b, tmpl.Name, source); err != nil {
		return file, fmt.Errorf("execute template %q: %w", tmpl.Name, err)
	}

	fn := tmpl.FormatName
	if tmpl.Format != nil {
		fn = tmpl.Format(ctx, source)
	}
	file.Name = fn
	file.Content = b.Bytes()
	if ge.Debug {
		log.Printf("generate %s with %T to %s", tmpl.Name, source, fn)
	}
	return
}

func GoFormatter(f *File) (err error) {
	f.Content, err = format.Source(f.Content)
	if err == nil {
		f.Content, err = imports.Process(f.Name, f.Content, nil)
	}
	return
}

func lineNumber(s string) string {
	lines := strings.Split(s, "\n")
	c := strings.Builder{}
	for i, v := range lines {
		c.WriteString(fmt.Sprintf("%v: ", i))
		c.WriteString(v)
		c.WriteRune('\n')
	}
	return c.String()
}
