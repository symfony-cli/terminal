package terminal

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
)

// Spinner struct
type Spinner struct {
	Writer          io.Writer
	PrefixIndicator string
	PrefixText      string
	SuffixIndicator string
	SuffixText      string

	chars    []string
	delay    time.Duration
	lock     *sync.RWMutex
	active   bool
	stopChan chan struct{}
	cursor   Cursor
}

// NewSpinner creates a spinner
func NewSpinner(w io.Writer) *Spinner {
	chars := []string{"◐", "◓", "◑", "◒"}
	if runtime.GOOS == "windows" {
		chars = []string{"|", "/", "-", "\\"}
	}

	return &Spinner{
		Writer:          w,
		SuffixIndicator: "</>",
		PrefixIndicator: " <fg=yellow>",
		PrefixText:      "",
		SuffixText:      "",
		chars:           chars,
		delay:           150 * time.Millisecond,
		lock:            &sync.RWMutex{},
		active:          false,
		stopChan:        make(chan struct{}),
		cursor:          NewCursor(w),
	}
}

// Active returns whether the spinner is currently spinning
func (s *Spinner) Active() bool {
	return s.active
}

// Start starts the spinner
func (s *Spinner) Start() {
	if !Stdin.IsInteractive() {
		return
	}

	s.lock.Lock()
	if s.active {
		s.lock.Unlock()
		return
	}
	s.active = true
	s.lock.Unlock()

	go func() {
		b := bytes.Buffer{}
		cursor := Cursor{Writer: &b}

		for {
			for i := 0; i < len(s.chars); i++ {
				select {
				case <-s.stopChan:
					return
				default:
					s.lock.Lock()
					b.Reset()
					cursor.SavePosition()
					fmt.Fprintf(&b, "%s%s%s  %s%s", s.PrefixText, s.PrefixIndicator, s.chars[i], s.SuffixIndicator, s.SuffixText)
					cursor.ClearLineAfter()
					cursor.RestorePosition()
					b.WriteTo(s.Writer)
					s.lock.Unlock()
					time.Sleep(s.delay)
				}
			}
		}
	}()
}

// Stop stops the spinner
func (s *Spinner) Stop() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.active {
		s.active = false
		s.stopChan <- struct{}{}
		s.cursor.ClearLineAfter()
	}
}
