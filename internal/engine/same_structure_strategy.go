package engine

import (
	"log"
	"path/filepath"
	"regexp"

	"github.com/mpragliola/sangennaro/internal/content"
	"github.com/mpragliola/sangennaro/internal/filesystem"
)

// SameStructureStrategy ...
type SameStructureStrategy struct {
	fileSystem filesystem.FileSystem
}

// NewSameStructureStrategy ...
func NewSameStructureStrategy(f filesystem.FileSystem) *SameStructureStrategy {
	return &SameStructureStrategy{
		fileSystem: f,
	}
}

// GetDestinationFilename ...
func (s *SameStructureStrategy) GetDestinationFilename(p content.Page) string {
	relativeSourcePath, err := filepath.Rel(
		s.fileSystem.GetDataPath()+string(filepath.Separator)+"posts",
		p.SourceFilename,
	)

	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(".md$")

	outpath := filepath.Join(
		s.fileSystem.GetPublicPath(), r.ReplaceAllString(relativeSourcePath, ".html"))

	return outpath
}
