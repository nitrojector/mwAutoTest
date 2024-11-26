package mwtest

type Player struct {
	isMinotaur bool
	// false implies player is Adventurer

	loc vec2
	// current location of player

	d Deck
	// deck of cards
}
