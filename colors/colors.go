package colors

import (
  "github.com/aybabtme/rgbterm"
)

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
