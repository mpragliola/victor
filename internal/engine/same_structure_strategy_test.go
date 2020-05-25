package engine

import (
	"testing"

	"github.com/mpragliola/sangennaro/internal/content"
)

func TestSum(t *testing.T) {

	s := new(SameStructureStrategy)
	page := content.Page{}

	got := s.GetDestinationFilename(page)
	want := ""

	if got != want {
		t.Errorf("No good %s %s", got, want)
	}
}
