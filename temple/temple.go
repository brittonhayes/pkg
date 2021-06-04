// Package temple provides useful utilities
// for rendering go text/templates and html/templates
// with less boilerplate
package temple

import (
	"bytes"
	"html/template"

	"github.com/Masterminds/sprig"
)

// Render renders the template string into a buffer and includes
// some useful helper functions
func Render(templateString string, data interface{}) (*bytes.Buffer, error) {
	t, err := template.New("temple").Funcs(sprig.FuncMap()).Parse(templateString)
	if err != nil {
		return nil, err
	}

	b, err := exec(t, data)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func exec(t *template.Template, data interface{}) (*bytes.Buffer, error) {
	var buff bytes.Buffer
	if err := t.Execute(&buff, data); err != nil {
		return nil, err
	}

	return &buff, nil
}
