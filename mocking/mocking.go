package main

import 
import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
	countdownEnd   = 0
	sleep          = "sleep"
	write          = "write"
)

// tracking how many times Sleep() is called so we can check our tests

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {
}

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > countdownEnd; i-- {
		_, err := fmt.Fprintln(w, i)
		if err != nil {
			return
		}
		s.Sleep()
	}
	_, err := fmt.Fprint(w, finalWord)
	if err != nil {
		return
	}
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
