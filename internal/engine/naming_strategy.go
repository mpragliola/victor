package engine

import "github.com/mpragliola/victor/internal/content"

// NamingStrategy ...
type NamingStrategy interface {
	GetDestinationFilename(p content.Page) string
}
