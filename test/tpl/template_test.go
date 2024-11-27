package tpl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	tpl2 "go-cheatsheet/tpl"
	"testing"
)

func TestTextTemplate(t *testing.T) {
	t.Run("parse", func(t *testing.T) {
		// given
		type testEntity struct {
			First  string
			Second string
		}

		var template tpl2.TextTemplate[testEntity] = `first: {{.First}}, second: {{.Second}}`

		entity := testEntity{
			First:  "1",
			Second: "2",
		}
		result, err := template.Parse(entity)

		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf(`first: %s, second: %s`, entity.First, entity.Second), *result)
	})
}
