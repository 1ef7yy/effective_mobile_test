package errors

import "fmt"

var (
	SongNotFound        = fmt.Errorf("song not found")
	OffsetOutOfRangeErr = fmt.Errorf("offset out of range")
)
