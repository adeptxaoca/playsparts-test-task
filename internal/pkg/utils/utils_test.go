package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuoteString(t *testing.T) {
	str := QuoteString("', pass = '123', WHERE id = 1; SELECT name = '")
	assert.Equal(t, "''', pass = ''123'', WHERE id = 1; SELECT name = '''", str)
}
