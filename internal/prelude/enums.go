package prelude

type GameStatus int8
type FinishReason uint8
type Color uint8

const (
	Ongoing  GameStatus = 10
	WhiteWin GameStatus = 1
	BDraw    GameStatus = 0
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
	White Color = 1
	Black Color = 0
)
