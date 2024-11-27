package tpl

import (
	"bytes"
	"errors"
	"go-cheatsheet/utils"
	"html/template"
)

var (
	LoadTemplateErr  = errors.New("load_template_error")
	ParseTemplateErr = errors.New("parse_template_error")
)

type TextTemplate[T any] string

func (s *TextTemplate[T]) Parse(t T) (*string, error) {
	tpl, err := template.New("TextTemplate").Parse(string(*s))
	if err != nil {
		return nil, LoadTemplateErr
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, t)
	if err != nil {
		return nil, ParseTemplateErr
	}

	return utils.Pointer(buf.String()), nil
}
