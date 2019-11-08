package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWorld = "Go!"
const countDownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Sleepy   func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.Sleepy(cs.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprintln(out, finalWorld)
}
