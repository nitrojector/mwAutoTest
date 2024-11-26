package mwtest

type Game struct {
	players []*Player
	spawns  []vec2
	board   *Board
}

type GameContext struct {
	exitLoc vec2
}

func NewGame(gameConfigFile string) *Game {
	return nil
}
