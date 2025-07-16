package schema

import (
	. "gopawn/internal/prelude"
	"time"
)

type Game struct {
	notation     string
	finishReason *FinishReason
	status       GameStatus
	seconds      time.Time
	increment    time.Time
	turnToMove   Color
}
