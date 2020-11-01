package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("initial test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := SpyCountdownSleeper{}
		Countdown(&buffer, &spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		assertStrings(t, got, want)

		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		}
	})

	t.Run("sleep before print", func(t *testing.T) {
		ssw := SpySleepWriter{}
		Countdown(&ssw, &ssw)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, ssw.Calls) {
			t.Errorf("wanted calls %v got %v", want, ssw.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}

}

type SpyTime struct {
	durationSlept time.Duration
}

func (st *SpyTime) Sleep(duration time.Duration) {
	st.durationSlept = duration
}

// SpyCountdownSleeper is a struct
type SpyCountdownSleeper struct {
	Calls int
}

// Sleep mocks time.Sleep, it increments Calls
func (ss *SpyCountdownSleeper) Sleep() {
	ss.Calls++
}

// type DefaultSleeper struct{}

// func (ds *DefaultSleeper) Sleep(duration time.Duration) {
// 	time.Sleep(duration)
// }

type SpySleepWriter struct {
	Calls []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (ssw *SpySleepWriter) Sleep() {
	ssw.Calls = append(ssw.Calls, sleep)
}

func (ssw *SpySleepWriter) Write(p []byte) (n int, e error) {
	ssw.Calls = append(ssw.Calls, write)
	return
}
