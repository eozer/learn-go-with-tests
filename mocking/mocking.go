package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
	oneSecond      = 1 * time.Second
)

// Sleeper is an interface
type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown counts down from 3, printing each number on a new line (with a 1 second pause) and when
// it reaches zero it will print "Go!" and exit.
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(writer, "%d\n", i)
	}
	sleeper.Sleep()
	fmt.Fprintf(writer, finalWord)
}

func main() {
	cs := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, cs)
}
