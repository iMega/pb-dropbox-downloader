package l10n

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestLoadCatalog(t *testing.T) {
	r := strings.NewReader(`
test=Test
`)
	lang, err := language.Parse(language.English.String())
	assert.NoError(t, err)

	cat, err := loadCatalog(lang, io.NopCloser(r))
	assert.NoError(t, err)

	message.DefaultCatalog = cat

	printer := message.NewPrinter(lang)
	actual := printer.Sprintf("test")

	assert.Equal(t, actual, "Test")
}
