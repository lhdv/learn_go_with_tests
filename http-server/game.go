package poker

// Game specifies what a poker game must perform
type Game interface {
	Start(numberOfPlayer int)
	Finish(winner string)
}
