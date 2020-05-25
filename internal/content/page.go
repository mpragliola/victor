package content

import "html/template"

// Page -
type Page struct {
	Title          string        `yaml:"title"`
	RawTags        string        `yaml:"tags"`
	Tags           []string      `yaml:"-"`
	Author         string        `yaml:"author"`
	SourceFilename string        `yaml:"-"`
	Body           template.HTML `yaml:"-"`
	Assets         string        `yaml:"-"`
}
