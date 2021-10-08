package common

import (
	"github.com/google/uuid"
)

const (
	EVENT_VIDEO_CONVERT = "video.convert"
)

type User struct {
	Id   int       `json:"-" db:"id"`
	Code uuid.UUID `json:"code" db:"code" binding:"required"`
}
