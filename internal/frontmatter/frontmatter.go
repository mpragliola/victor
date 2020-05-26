package frontmatter

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

var delimiter = []byte("---")

// Unmarshal parses a text content in form of []byte,
// extracting the front matter and returning the remaining
// content body
func Unmarshal(content []byte, v interface{}) (body []byte, err error) {
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
