package utils

import (
	"bytes"
	"html/template"
	"strconv"
	"strings"
)

func mustExecuteTemplate(t *template.Template, data any) (result string) {
	buffer := new(bytes.Buffer)
	err := t.Execute(buffer, data)
	if err != nil {
		panic(err)
	}
	result = buffer.String()
	return
}

// BindStringTemplateByMap key in string format = {{.key}}
func BindStringTemplateByMap(stringTemplate string, params map[string]any) string {
	stringTemplate = strings.ReplaceAll(stringTemplate, "\n", " ")
	quoteString := strconv.Quote(stringTemplate)
	t := template.Must(template.New("").Parse(quoteString))

	return mustExecuteTemplate(t, params)
}

/*
BindStringTemplateByStructWithFuncMap
ex)

	funcMap := map[string]any {
		"formatDate": func(t time.Time){ return t.Format(time.DateOnly)},
	}
	{{formatDate .Key}}
*/
func BindStringTemplateByStructWithFuncMap(stringTemplate string, params any, funcMap map[string]any) string {
	t := template.Must(template.New("").Funcs(funcMap).Parse(stringTemplate))
	return mustExecuteTemplate(t, params)
}

// BindStringTemplateByStruct key in string format = {{.Key}}
func BindStringTemplateByStruct(stringTemplate string, params any) string {
	t := template.Must(template.New("").Parse(stringTemplate))
	return mustExecuteTemplate(t, params)
}
