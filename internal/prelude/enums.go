package prelude

type GameStatus int8
type FinishReason uint8
type Color bool

const (
	ongoing  GameStatus = 10
	whiteWin GameStatus = 1
	blackWin GameStatus = -1
	draw     GameStatus = 0
)

const (
	checkmate        FinishReason = 10
	resign           FinishReason = 11
	timeout          FinishReason = 12
	disconnect       FinishReason = 13
	drawAgreed       FinishReason = 20
	drawRepetition   FinishReason = 21
	drawInsufficient FinishReason = 22
	drawStalemate    FinishReason = 23
)

const (
	white Color = true
	black Color = false
)
