package content

import "time"

// Metadata .
type Metadata struct {
	Title     string
	Author    string
	Created   time.Time
	Updated   time.Time
	Published time.Time
}
