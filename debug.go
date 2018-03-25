// Package debug implements a development-friendly textual handler.
package debug

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/sqrthree/debug/colors"
)

// color function.
type colorFunc func(string) string

// Colors mapping.
var Colors = [...]colorFunc{
	log.DebugLevel: colors.Purple,
	log.InfoLevel:  colors.Blue,
	log.WarnLevel:  colors.Yellow,
	log.ErrorLevel: colors.Magenta,
	log.FatalLevel: colors.Red,
}

// Strings mapping.
var Strings = [...]string{
	log.DebugLevel: "DEBUG",
	log.InfoLevel:  "INFO",
	log.WarnLevel:  "WARN",
	log.ErrorLevel: "ERROR",
	log.FatalLevel: "FATAL",
}

// Handler implementation.
type Handler struct {
	mu sync.Mutex
	Writer io.Writer
}

// New returns a new handle.
func New(w io.Writer) *Handler {
	return &Handler{
		Writer: w,
	}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	color := Colors[e.Level]
	level := Strings[e.Level]

	names := e.Fields.Names()

	time := formatDateString(e.Timestamp.Local())

	h.mu.Lock()
	defer h.mu.Unlock()

	fmt.Fprintf(h.Writer, "%s %s %s: ", colors.Gray(time), color(level), color(e.Message))

	for _, name := range names {
		fmt.Fprintf(h.Writer, "%s%s%v ", color(name), "=", e.Fields.Get(name))
	}

	fmt.Fprintln(h.Writer)

	return nil
}

// formatDateString formats t to a human-friendly string.
func formatDateString(t time.Time) string {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()

	return fmt.Sprintf("%v-%02d-%02v %02v:%02v:%02v", year, month, day, hour, minute, second)
}
