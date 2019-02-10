package main

import (
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/generate"
	"github.com/azer/snakecase"
	"os"
	"reflect"
	"strings"
	"text/template"
)

const (
	skipTag     = "-"
	optionalTag = "omitempty"
	tplText     = `map[string]*schema.Schema{ {{range $tag := .Tags}}
	"{{ $tag.Name }}": { {{ if $tag.IsArray }}
		Type:     schema.TypeSet,
{{if $tag.Optional}}		Optional: true,{{else}}		Required: true,{{end}}
		Elem:     &schema.Schema{Type: schema.Type{{ $tag.Type }}},
		Set:      schema.Hash{{ $tag.Type }},
{{else}}
		Type:     schema.Type{{ $tag.Type }},
{{if $tag.Optional}}		Optional: true,{{else}}		Required: true,{{end}}{{end}}
	},{{end}}
}


generate.{{.Name}}{
{{range $tag := .Tags}} {{ if $tag.IsArray }}
	{{ $tag.FieldName }}: SchemaSetTo{{$tag.Type}}List(d.Get("{{$tag.Name}}").(*schema.Set)),{{else}}
	{{ $tag.FieldName }}: d.Get("{{$tag.Name}}").({{$tag.TypeLower}}),{{end}}{{end}}
}
`
)

type Tag struct {
	Name      string
	FieldName string
	Skip      bool
	Optional  bool
	IsArray   bool
	Type      string
	TypeLower string
}

func main() {
	inter := generate.Certificate{}
	v := reflect.ValueOf(inter)
	t := reflect.TypeOf(inter)
	tags := make([]Tag, 0)
	for index := 0; index < v.NumField(); index++ {
		tField := t.Field(index)
		tag := parseInTag(tField.Tag.Get("json"), tField.Name)
		typeName := tField.Type.Name()
		if tField.Type.Kind() == reflect.Slice {
			typeName = tField.Type.Elem().Name()
			tag.IsArray = true
		}
		tag.Type = strings.Title(typeName)
		tag.TypeLower = typeName
		tags = append(tags, tag)

	}
	tpl := template.New("toschema")
	tpl, err := tpl.Parse(tplText)
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "toschema", struct {
		Name string
		Tags []Tag
	}{t.Name(), tags})
}
func parseInTag(tag, fieldName string) Tag {
	if tag == "" {
		return Tag{
			FieldName: fieldName,
			Name:      snakecase.SnakeCase(fieldName),
		}
	}
	tag = strings.TrimSpace(tag)
	splitedTag := strings.Split(tag, ",")
	name := splitedTag[0]
	if name == skipTag {
		name = snakecase.SnakeCase(fieldName)
	}
	if name == "" {
		name = snakecase.SnakeCase(fieldName)
	}

	return Tag{
		FieldName: fieldName,
		Name:      name,
		Optional:  hasOptionalTag(splitedTag[1:]),
	}
}
func hasOptionalTag(tags []string) bool {
	for _, tag := range tags {
		if tag == optionalTag {
			return true
		}
	}
	return false
}
