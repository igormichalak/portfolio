package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
)

type Parser struct {
	MD goldmark.Markdown
}

func (p *Parser) Convert(source []byte) (string, error) {
	var buf bytes.Buffer
	if err := p.MD.Convert(source, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
