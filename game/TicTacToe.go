package game

import (
	"strconv"
)

const (
	LAST_INDEX = 1 << 8
	DRAW_STATE = 7 | 7<<3 | 7<<3<<3
)

var winCombos = map[string]WinCombos{
	"ROW_0":          {boardSlice: 7, indexes: 1<<0 | 1<<1 | 1<<2},
	"ROW_1":          {boardSlice: 7 << 3, indexes: 1<<3 | 1<<4 | 1<<5},
	"ROW_2":          {boardSlice: 7 << 3 << 3, indexes: 1<<6 | 1<<7 | 1<<8},
	"COLUMN_0":       {boardSlice: 4 | (4 << 3) | (4 << 3 << 3), indexes: 1<<8 | 1<<5 | 1<<2},
	"COLUMN_1":       {boardSlice: 2 | (2 << 3) | (2 << 3 << 3), indexes: 1<<7 | 1<<4 | 1<<1},
	"COLUMN_2":       {boardSlice: 1 | (1 << 3) | (1 << 3 << 3), indexes: 1<<6 | 1<<3 | 1<<0},
	"DIAGONAL_RIGHT": {boardSlice: 84, indexes: 1<<6 | 1<<4 | 1<<2},
	"DIAGONAL_LEFT":  {boardSlice: 273, indexes: 1<<8 | 1<<4 | 1<<0},
}

type WinCombos struct {
	boardSlice int
	indexes    int
}

const (
	WIN = iota
	DRAW
	NONE
	INVALID
)

type Result int

type TicTacToe struct {
	CurrentPlayer *Player
	xBoard        int
	oBoard        int
}

func NewTicTacToe() *TicTacToe {
	return &TicTacToe{
		CurrentPlayer: &Player{Type: &CROSS, OpponentType: &NOUGHT},
		xBoard:        0,
		oBoard:        0,
	}
}

func (t *TicTacToe) GetXBoard() string {
	if t.xBoard == 0 {
		return "0"
	} else {
		return strconv.Itoa(t.xBoard)
	}
}

func (t *TicTacToe) GetOBoard() string {
	if t.oBoard == 0 {
		return "0"
	} else {
		return strconv.Itoa(t.oBoard)
	}
}

func (t *TicTacToe) Place(index int) (Result, int) {
	cell := 1 << index
	if cell > LAST_INDEX || cell < 0 || ((t.xBoard|t.oBoard)&cell == cell) {
		return INVALID, 0
	}
	switch *t.CurrentPlayer.Type {
	case CROSS:
		t.xBoard |= cell
		if win, indexes := t.checkWin(t.CurrentPlayer.Type); win {
			return WIN, indexes
		}
	case NOUGHT:
		t.oBoard |= cell
		if win, indexes := t.checkWin(t.CurrentPlayer.Type); win {
			return WIN, indexes
		}
	}
	if t.checkDraw() {
		return DRAW, 0
	}
	return NONE, 0
}

func (t *TicTacToe) checkWin(playerType *PlayerType) (bool, int) {
	switch *playerType {
	case CROSS:
		for _, combos := range winCombos {
			if (t.xBoard & combos.boardSlice) == combos.boardSlice {
				return true, combos.indexes
			}
			t.CurrentPlayer.Type = &NOUGHT
			t.CurrentPlayer.OpponentType = &CROSS
		}
	case NOUGHT:
		for _, combos := range winCombos {
			if (t.oBoard & combos.boardSlice) == combos.boardSlice {
				return true, combos.indexes
			}
			t.CurrentPlayer.Type = &CROSS
			t.CurrentPlayer.OpponentType = &NOUGHT
		}
	}
	return false, 0
}

func (t *TicTacToe) checkDraw() bool {
	return t.xBoard|t.oBoard == DRAW_STATE
}
