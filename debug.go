// Package debug implements a development-friendly textual handler.
package debug

import (
	"fmt"
	"io"
	"time"

	"github.com/apex/log"
	"github.com/aybabtme/rgbterm"
)

// color function.
type colorFunc func(string) string

// Gray returns a grayed string.
func Gray(s string) string {
	return rgbterm.FgString(s, 140, 140, 140)
}

// Blue returns a blued string.
func Blue(s string) string {
	return rgbterm.FgString(s, 9, 109, 217)
}

// Purple returns a purpled string.
func Purple(s string) string {
	return rgbterm.FgString(s, 96, 97, 190)
}

// Yellow returns a yellowed string.
func Yellow(s string) string {
	return rgbterm.FgString(s, 212, 177, 6)
}

// Red returns a redden string.
func Red(s string) string {
	return rgbterm.FgString(s, 207, 19, 34)
}

// Magenta returns a megenta string.
func Magenta(s string) string {
	return rgbterm.FgString(s, 235, 47, 150)
}

// Colors mapping.
var Colors = [...]colorFunc{
	log.DebugLevel: Purple,
	log.InfoLevel:  Blue,
	log.WarnLevel:  Yellow,
	log.ErrorLevel: Magenta,
	log.FatalLevel: Red,
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

	fmt.Fprintf(h.Writer, "%s %s %s: ", Gray(time), color(level), color(e.Message))

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
