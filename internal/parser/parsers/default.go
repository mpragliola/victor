package parsers

import (
	"html/template"
	"log"
	"strings"

	"github.com/mpragliola/sangennaro/internal/content"
	"github.com/mpragliola/sangennaro/internal/frontmatter"
	"github.com/russross/blackfriday/v2"
)

// Default ...
type Default struct {
}

// NewDefault ...
func NewDefault() *Default {
	return &Default{}
}

// Parse ...
func (p *Parser) Parse(source []byte, sourceFilename string) content.Page {
	page := content.Page{
		SourceFilename: sourceFilename,
	}

	fm := frontmatter.NewFrontmatter()
	body, err := fm.Unmarshal(source, &page)

	if err != nil {
		log.Fatal(err)
	}

	page.Body = template.HTML(blackfriday.Run([]byte(body)))

	if len(page.RawTags) == 0 {
		return page
	}

	p.parseTags(&page)

	return page
}

func (p *Parser) parseTags(page *content.Page) *content.Page {
	if len(page.RawTags) == 0 {
		return page
	}

	var tags []string

	for _, tag := range strings.Split(page.RawTags, ",") {
		tags = append(tags, strings.Trim(tag, " "))
	}

	page.Tags = tags

	return page
}
