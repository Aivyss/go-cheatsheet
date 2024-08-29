package reserved_words

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeferExample(t *testing.T) {
	testFunction(t)
}

func testFunction(t *testing.T) (s string) {
	defer func(ss string) {
		fmt.Println("(2) ss =", ss)
		assert.Equal(t, "", ss)
		fmt.Println("(3) s =", s)
		assert.Equal(t, "apple", s)
	}(s)

	s = "banana"
	fmt.Println("(1) s =", s)
	assert.Equal(t, "banana", s)
	return "apple"
}

/*
	(1) s = banana
	(2) ss =
	(3) s = apple
*/
