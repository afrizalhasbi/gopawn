package schema

import (
	. "gopawn/internal/prelude"
)

type Game struct {
	Notation     string
	FinishReason *FinishReason
	Status       GameStatus
	// int32 because postgres has no native uint
	// corresponds to int4 in postgres (4 bytes)
	StartTimeMillis int32
	IncrTimeMillis  int32
	TurnToMove      Color
	PlayerWhite     string
	PlayerBlack     string
}
