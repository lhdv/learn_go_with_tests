package poker

import "io"

// Game specifies what a poker game must perform
type Game interface {
	Start(numberOfPlayer int, alertDestination io.Writer)
	Finish(winner string)
}
