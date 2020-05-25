package engine

import "github.com/mpragliola/sangennaro/internal/content"

// NamingStrategy ...
type NamingStrategy interface {
	GetDestinationFilename(p content.Page) string
}
