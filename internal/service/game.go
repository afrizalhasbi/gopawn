package service

import (
	"database/sql"
	"fmt"
	"gopawn/internal/data/schema"
	. "gopawn/internal/prelude"
	"math"
	"math/rand/v2"
	"time"
)

type GameRequest struct {
	reqTimeMillis int32
	reqIncrMillis int32
}

type WaitedPlayer struct {
	Name         string
	Elo          uint16
	WaitedMillis uint16
	Request      GameRequest
}

type GameService struct {
	DB          *sql.DB
	QueueBuffer chan WaitedPlayer
	Queue       []WaitedPlayer
	ActiveGames chan schema.Game
}

func (s *GameService) Serve() {
	for true {
		for v := range s.QueueBuffer {
			s.Queue = append(s.Queue, v)
		}
		s.matchQueue()
	}
}

func (s *GameService) newGame(p1 *WaitedPlayer, p2 *WaitedPlayer) schema.Game {
	var players = []string{p1.Name, p2.Name}
	var whitePlayer = players[rand.IntN(2)]
	var blackPlayer string

	if whitePlayer == p1.Name {
		blackPlayer = p2.Name
	} else {
		blackPlayer = p1.Name
	}

	return schema.Game{
		Notation:        "string",
		FinishReason:    nil,
		Status:          Ongoing,
		StartTimeMillis: p1.Request.reqTimeMillis,
		IncrTimeMillis:  p1.Request.reqIncrMillis,
		TurnToMove:      White,
		PlayerWhite:     whitePlayer,
		PlayerBlack:     blackPlayer,
	}
}

func (s *GameService) matchQueue() {
	var target = 0
	var src = 1
	for target < len(s.Queue) {
		game, err := s.matchPlayers(&s.Queue[target], &s.Queue[src])
		if err != nil {
			src += 1
		} else {
			s.ActiveGames <- game
			return
		}
	}
}

func (s *GameService) matchPlayers(p1 *WaitedPlayer, p2 *WaitedPlayer) (schema.Game, error) {
	p1.WaitedMillis = uint16(time.Now().UnixMilli()) - p1.WaitedMillis
	p2.WaitedMillis = uint16(time.Now().UnixMilli()) - p2.WaitedMillis
	if p1.WaitedMillis > 30_000 || p2.WaitedMillis > 30_000 {
		if math.Abs(float64(p1.Elo)-float64(p2.Elo)) < 80 {
			return s.newGame(p1, p2), nil
		}
	} else {
		if math.Abs(float64(p1.Elo)-float64(p2.Elo)) < 40 {
			return s.newGame(p1, p2), nil
		}
	}
	return schema.Game{}, fmt.Errorf("Invalid elo gap")
}

// func (s *GameService) funcname() {}

// func (s *GameService) funcname() {}
// // func (s *GameService) funcname() {}
//
