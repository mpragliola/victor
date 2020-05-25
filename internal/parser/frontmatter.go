package parser

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

var delimiter = []byte("---")

// Frontmatter is an utility struct intended to parse the "front matter" as
// defined by Jekyll; markdown files can have an header delimited by three dashes
// (---)  at the very beginning of the file and between the header and the real
// content. This header can contain metadata as YAML fragment, which can be
// parsed with these methods.
type Frontmatter struct {
}

// NewFrontmatter ...
func NewFrontmatter() *Frontmatter {
	return &Frontmatter{}
}

// Unmarshal parses a text content in form of []byte,
// extracting the front matter and returning the remaining
// content body
func (f Frontmatter) Unmarshal(content []byte, v interface{}) (body []byte, err error) {
	// Let's short circuit if content doesn't begin with delimiter
	if !bytes.HasPrefix(content, delimiter) {
		return content, nil
	}

	// Extract front matter and parse its YAML
	parts := bytes.SplitN(content, delimiter, 3)
	body = parts[2]
	err = yaml.Unmarshal(parts[1], v)

	return
}
