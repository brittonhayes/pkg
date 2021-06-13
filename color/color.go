// Package color provides simple utilities for
// using colors based on their message significance.
//
// e.g. Success is light green, Error is bright red
//
package color

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type Color int

const (

	// Info is a light blue color
	// to represent informational level text
	Info Color = iota

	// Error is a bright red color
	// for error messages
	Error

	// Success is a bright green color
	// for success messages
	Success

	// Warn is a bright yellow color
	// for indicating a potential issue
	Warn
)

// String returns an adaptive color to represent
// info, error, warn, or success
// based on the color theme of the terminal
func (c Color) String() string {
	var (
		infoColor    = lipgloss.AdaptiveColor{Light: "#073b4c", Dark: "#4cc9f0"}
		errorColor   = lipgloss.AdaptiveColor{Light: "#9d0208", Dark: "#f72585"}
		warnColor    = lipgloss.AdaptiveColor{Light: "#e85d04", Dark: "#f48c06"}
		successColor = lipgloss.AdaptiveColor{Light: "#469d89", Dark: "#06d6a0"}
	)
	if termenv.HasDarkBackground() {
		return [...]string{infoColor.Dark, errorColor.Dark, warnColor.Dark, successColor.Dark}[c]
	}

	return [...]string{infoColor.Light, errorColor.Light, warnColor.Light, successColor.Light}[c]
}

