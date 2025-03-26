package errors

import "fmt"

var (
	SongNotFoundErr     = fmt.Errorf("song not found")
	OffsetOutOfRangeErr = fmt.Errorf("offset out of range")
	AlreadyExistsErr    = fmt.Errorf("song already exists")
)
