package mwtest

/*

 For LLM

	Fixed Game Element
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | | | | | | | | | |
	+ +-+-+-+-+-+-+-+-+-+-+
	| | | | | | | | | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | |1| | | | | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | |     | | | | | | |
	+-+-+ + + +-+-+-+-+-+-+
	| | |  x  | | |2| | | |
	+-+-+ + + +-+-+-+-+-+-+
	|2| |     | | | | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | | | | |1| | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | |#| | | | | | | | |
	+ +-+-+-+-+-+-+-+-+-+-+
	|    *| | | | | | | | |
	+-+-+ +-+-+-+-+-+-+-+-+
	|x|x|E|x|x|x|x|x|x|x|x|
	+-+-+-+-+-+-+-+-+-+-+-+

	*, A, B - Spawn Points (* is yours, A is your 1st teammate, B is your 2nd teammate)
	0, 1, 2, 3, ... - Teleporter (ID)
	x - Inaccessible Region
	E - Exit


	Player/Minotaur Location
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | | | | | | | | | |
	+ +-+-+-+-+-+-+-+-+-+-+
	| | | | | | | | | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | | | | | | | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | | | | | | | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | | | | |2| | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | |1| | | | | | | |
	+-+-+-+-+-+-+-+-+-+-+-+
	| | | | | | | | | | | |
	+-+-+ +-+-+-+-+-+-+-+-+
	|D   #| | | | | | | | |
	+ +-+-+-+-+-+-+-+-+-+-+
	|D   *| | | | | | | | |
	+-+-+-+-+-+-+-+-+-+-+-+

	# - Minotaur
	* - You
	1 - Your 1st teammate
	2 - Your 2nd teammate
	D - Dark Region

	Your teammates are marked with


*/
