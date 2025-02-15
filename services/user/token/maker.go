package token

import (
	"time"
)

type Maker interface {
	CreateToken(userId int32, userRole string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
